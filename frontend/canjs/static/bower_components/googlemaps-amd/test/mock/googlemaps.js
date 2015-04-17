define(function() {
  var root = this;

  /**
   * Mock implementation of google.maps
   *
   * @constructor
   */
  var MockGoogleMaps = function() {
    this.googleOrig_ = root.google;
  };


  /**
   * Expose in place of the global google.maps library.
   */
  MockGoogleMaps.prototype.exposeMockLibrary = function() {
    root.google = {
      maps: this
    };
  };


  /**
   * Restore the original reference to the
   * global google object.
   */
  MockGoogleMaps.prototype.restore = function() {
    root.google = this.googleOrig_;
  };


  /**
   * Set the global google object to null.
   */
  MockGoogleMaps.prototype.removeGlobalLibrary = function() {
    root.google = null;
  };


  return MockGoogleMaps;
});
