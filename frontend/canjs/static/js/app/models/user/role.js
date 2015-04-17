/*global define*/
define(['can', 'app/models/model', 'app/models/api-path'], function (can, Model, API) {
    'use strict';

    /**
     * Admin related
     * @author dorajistyle
     * @namespace admin
     */

    /**
     * Role model.
     *  @constructor
     * @type {*}
     * @name admin#Role
     */
    var Role = Model({
        create: 'POST '+API+'/roles',
        update: 'PUT '+API+'/roles/{id}',
        destroy: 'DELETE '+API+'/roles/{id}',
        findOne: 'GET '+API+'/roles/{id}',
        findAll: 'GET '+API+'/roles'
    }, {
    });
    return Role;
});
