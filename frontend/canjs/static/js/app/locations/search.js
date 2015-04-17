define(['can', 'app/models/user/user', 'app/models/location/item',
        'googlemaps!', 'refresh',
    'util', 'info', 'i18n', 'jquery', 'jquery.trumbowyg'],
    function (can, User, Location, GoogleMap, Refresh, util, info, i18n, $) {
    'use strict';

  /**
     * Instance of Search Contorllers.
     * @private
     */
    var search, map, markers, markerIcon;
    markerIcon = util.getStaticPath()+'/images/icon/location/cat.png';

    var Search = can.Control.extend({
        init: function () {
            util.logInfo('*search', 'Initialized');
            this.initialize();
        },
        '.search-location click': function () {
            util.logInfo('search-location', 'clicked');
            this.searchAddr();
        },
        '.search-mylocation click': function () {
            $('#term').val('');
            util.logInfo('search-mylocation', 'clicked');
            this.searchMyLocation();
        },
        initialize: function() {
          var mapOptions = {
            zoom: 15,
            mapTypeId: google.maps.MapTypeId.ROADMAP
          };
          map = new google.maps.Map(document.getElementById('map-canvas'),
                mapOptions);
          util.logDebug('map', 'initialized');
        },
        searchMyLocation: function() {
             // Try HTML5 geolocation
          if(navigator.geolocation) {
            navigator.geolocation.getCurrentPosition(function(position) {
              var pos = new google.maps.LatLng(position.coords.latitude,
                                               position.coords.longitude);

//              var infowindow = new google.maps.InfoWindow({
//                map: map,
//                position: pos,
//                content: 'Location found using HTML5.'
//              });
              map.setCenter(pos);
              search.setLatLng();
              search.addMarkers();
            }, function() {
              search.handleNoGeolocation(true);
            });
          } else {
            // Browser doesn't support Geolocation
            search.handleNoGeolocation(false);
          }
        },
        handleNoGeolocation: function(errorFlag) {
         var content = '';
         if (errorFlag) {
            content = i18n.t('location.validation.geolocationFaild');
         } else {
            content = i18n.t('location.validation.geolocationNotSupported');
         }
         var options = {
            map: map,
            position: new google.maps.LatLng(37.5770605,126.9767304),
            content: content
         };
         var infowindow = new google.maps.InfoWindow(options);
         map.setCenter(options.position);
         search.setLatLng();
         util.logDebug('map', 'noGeoLocation');
        },

        showLocation: function(marker){
            util.logDebug('marker',marker);
            // util.replaceHash('location/'+marker.name+'/'+marker.id);
            can.route.attr({'route': 'location/:name/:id', name: util.slashToBlank(util.truncate40(marker.name)), id: marker.id}, true);
//            alert('marker('+marker.id+') - Latitude: ' + marker.position.lat() + '\nLongitude: ' + marker.position.lng());
        },
        searchAddr: function(){
            var $term = $('#term');
            var term = $term.val();
            if(term.length > 0) {
              new google.maps.Geocoder().geocode( { 'address': term}, function(results, status) {
                  if (status == google.maps.GeocoderStatus.OK) {
                      if(results != undefined && results.length > 0) {
                          var location = results[0].geometry.location;
                          search.latitude = location.lat();
                          search.longitude =  location.lng();
                          search.addMarkers();
    //                    if(!marker){
    //                        marker = new google.maps.Marker({
    //                        map: map,
    //                        draggable: true
    //                      });

    //                    google.maps.event.addListener(marker, 'click', function() {
    //                        map.setZoom(8);
    //                        map.setCenter(marker.getPosition());
    //                    });

    //                     google.maps.event.addListener(marker, 'click', search.showLocation());

    //                    marker.setPosition(results[0].geometry.location);
                        map.setCenter(results[0].geometry.location);
                        map.setZoom(15);
                        $term.val(results[0].formattedAddress);
                      } else {
                          alert("Geolocation not found.");
                      }


                  } else {
                    alert("Geocode was not successful for the following reason: " + status);
                  }
              });
            } else {
                util.logWarn('search', 'Search term is empty.');
                search.searchMyLocation();
            }
        },
        setLatLng: function () {
            search.latitude =  map.getCenter().lat();
            search.longitude =  map.getCenter().lng();
        },
        searchLocation: function (locationId) {
            util.logDebug('map data', map);
            Location.findOne({id: locationId}, function (locationData) {
                 var $term = $('#term');
                 $term.val(locationData.location.address);
                 search.latitude =  locationData.location.latitude;
                 search.longitude =  locationData.location.longitude;
                 var pos = new google.maps.LatLng(search.latitude,
                                                  search.longitude);
                 map.setCenter(pos);
                 var center = map.getCenter();
                 util.logDebug('map Center', center);
                 search.addMarkers();
             });
        },
        addMarkers: function () {
            if(markers == undefined) {
               markers = new Array();
            }
            markers.length = 0;
            Location.findAll({page: 1, 'filters': util.stringify({'latitude': search.latitude, 'longitude': search.longitude})},
                function (locationsData) {
                    util.logJson('get Locations',locationsData);
                    var templateData = new can.Observe({
                        locations: locationsData.locations
                    });
                    search.locationData = templateData;
                    $.each(search.locationData.locations, function(idx, location) {
                        markers.push(new google.maps.Marker({
                            map: map,
                            position: new google.maps.LatLng(location.latitude, location.longitude),
//                            draggable: true,
                            icon: markerIcon,
                            name: location.name,
                            id: location.id
                        }));
//                        util.logDebug('locationId', location.id);
                    });
                    $.each(markers, function(idx, marker) {
                        google.maps.event.addListener(marker, 'click', function() {
                            search.showLocation(marker);
                        });
                    });

            });
        },

        load: function () {
            search.initialize();
            search.searchAddr();
        }

    });

    var Router = can.Control.extend({
        defaults: {}
    }, {
        init: function () {
            search = undefined;
            util.logInfo('*search/Search/Router', 'Initialized');
        },
        allocate: function ($app) {
            search = new Search($app);
        },
//        'route': function () {
//            if (can.route.attr('_') == '_') {
//                Refresh.load({'route': 'setting/:tab', 'tab': 'connection'});
//                return false;
//            }
//            this.show();
//        },
        'locations/search/:id route': function (data) {
            this.show(data.id);
        },
        'locations/search route': function (data) {
          this.show();
        },
        show: function (id) {
            var router = this;
            User.findAll({}, function (data) {
                util.logJson('users', data.users);
                var $app = util.getFreshApp();
                $($app).html(can.view('views_location_search_stache', {
                    users: data.users
                }));
               util.refreshTitle();
               router.allocate($app);
               if(id != undefined) {
                   search.searchLocation(id);
               } else {
                   search.searchMyLocation();
               }
            });
        }
    });

    return Router;
});
