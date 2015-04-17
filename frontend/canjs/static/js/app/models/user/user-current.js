/*global define*/
define(['can', 'app/models/model', 'app/models/api-path'], function (can, Model, API) {
    'use strict';

    /**
     * User related
     * @author dorajistyle
     * @namespace user-current
     */

    /**
     * UserCurrent model.
     *  @constructor
     * @type {*}
     * @name user-current#UserCurrent
     */
    var UserCurrent = Model({
        findOne: 'GET '+API+'/user/current'
    }, {
    });

    return UserCurrent;
});
