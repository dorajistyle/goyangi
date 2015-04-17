define([
  'underscore',
  'mock/require',
  'sinon'
], function(_, MockRequire, sinon) {
  var requireOrig = window.require;
  var defineOrig = window.define;


  describe('A MockRequire', function() {
    var mockRequire;
    var clock;

    beforeEach(function() {
      mockRequire = new MockRequire();
      clock = sinon.useFakeTimers();
    });

    afterEach(function() {
      window.require = requireOrig;
      window.define = defineOrig;
      clock.restore();
    });


    describe('useMockRequire', function() {

      beforeEach(function() {
        mockRequire.require.andReturn(void 0);
      });


      it('should replace the global require with its own', function() {
        mockRequire.useMockRequire();
        window.require('foo', 'bar');

        expect(mockRequire.require).toHaveBeenCalledWith('foo', 'bar');
      });

      it('should bind the global require to its own context', function() {
        mockRequire.useMockRequire();

        mockRequire.require.andCallFake(function() {
          expect(this).toEqual(mockRequire);
        });

        window.require();

        expect(mockRequire.require).toHaveBeenCalled();
      });

    });


    describe('useMockDefine', function() {

      beforeEach(function() {
        spyOn(mockRequire, 'define').andReturn(void 0);
      });

      it('should replace the global define with its own', function() {
        mockRequire.useMockDefine();
        window.define('foo', 'bar');

        expect(mockRequire.define).toHaveBeenCalledWith('foo', 'bar');
      });

      it('should bind the global define to its own context', function() {
        mockRequire.useMockDefine();

        mockRequire.define.andCallFake(function() {
          expect(this).toEqual(mockRequire);
        });

        window.define();

        expect(mockRequire.define).toHaveBeenCalled();
      });

    });


    describe('restore', function() {

      it('should restore the global require', function() {
        mockRequire.useMockRequire();
        mockRequire.restore();

        expect(window.require).toEqual(requireOrig);
      });

      it('should restore the global define', function() {
        mockRequire.useMockDefine();
        mockRequire.restore();

        expect(window.define).toEqual(defineOrig);
      });

    });


    describe('require/define', function() {
      var fooModule, barModule;
      var callback, errback;

      beforeEach(function() {
        fooModule = 'FOO_MODULE_STUB';
        barModule = 'BAR_MODULE_STUB';
        mockRequire.define('foo', function(){ return fooModule; });
        mockRequire.define('bar', function() { return barModule; });

        callback = jasmine.createSpy('callback');
        errback = jasmine.createSpy('errback');
      });


      it('should callback with a defined module', function() {
        mockRequire.require(['foo'], callback);

        expect(callback).toHaveBeenCalledWith(fooModule);
      });

      it('should callback with multiple defined modules', function() {
        mockRequire.require(['foo', 'bar'], callback);

        expect(callback).toHaveBeenCalledWith(fooModule, barModule);
      });

      it('should not errback if all the modules are defined', function() {
        mockRequire.require(['foo', 'bar'], callback, errback);

        expect(errback).not.toHaveBeenCalled();
      });

      it('shoudl errback if a single required module is not defined', function() {
        mockRequire.require(['notDefined'], callback, errback);

        expect(errback).toHaveBeenCalled();
        expect(callback).not.toHaveBeenCalled();
      });

      it('should errback if all the required modules are not defined', function() {
        mockRequire.require(['notDefined1', 'notDefined2'], callback, errback);

        expect(errback).toHaveBeenCalled();
        expect(callback).not.toHaveBeenCalled();
      });

      it('should errback if some of the required modules are not defined', function() {
        mockRequire.require(['foo', 'notDefined', 'bar'], callback, errback);

        expect(errback).toHaveBeenCalled();
        expect(callback).not.toHaveBeenCalled();
      });

      it('should include ids of missing modules in errback', function() {
        var errorObj;

        mockRequire.require(['notDefined1', 'foo', 'notDefined2'], callback, errback);

        errorObj = errback.mostRecentCall.args[0];
        expect(errorObj.requireModules).toEqual(['notDefined1', 'notDefined2']);
      });

      describe('asynchronous specs', function() {
        var DELAY;

        beforeEach(function() {
          DELAY = 100;
          mockRequire.setRequireDelay(DELAY);
        });


        it('should callback only after a specified delay', function() {
          mockRequire.require(['foo', 'bar'], callback);

          expect(callback).not.toHaveBeenCalled();

          clock.tick(DELAY);

          expect(callback).toHaveBeenCalled();
        });

        it('should errback only after a specified delay', function() {
          mockRequire.require(['undefinedModuleId'], callback, errback);

          expect(errback).not.toHaveBeenCalled();

          clock.tick(DELAY);

          expect(errback).toHaveBeenCalled();
        });

        describe('with delay set to null', function() {
          beforeEach(function() {
            mockRequire.setRequireDelay(null);
          });


          it('should callback immediately if delay is null', function() {
            mockRequire.require(['foo', 'bar'], callback);

            expect(callback).toHaveBeenCalled();
          });

          it('should errback immediately if delay is null', function() {
            mockRequire.require(['undefinedModule'], callback, errback);

            expect(errback).toHaveBeenCalled();
          });

        });

      });

    });


    describe('shouldHaveRequired', function() {

      it('should pass the if the moduleIds were required', function() {
        mockRequire.require(['foo', 'bar']);
        mockRequire.require(['shnaz', 'lorsch']);

        mockRequire.shouldHaveRequired('foo', 'lorsch');
      });

    });


    describe('unssetModule', function() {

      it('should errback when trying to require an unsset module', function() {
        var errback = jasmine.createSpy('errback');
        var MODULE_ID = 'someModuleToBeUnsset';
        mockRequire.define(MODULE_ID, function() {
          return 'STUB_MODULE';
        });
        mockRequire.unssetModule(MODULE_ID);

        mockRequire.require([MODULE_ID], null, errback);
        expect(errback).toHaveBeenCalled();
      });

      it('should throw an error if the module was never set', function() {
        expect(function() {
          mockRequire.unssetModule('someModuleWeNeverDefined');
        }).toThrow();
      });

    });

  });

});