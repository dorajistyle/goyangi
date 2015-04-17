define(['can', 'app/models/location/item', 'app/partial/comment', 'app/models/location/comment', 'app/partial/destroy',
        'app/partial/liking', 'app/models/location/liking', 'app/partial/sharing',
        'facebook', 'settings', 'util', 'validation', 'i18n', 'jquery', 'jquery.fileupload', 'jquery.magnific-popup'],
    function (can, Location, Comment, CommentModel, Destroy, Liking, LikingModel, Sharing, FB, settings, util, validation, i18n, $) {
    'use strict';


    /**
     * Instance of ShowLocation Contorllers.
     * @private
     */
    var showLocation, updateLocation;
    var modalDestroyId =  "destroyLocationConfirm";
    var modalDestroyName =  "location";
    /**
     * Control for show Location
     * @private
     * @author dorajistyle
     * @param {string} target
     * @name location#ShowLocation
     * @constructor
     */

    var ShowLocation = can.Control.extend({
        init: function () {
            util.logInfo('*share/ShowLocation', 'Initialized');
            // this.maxLocations = 0;        // This indicates how many locations user actually uploaded
            // this.locationNames = [];      // This is an array that holds the uploaded location target names
        },
        load: function (locationId) {
            var locationItem = this;
            util.logInfo('*share/ShowLocation', 'loaded');
            Location.findOne({id: locationId}, function (locationData) {
                locationItem.locationId = locationId;
                locationItem.loadLocationData(locationData);
                locationItem.element.hammer({drag: false, hold: true, transform: false,tapAlways: false, swipeVelocityX: 0.3});
             });
        },
        loadLocationData: function(locationData) {
            var locationItem = this;
            util.logJson('ShowLocation locationData',locationData);
            var templateData = new can.Observe({
                location: locationData.location,
                isAuthor: locationData.isAuthor,
//                isLiked: locationData.isLiked,
                currentUserId: locationData.currentUserId,
//                likingCount: locationData.likingCount,
                staticPath: util.getStaticPath(),
                modalDestroyId: modalDestroyId,
                modalDestroyName: modalDestroyName,
                parentId: locationData.location.id,
                parentTitle: util.truncate(locationData.location.name),
                parentDescription: util.truncate(locationData.location.content),
                parentTagNames: locationData.location.tagNames,
                commentList: locationData.location.commentList,
                likingList: locationData.location.likingList
            });
            locationItem.locationData = templateData;
//            var descriptionAndTags = util.linkHashtag('#!locations/tag', locationData.location.description);
//            locationItem.locationData.location.attr('description',descriptionAndTags.text);
//            locationItem.locationData.location.attr('hashtag',descriptionAndTags.hashtag);
//            locationItem.locationData.location.attr('tags',descriptionAndTags.tags);


            can.when(locationItem.show()).then(function () {
                       util.refreshMeta();
//                       locationItem.initFacebook();
//                    $(document).ready(function(){
                       locationItem.initControl();
                        // Hide the arrows if there's only one image
//                   });

                    /*$('.preview-wrapper').magnificPopup({
                      delegate: 'a', // child locations selector, by clicking on it popup will open
                      type: 'image',
                      closeOnContentClick: true,
                      gallery: {enabled:true},
                      image: {
                        verticalFit: true,
                        tError: '<a href="%url%">'+i18n.t('magnificPopup.error.theImage')+'</a> '+i18n.t('magnificPopup.error.couldNotBeLoaded')
                      }
                      // other options
                    });*/
                }
            );

        },
        'tab': function (el, ev) {
            util.logDebug('location', 'tab');
        },
        'hold': function (el, ev) {
            util.logDebug('location', 'hold');
        },

        'swipeleft': function (el, ev) {
            util.logDebug('location', 'swipeleft');
            this.element.find('.right').click();
        },
        'swiperight': function (el, ev) {
            util.logDebug('location', 'swiperight');
            this.element.find('.left').click();
        },


        initFacebook: function () {
            var facebookScript = document.createElement('script');
            facebookScript.type = 'text/javascript';
            facebookScript.text = '(function(d, s, id) {var js, fjs = d.getElementsByTagName(s)[0];if (d.getElementById(id)) return;js = d.createElement(s); js.id = id;js.src = "//connect.facebook.net/ko_KR/sdk.js#xfbml=1&appId='+settings.facebookAppId+'&version=v2.0";fjs.parentNode.insertBefore(js, fjs);}(document, "script", "facebook-jssdk"));';
            util.logDebug('tourl', facebookScript.text);
            $('body').append(facebookScript);

//            var url = util.getCurrentURL().replace('#!','');;
//            $('#likeLocation').html("<fb:like href='" + url + "' send='false' layout='button_count' width='20' show_faces='false'></fb:like>");
            FB.XFBML.parse();
//            FB.getLoginStatus(function(response) {
//                if (response.status === 'connected') {
//                    // the user is logged in and has authenticated your
//                    // app, and response.authResponse supplies
//                    // the user's ID, a valid access token, a signed
//                    // request, and the time the access token
//                    // and signed request each expire
//                    var uid = response.authResponse.userID;
//                    var accessToken = response.authResponse.accessToken;
//                    var likeStatus = util.fetchFBLikingCount(url);
//                    if(likeStatus == 1) {
//                        $('#fbLikeBtn').addClass('btn-warning');
//                    }
//                }
//            });
//            util.fetchFBLikingCount(url).always(function(res){
//                $('#fbCount').text(res);
//            });


        },
        initControl: function (){

            var $locationWrapper = $(this.element.find('#locationWrapper'));
            var $headerDiv = $(this.element.find('#locationHeader'));
            var $commentDiv = $(this.element.find('#commentWrapper'));
            updateLocation = new UpdateLocation($headerDiv);
            new Liking($headerDiv, {parentId: showLocation.locationId, parentName: i18n.t("location.modelName"), view: showLocation, model: LikingModel, likePrefix: "liking.like", unlikePrefix: "liking.unlike"});
            new Destroy($headerDiv, {modalDestroyId: modalDestroyId, name: modalDestroyName, view: showLocation, model: Location});
            new Comment($commentDiv, {parentId: showLocation.locationId, parentData: showLocation.locationData.location, view: showLocation, model: CommentModel});
            new Sharing($locationWrapper);
            // destroyLocation = new DestroyLocation($topDiv);
//            FB.XFBML.parse(document.getElementById("locationWrapper"));
//            FB.XFBML.parse();
//            FB.Event.subscribe('edge.create', this.pageLikeOrUnlikeCallback);
//            FB.Event.subscribe('edge.remove', this.pageLikeOrUnlikeCallback);
            // createComment = new CreateComment($commentDiv);
            // destroyComment = new DestroyComment($commentDiv);
            // showComment = new ShowComment($commentDiv);
        },

        initGestures: function (){

        },
        pageLikeOrUnlikeCallback: function(url, htmlElement) {
          util.fetchFBLikingCount(url).always(function(res){
                $('#fbCount').text(res);
            });

          util.logDebug("pageLikeOrUnlikeCallback", 'yeah!');
          util.logDebug('callback url', url);
          util.logDebug('callback elem', htmlElement);
        },

// In your onload handler



        loadImage: function (){
             var locationItem = this;
        },
        reload: function () {
          can.route.attr({'route': 'locations'}, true);
        },
        /**
         * @memberof share#ShowLocation
         */
        show: function () {
            this.element.html(can.view('views_location_item_stache', this.locationData));
        }

    });

    var UpdateLocation = can.Control.extend({
      init: function () {
        util.logInfo('*location/UpdateLocation', 'Initialized');
      },
      '.update-location click': function (el, ev) {
        ev.preventDefault();
        ev.stopPropagation();
        updateLocation.id = util.getId(ev);
        this.performUpdate();
        return false;
      },
      /**
      * perform Update action.
      * @memberof location#DestroyLocation
      */
      performUpdate: function () {
        // if(updateLocation.id == undefined) {
        //   updateLocation.id = -1;
        // }
        can.route.attr({'route': 'locations/write/:id', 'id': updateLocation.id}, true);
      }

    });

    /**
     * Router for profile.
     * @author dorajistyle
     * @param {string} target
     * @function Router
     * @name location#ShowLocation
     * @constructor
     */
    var Router = can.Control.extend({
        defaults: {}
    }, {
        init: function () {
            util.logInfo('*location/ShowLocation/Router', 'Initialized');
        },
        allocate: function () {
            var $app = util.getFreshApp();
            showLocation = new ShowLocation($app);
            // destroyLocation = new DestroyLocation($app);
            // new Destroy($app, {modalDestroyId: modalDestroyId, name: modalDestroyName, view: showLocation, model: Location});
            // likeLocation = new LikingLocation($app);
        },
        show: function(locationId) {
            util.allocate(this, showLocation);
            showLocation.load(locationId);

        },
        'location/:name/:id route': function (data) {
            this.show(data.id);
        }
    });
    return Router;
});
