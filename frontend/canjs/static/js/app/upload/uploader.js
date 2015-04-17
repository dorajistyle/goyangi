define(['can/component', 'can/model', 'can/util/fixture', 'can',
        'app/models/upload/upload', 'app/models/upload/file', 'app/models/upload/files', 'app/models/user/user', 'app/models/user/user-current',
        'app/models/api-path', 'util', 'i18n', 'settings', 'jquery', 'can/construct/proxy'],
    function (Component, Model, fixture, can, Upload, File, Files, User, UserCurrent, API, util, i18n, settings, $) {
    'use strict';

    /**
     * Instance of User Contorllers.
     * @private
     */
    var uploader;

    /**
     * Control for Uploader
     * @private
     * @author dorajistyle
     * @param {string} target
     * @name users#Uploader
     * @constructor
     */
    var Uploader = can.Control.extend({
        init: function () {
            util.logInfo('*User/Uploader', 'Initialized');
            this.loadForm();
        },
        load: function () {
          UserCurrent.findOne({}, function (currentUserData) {
            util.logInfo('*Upload/Uploader/Router', 'Current user loaded');
            if(currentUserData.user == undefined) {
              can.route.attr({'route': 'login'}, true);
            }  else {
              User.findOne({id: currentUserData.user.id}, function (userdata) {
                uploader.user = userdata;
                uploader.show();
                util.refreshTitle();
              });

            }
          }, function(xhr) {
            can.route.attr({'route': ''}, true);
          });

        },

        show: function () {
            util.logJson('*Upload/Uploader.show data', this.user);
            this.element.html(can.view('views_upload_form_stache', {
                userData: this.user
            }));
        },
        '.perform-follow click': function (el, ev) {
            ev.preventDefault();
            ev.stopPropagation();
            return false;
        },

        '.perform-unfollow click': function (el, ev) {
            ev.preventDefault();
            ev.stopPropagation();
            return false;
        },
        '.show-more-following click': function (el, ev) {
            ev.preventDefault();
            ev.stopPropagation();
            return false;
        },
        '.show-more-followers click': function (el, ev) {
            ev.preventDefault();
            ev.stopPropagation();
            return false;
        },
        loadFixture: function () {
          // Create fixtures for the file metadata for testing
          var id = 0;
          fixture.delay = 2000;
          util.logDebug('id before', id);
          fixture('POST /api/files', function(req){
            util.logDebug('POST /api/files', 'start');
            util.logDebug('id before', id);
            if(id == undefined) {
              var id = 0;
            }
            id = id++;
            var	data = req.data;

            data.isProcessed = true;
            data.id = id;
            util.logDebug('id', id);
            util.logDebug('data', data);
            return data;
          });

          fixture('DELETE /api/files/{id}', function(){
            util.logDebug('DELETE /api/files{id}', 'start');
            return {}
          });
        },

        loadForm: function () {
            Component.extend({
                tag : 'rt-uploader',
                template : can.view('views_upload_init_stache'),
                scope : {
                    init : function(){
                        // Initialize files as an empty array
                        this.attr('files', []);
                    },
                    // When user clicks the remove file button,
                    // just remove it from the list, we'll handle
                    // everything else in the control (`events` object)
                    removeUpload : function(file, ev, el){
                        var idx = this.files.indexOf(file);
                        if(confirm('Are you sure?')){
                            this.files.splice(idx, 1);
                        };
                    }
                },
                helpers : {
                    // Pretty format for the file size
                    formatByteSize : function(size){
                        var i = -1;
                        var byteUnits = ['kB', 'MB', 'GB', ' TB', 'PB', 'EB', 'ZB', 'YB'];

                        size = size();

                        do {
                            size = size / 1024;
                            i++;
                        } while (size > 1024);

                        return Math.max(size, 0.1).toFixed(1) + byteUnits[i];
                    }
                },
                events : {
                    // Initialize the fileupload library when the `rt-uploader` element
                    // is inserted in the page
                    //
                    // `rt-uploader` will be used as a target for the drag and drop of files
                    " inserted" : function(){
                        this.element.find('.select-files').fileupload({
                            singleFileUploads: false,
                            dropZone : this.element,
                            formData : {},
        //					limitMultiFileUploads : 1,
                            url : API+'/upload/images',
                            add : this.proxy('fileUploadAdd')
                        });
                    },
                    // When file is uploaded create a new instance of the `File` and
                    // populate it with the data provided by the fileupload lib
                    //
                    // `uploadId` is used to so we can grab the uploader and abort
                    // the file upload
                    //
                    // we save the function that aborts the upload in the object so
                    // we can grab it later
                    fileUploadAdd : function(ev, data){
                        util.logDebug("fileUploadAdd files length", data.files.length);
                        var uploaders = new Array();
                        for(var i=0;i<data.files.length;i++) {
                            var uploader = new File({
                                name: data.files[i].name,
                                size: data.files[i].size,
                                progress: 0,
                                uploadId: (new Date()).getTime()
                            }), jqXHR;
                            uploaders.push(uploader);
                        }
                            data.uploaders = uploaders;
                            jqXHR = data.submit();
                        for(var j=0;j<data.uploaders.length;j++) {
                            this._uploads = this._uploads || {};
                            this._uploads[uploaders[j].uploadId] = $.proxy(jqXHR.abort, jqXHR);
//                            this._uploads[uploader.uploadId] = $.proxy(jqXHR.abort, jqXHR);

                            this.scope.files.push(uploaders[j]);
//                            this.scope.files.push(uploader);
                        }

                    },
                    // listen to the `fileuploadprogress` event and update the progress
                    ".select-files fileuploadprogress" : function(el, ev, data){
                        for(var i=0;i<data.uploaders.length;i++) {
                            data.uploaders[i].attr('progress', parseInt(data.loaded / data.total * 100, 10));
                        }

                    },
                    // when file is uploaded we will save metadata to a different location
                    // this way we can upload the file to a different server (eg. S3)
                    ".select-files fileuploaddone" : function(el, ev, data){
                        if(settings.savefilesMeta) {
                          var files = new Array();
                          for(var i=0;i<data.uploaders.length;i++) {
                              var file = data.uploaders[i],
                              uploadId = file.attr('uploadId');

                              file.attr({
                                  progress : 100,
                                  // url : data.location
                              });
                              file.removeAttr('uploadId');
                              if(settings.saveAllfilesMetaAtOnce){
                                files.push({name: file.name, progress: file.progress, size: file.size});
                              } else {
                                can.when(file.save()).then(function(data){
                                  util.logDebug('saved data', data);
                                });
                              }
                              delete this._uploads[uploadId];

                          }

                          if(settings.saveAllfilesMetaAtOnce) {
                            var data = {};
                            data.files = files;
                            util.logDebug("data",data);
                            can.when(Files.create({data: JSON.stringify(data)})).then(function(data){
                              util.logDebug('saved data', data);
                            });
                          }
                          delete data.uploaders;
                        }
                    },

                    // when file is removed from the list (user clicked on the remove file button)
                    // we iterate through the list of the removed files and check if it has the `uploadId`
                    // attribute. If it has, it means that file is still uploading and we cancel the upload
                    // otherwise we delete the file
                    "{scope.files} remove" : function(files, ev, removed, prop){
                        var self = this;
                        if(typeof prop === 'number'){
                            can.each(removed, function(file){
                                var uploadId = file.attr('uploadId');
                                if(self._uploads[uploadId]){
                                    self._uploads[uploadId]();
                                    delete self._uploads[uploadId];
                                } else {
                                    util.logJson('upload Item is undefined. file : ', file);
                                    file.destroy();
                                }
                            });
                        }
                    },
                    // We prevent the default behavior for the `dragover` and `drop` events on the
                    // document since drag and drop is handled by the file upload component
                    "{document} drop" : function(el, ev){
                        ev.preventDefault();
                    },
                    "{document} dragover" : function(el, ev){
                        ev.preventDefault();
                    }
                }
            });
        }

    });

    /**
     * Router for upload.
     * @author dorajistyle
     * @param {string} target
     * @function Router
     * @name Upload#Uploader/Router
     * @constructor
     */
    var Router = can.Control.extend({
        defaults: {}
    }, {
        init: function () {
            util.logInfo('*Upload/Uploader/Router', 'Initialized');
        },
        allocate: function () {
          var $app = util.getFreshApp();
          uploader = new Uploader($app);
        },

        'upload route': function () {
            this.allocate();
            uploader.load();
        }
    });

    return Router;
});
