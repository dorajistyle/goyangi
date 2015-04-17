/*global define*/
define(['can', 'app/models/model', 'app/models/api-path'], function (can, Model, API) {
    'use strict';

    /**
     * Upload
     * @author dorajiupload
     * @namespace upload
     */

    /**
     * Upload model.
     *  @constructor
     * @type {*}
     * @name blog#Upload
     */
    var Upload = Model({
        create: 'POST '+API+'/upload/images',
        update: 'PUT '+API+'/upload/{id}',
        destroy: 'DELETE '+API+'/upload/{id}',
        findOne: 'GET '+API+'/upload/{id}',
        findAll: 'GET '+API+'/upload'
    }, {
    });
    return Upload;
});
