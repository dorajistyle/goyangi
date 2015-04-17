/*global define*/
define(['can', 'app/models/model', 'app/models/api-path'], function (can, Model, API) {
    'use strict';

    /**
     * Authentication
     * @author dorajistyle
     * @namespace authentications
     */

    /**
     * Authentication model.
     *  @constructor
     * @type {*}
     * @name authentications#Authentication
     */
    var Authentication = Model({
        findAll: 'GET '+API+'/authentications',
        create: 'POST '+API+'/authentications',
        destroy: 'DELETE '+API+'/authentications'
    }, {
    });
    return Authentication;
});
