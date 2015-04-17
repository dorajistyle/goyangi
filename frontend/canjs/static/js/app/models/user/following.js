/*global define*/
define(['can', 'app/models/model', 'app/models/api-path'], function (can, Model, API) {
    'use strict';

    /**
     * User related
     * @author dorajistyle
     * @namespace following
     */

    /**
     * User model.
     *  @constructor
     * @type {*}
     * @name following#Following
     */
    var Following = Model({
        findAll: 'GET '+API+'/following'
    }, {
    });
    return Following;
});
