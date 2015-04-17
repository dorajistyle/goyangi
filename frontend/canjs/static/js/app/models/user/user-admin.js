/*global define*/
define(['can', 'app/models/model', 'app/models/api-path'], function (can, Model, API) {
    'use strict';

    /**
     * UserAdmin related
     * @author dorajistyle
     * @namespace users
     */

    /**
     * UserAdmin model.
     * @constructor
     * @type {*}
     * @name users#UserAdmin
     */
    var UserAdmin = Model({
        findAll: 'GET '+API+'/user/admin',
        findOne: 'GET '+API+'/user/admin/{id}'
    }, {
    });
    return UserAdmin;
});
