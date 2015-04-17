/*global define*/
define(['can', 'app/models/model', 'app/models/api-path'], function (can, Model, API) {
    'use strict';

    /**
     * Files
     * @author dorajiupload
     * @namespace upload
     */

    /**
     * Files model.
     *  @constructor
     * @type {*}
     * @name blog#Files
     */
    var Files = Model({
        create : function(params){
          console.log("params.data : "+params.data);
          return $.ajax({
            type: 'POST',
            url: API+'/upload/files/all',
            contentType: "application/json",
            async: false,
            data: params.data
            // data: '{"files":{id: 3, name: "testName", progress: 77, size: 1312981}}',
            // data: params.data
            // data: {files: [{id: 3, name: "testName", progress: 77, size: 1312981}]}
            // data: 'files={id: 3, name: "testName", progress: 77, size: 1312981}'
            // dataType:"json"
          });
        }
    }, {
    });
    return Files;
});
