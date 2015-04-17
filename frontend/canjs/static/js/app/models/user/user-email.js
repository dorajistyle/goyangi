/*global define*/
define(['can', 'app/models/model', 'app/models/api-path'], function (can, Model, API) {
    'use strict';

    /**
     * User related
     * @author dorajistyle
     * @namespace user-email
     */

    /**
     * UserEmail model.
     *  @constructor
     * @type {*}
     * @name user-email#UserEmail
     */
    var UserEmail = Model({
        findOne: 'GET '+API+'/user/email/{email}'
    }, {
    });
    return UserEmail;
});
