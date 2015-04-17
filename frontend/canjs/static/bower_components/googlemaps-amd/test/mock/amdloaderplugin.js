define(function(_) {
  var root = this;

  /**
   * Implements https://github.com/amdjs/amdjs-api/wiki/Loader-Plugins
   * @constructor
   */
  var MockAmdLoaderPlugin = function() {
    spyOn(this, 'load').andCallThrough();
  };

  MockAmdLoaderPlugin.prototype.load = function(name, parentRequire, onload, config) {
  };

  MockAmdLoaderPlugin.prototype.resolveWith = function(var_args) {
    this.getOnload().apply(root, arguments);
  };

  MockAmdLoaderPlugin.prototype.rejectWith = function(var_args) {
    this.getOnerror().apply(root, argument);
  };

  MockAmdLoaderPlugin.prototype.getLoadPath = function() {
    return this.getLoadArgs_()[0];
  };

  MockAmdLoaderPlugin.prototype.getOnload = function() {
    return this.getLoadArgs_()[2];
  };

  MockAmdLoaderPlugin.prototype.getOnerror = function() {
    return this.getOnload().onerror;
  };


  MockAmdLoaderPlugin.prototype.getLoadArgs_ = function() {
    if (!this.load.callCount) {
      throw new Error('Unable to get load arguments: load was never called');
    }

    return this.load.mostRecentCall.args;
  };


  return MockAmdLoaderPlugin;
});
