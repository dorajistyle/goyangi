define([
  'async'
], function(async) {
  var root = this;

  /**
   * Google maps AMD loader plugin.
   *
   * Example:
   *  // All configs options are optional.
   *  require.config({
   *      googlemaps: {
   *        url: 'https://maps.googleapis.com/maps/api/js',
   *        params: {
   *          key: 'abcd1234',
   *          libraries: 'geometry',
   *          sensor: true                // Defaults to false
   *        },
   *        async: asyncLoaderPlugin      // Primarly for providing test stubs.
   *      }
   *  });
   *
   *  require(['googlemaps!'], function(gmaps) {
   *    var map = new gmaps.map('map-canvas);
   *  });
   *
   */
  var googlemapsPlugin = {
    load: function(name, parentRequire, onload, opt_config) {
      var googleMapsLoader;
      var config = opt_config || {};

      if (config.isBuild) {
        onload(null);
        return;
      }

      googleMapsLoader = new GoogleMapsLoader(parentRequire, onload, config.googlemaps || {});
      googleMapsLoader.load();
    }
  };


  /**
   * Helper class for googlemaps loader plugin.
   */
  var GoogleMapsLoader = function(require, onload, config) {
    this.require_ = require;
    this.onload_ = onload || NOOP;
    this.baseUrl_ = config.url || GoogleMapsLoader.DEFAULT_BASE_URL;
    this.async_ = config.async || async;
    this.params_ = this.normalizeParams_(config.params);
  };


  GoogleMapsLoader.prototype.load = function() {
    if (this.isGoogleMapsLoaded_()) {
      this.resolveWith_(this.getGlobalGoogleMaps_());
    }
    else {
      this.loadGoogleMaps_();
    }
  };


  GoogleMapsLoader.prototype.loadGoogleMaps_ = function() {
    var self = this;

    var onAsyncLoad = function() {
      // Ensure correct context
      self.resolveWithGoogleMaps_(self);
    };
    onAsyncLoad.onerror = this.onload_.onerror;

    this.async_.load(this.getGoogleUrl_(), this.require_, onAsyncLoad, {});
  };


  GoogleMapsLoader.prototype.getGoogleUrl_ = function() {
    return this.baseUrl_ + '?' + this.serializeParams_();
  };


  GoogleMapsLoader.prototype.resolveWithGoogleMaps_ = function() {
    if (!this.isGoogleMapsLoaded_()) {
      this.reject_();
      return;
    }

    this.resolveWith_(this.getGlobalGoogleMaps_());
  };


  /** Thanks to http://jsfiddle.net/rudiedirkx/U5Tyb/1/ */
  GoogleMapsLoader.prototype.serializeParams_ = function() {
    var encodedParams = [];
    for (var key in this.params_) {
      if (this.params_.hasOwnProperty(key)) {
        var value = this.params_[key];
        var isObject = (typeof value === 'object');
        var encodedParam = encodeURIComponent(key) + "=" + encodeURIComponent(value);
        var serializedValue = isObject ? this.serializeParams_(value, key) : encodedParam;

        encodedParams.push(serializedValue);
      }
    }

    return encodedParams.join("&");
  };


  GoogleMapsLoader.prototype.normalizeParams_ = function(params) {
    var defaultParams = {
      sensor: false
    };
    params || (params = {});

    params.sensor = (params.sensor == void 0) ? defaultParams.sensor : params.sensor;

    return params;
  };


  GoogleMapsLoader.prototype.isGoogleMapsLoaded_ = function() {
    return root.google && root.google.maps;
  };


  GoogleMapsLoader.prototype.getGlobalGoogleMaps_ = function() {
    return root.google ? root.google.maps : undefined;
  };


  GoogleMapsLoader.prototype.resolveWith_ = function(var_args) {
    this.onload_.apply(root, arguments);
  };


  GoogleMapsLoader.prototype.reject_ = function(opt_error) {
    var error = opt_error || new Error('Failed to load Google Maps library.');

    if (this.onload_.onerror) {
      this.onload_.onerror.call(root, error);
    }
    else {
      throw error;
    }
  };


  GoogleMapsLoader.DEFAULT_BASE_URL = 'https://maps.googleapis.com/maps/api/js';


  function NOOP() {
  }


  return googlemapsPlugin;
});
