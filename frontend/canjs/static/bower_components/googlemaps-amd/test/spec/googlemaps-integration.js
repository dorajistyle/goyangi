define([
  'googlemaps!'
], function(googleMaps) {
  var root = this;

  describe('The googlemaps AMD loader plugin (integration)', function() {

    it('should work as a loader plugin', function() {
      expect(root.google).toBeDefined();
      expect(root.google.maps).toBeDefined();
      expect(googleMaps).toEqual(root.google.maps);
    });

  });
});