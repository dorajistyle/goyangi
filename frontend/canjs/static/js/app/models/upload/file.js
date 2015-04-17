/*global define*/
define(['can', 'app/models/model', 'app/models/api-path'], function (can, Model, API) {
    'use strict';

    /**
     * File
     * @author dorajiupload
     * @namespace upload
     */

    /**
     * File model.
     *  @constructor
     * @type {*}
     * @name blog#File
     */
    var File = Model({
        create: 'POST '+API+'/upload/files',
        destroy: 'DELETE '+API+'/upload/files/{id}'
    }, {
    });
    return File;
});
