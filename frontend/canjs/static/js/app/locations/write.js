define(['can', 'app/models/location/item',
    'util', 'validation', 'info', 'i18n', 'jquery', 'jquery.trumbowyg'],
    function (can, Location, util, validation, info, i18n, $) {
    'use strict';


    /**
     * Instance of UploadImage Contorllers.
     * @private
     */
    var writeLocation;

    var WriteLocation = can.Control.extend({
        init: function () {
            util.logInfo('location/WriteLocation', 'Initialized');
        },
        load: function(locationId) {
//             var locationCategories = info.getLocationCategories();

//             util.logJson("location categories", locationCategories);
          if(locationId == undefined || locationId.length <= 0 || locationId <= 0) {
            var templateData = new can.Observe({
              location: null,
            });
            writeLocation.locationData = templateData;
            writeLocation.locationId = locationId;
            writeLocation.show();
          }
          else {
            Location.findOne({id: locationId}, function (locationData) {
              var templateData = new can.Observe({
                location: locationData.location,
              });
              writeLocation.locationData = templateData;
              writeLocation.locationId = locationId;
              writeLocation.show();
              //                writeArticle.initForm();
              //                writeArticle.activeForm();
            }, function (xhr) {

            });
          }
//
//
//              if(locationId == undefined || locationId.length <= 0) { locationId = -1; }
//              Location.findOne({id: locationId}, function (locationData) {
//                 var templateData = new can.Observe({
//                     location: locationData.location,
// //                    categories: locationCategories.categories,
//                     hasPrev: locationData.hasPrev,
//                     hasNext: locationData.hasNext,
//                     currentPage: locationData.currentPage
//                 });
// //                util.logDebug("location data", templateData);
//                 writeLocation.locationData = templateData;
//                 writeLocation.locationId = locationId;
//                 writeLocation.show();
//
//                 $('#locationContent').trumbowyg({
//                     mobile: true,
//                     tablet: true
//                 });
// //                writeLocation.initForm();
// //                writeLocation.activeForm();
//             });
        },
        /**
         * Refresh will be performed after image uploaded.
         */
        refresh: function() {
            var locationId = $('#id').val();
            var locationName = $('#locationName').val();
                if(locationId != undefined && locationId.length > 0 && locationName != undefined && locationName.length > 0) {
                    can.route.attr({'route': 'location/:name/:id', name: util.slashToBlank(util.truncate40(locationName)), id: locationId}, true);
                }
                return true;
            if(writeLocation.locationData.location != undefined) {
                if(locationId != undefined && locationId.length > 0) can.route.attr({'route': 'location/write/:id', id: locationId}, true);
                return false;
            }
            can.route.attr({'route': 'location/write'}, true);
            writeLocation.load();
            return false;
        },
        reload: function() {
            var locationId = $('#id').val();
            if(locationId != undefined && locationId.length > 0) {
                can.route.attr({'route': 'location/write/:id', id: locationId}, true);
                writeLocation.load(locationId);
            }
        },
//        initForm : function () {
//            if(writeLocation.locationData.location != undefined) {
//            }
//        },
//        activeForm: function() {
//            $($('#location-image').find('.location-image-form')).addClass('active');
//        },
//        '.save-location click': function(el, ev){
//            ev.preventDefault();
//            ev.stopPropagation();
//        },
//        loadImage: function() {
//            image.loadEditableImages(writeLocation, image.getLocationImageNames());
//        },

        validate: function () {
             var isValid = validation.minLength('name', 1,  'location.view.write.validation.name', false) &&
             validation.isDecimal('latitude', 'location.view.write.validation.latitude') &&
             validation.isDecimal('longitude','location.view.write.validation.longitude') &&
             validation.minLength('address', 1,  'location.view.write.validation.address', false) &&
                    validation.minLengthTextarea('content', 1, 'location.view.write.validation.content');
//                    validation.minLength('tags', 1, 'location.view.write.validation.tags', false);
             return isValid;
        },
//        validateUpdate: function () {
//             return util.validateCondition($('.collaborate-id:checked').length > 0, 'location.view.write.validation.chooseProduct') &&
//                    validation.minLengthTextarea('description', 1, 'location.view.write.validation.description') &&
//                    validation.maxLengthTextarea('description', 65535, 'location.view.write.validation.descriptionTooLong');
//        },
        '.write-location click': function (el, ev) {
//            util.logDebug('write-location', writeLocation.locationId);
            ev.preventDefault();
            ev.stopPropagation();
            if(writeLocation.locationId == -1) {
                writeLocation.performCreate();
                return false;
            }
            writeLocation.performUpdate();

        },

        performCreate: function () {
            if (this.validate()) {
                var form = this.element.find('#locationWriteForm');
                var values = can.deparam(form.serialize());
                var $form = $(form);
                values.name = validation.replaceBadWords(values.name);
                values.content = validation.replaceBadWords(values.content);
                if (!$form.data('submitted')) {
                    $form.data('submitted', true);
                    var updateBtn = this.element.find('.change-status');
                    updateBtn.attr('disabled', 'disabled');
                    can.when(Location.create(values)).then(function (result) {
                        var id = writeLocation.element.find('#id');
                        var $document = $(document);
                        $document.ajaxError(function() {
                            can.route.attr({'route': ''}, true);
                        });
                        id.val(result.location.id);
                        util.showSuccessMsg(i18n.t('location.view.created.done'));
                        writeLocation.refresh();
                        can.route.attr({'route': 'location/:name/:id', name: result.location.name, id: result.location.id}, true);
                    }, function (xhr) {
                        updateBtn.removeAttr('disabled');
                        $form.data('submitted', false);
                        util.showErrorMsg(i18n.t('location.view.created.fail'));
                    });
                }
            } else {
                util.showMessages();
            }
        },
        performUpdate: function () {
            if (this.validate()) {
                var form = this.element.find('#locationWriteForm');
                var values = can.deparam(form.serialize());
                var $form = $(form);
                values.name = validation.replaceBadWords(values.name);
                values.content = validation.replaceBadWords(values.content);
                 if (!$form.data('submitted')) {
                    $form.data('submitted', true);
                    var updateBtn = this.element.find('.change-status');
                    updateBtn.attr('disabled', 'disabled');
                    can.when(Location.update(writeLocation.locationId, values)).then(function (result) {
                        writeLocation.refresh();
                        util.showSuccessMsg(i18n.t('location.view.updated.done'));
                        $form.data('submitted', false);
                        can.route.attr({'route': 'location/:name/:id', name: values.name, id: writeLocation.locationId}, true);
                    }, function (xhr) {
                        updateBtn.removeAttr('disabled');
                        $form.data('submitted', false);
                        util.handleError(xhr);
                        // util.showErrorMsg(i18n.t('location.view.updated.fail'));
                    });
                }
            } else {
                util.showMessages();
            }
        },

        show: function () {
            this.element.html(can.view('views_location_write_stache', writeLocation.locationData));
            $('#locationContent').trumbowyg({
              mobile: true,
              tablet: true
            });
        }
    });


    /**
     * Router for location image.
     * @author dorajilocation
     * @param {string} target
     * @function Router
     * @name users#LocationWriteRouter
     * @constructor
     */
    var Router = can.Control.extend({
        defaults: {}
    }, {
        init: function () {
            writeLocation = undefined;
            util.logInfo('*location/Router', 'Initialized');
        },
        allocate: function () {
            var $app = util.getFreshApp();
            writeLocation = new WriteLocation($app);
//            destroyImage = new DestroyImage($app);
        },
        load: function (locationId) {
            util.allocate(this, writeLocation);
//            if(writeLocation.locationData === undefined){
            writeLocation.load(locationId);
//                return false;
//            }
//            writeLocation.reload(locationId);
        },
        'locations/write route': function () {
            this.load(-1);
        },
        'locations/write/:id route': function (data) {
            this.load(data.id);
        }
    });
    return Router;
});
