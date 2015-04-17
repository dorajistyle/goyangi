define([
  'googlemaps',
  'sinon',
  'mock/require',
  'mock/googlemaps',
  'mock/amdloaderplugin'
], function(googlemapsplugin, sinon, MockRequire, MockGoogleMaps, MockAsync) {

  describe('The googlemaps AMD loader plugin', function() {
    var mockRequire, mockAsync, mockGoogleMaps, onload, config;
    var REQUIRE_DELAY = 100;
    var clock;
    var STUB_API_KEY = 'STUB_API_KEY';
    var STUB_URL = 'STUB_URL'

    function restoreGlobalMocks() {
      mockRequire.restore();
      mockGoogleMaps.restore();
      clock.restore();
    }

    beforeEach(function() {
      mockGoogleMaps = new MockGoogleMaps();
      mockAsync = new MockAsync();
      mockRequire = new MockRequire();

      mockRequire.useMockRequire();
      mockRequire.setRequireDelay(REQUIRE_DELAY);

      mockGoogleMaps.removeGlobalLibrary();

      onload = jasmine.createSpy('onload');
      onload.onerror = jasmine.createSpy('onerror');

      config = {
        googlemaps: {
          async: mockAsync,
          url: STUB_URL
        }
      };

      clock = sinon.useFakeTimers();
    });

    afterEach(restoreGlobalMocks);


    describe('load', function() {

      function resolveGoogleMapsLibrary() {
        mockGoogleMaps.exposeMockLibrary();
        mockAsync.resolveWith([]);
        clock.tick(REQUIRE_DELAY);
      }

      it('should load the google maps library using the async plugin (without params)', function() {
        googlemapsplugin.load('', mockRequire, onload, config);

        expect(mockAsync.getLoadPath()).toEqual(STUB_URL + '?sensor=false');
      });

      it('should load the google maps library using the async plugin (with params)', function() {
        config.googlemaps.params = {
          sensor: true,
          key: STUB_API_KEY,
          foo: 'bar'
        }

        googlemapsplugin.load('', mockRequire, onload, config);

        expect(mockAsync.getLoadPath()).toMatch(STUB_URL + '?');
        expect(mockAsync.getLoadPath()).toMatch('key=' + STUB_API_KEY);
        expect(mockAsync.getLoadPath()).toMatch('sensor=true');
        expect(mockAsync.getLoadPath()).toMatch('foo=bar');
      });

      it('should not require a url or apiKey', function() {
        config.googlemaps = {
          async: mockAsync
        };
        googlemapsplugin.load('', mockRequire, onload, config);

        resolveGoogleMapsLibrary();

        expect(onload).toHaveBeenCalled();
        expect(onload.onerror).not.toHaveBeenCalled();
      });

      it('should resolve with the global google maps library', function() {
        googlemapsplugin.load('', mockRequire, onload, config);

        resolveGoogleMapsLibrary();

        expect(onload).toHaveBeenCalledWith(mockGoogleMaps);
        expect(onload.onerror).not.toHaveBeenCalled();
      });

      it('should not load the google maps library, if it is already in the global namespace', function() {
        mockGoogleMaps.exposeMockLibrary();

        googlemapsplugin.load('', mockRequire, onload, config);

        // AMD loaders were not used
        expect(mockAsync.load).not.toHaveBeenCalled();
        expect(mockRequire.require).not.toHaveBeenCalled();

        // Resolve with exposed google.maps
        expect(onload).toHaveBeenCalledWith(mockGoogleMaps);
        expect(onload.onerror).not.toHaveBeenCalled();
      });

      it('should reject if there is no global google maps library', function() {
        var onErrorArg;

        googlemapsplugin.load('', mockRequire, onload, config);

        mockGoogleMaps.removeGlobalLibrary();
        mockAsync.resolveWith([]);
        clock.tick(REQUIRE_DELAY);

        expect(onload.onerror).toHaveBeenCalled();

        onErrorArg = onload.onerror.mostRecentCall.args[0];
        expect(onErrorArg instanceof Error).toEqual(true);
      });

      it('should resolve with null during a build', function() {
        config.isBuild = true;

        googlemapsplugin.load('', mockRequire, onload, config);

        expect(onload).toHaveBeenCalledWith(null);

        // Should not have other side effects
        expect(mockAsync.load).not.toHaveBeenCalled();
        expect(mockRequire.require).not.toHaveBeenCalled();
        expect(onload.onerror).not.toHaveBeenCalled();
      });

    });

  });
});
