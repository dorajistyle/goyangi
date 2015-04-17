/*global define*/
define(['can', 'app/models/model', 'app/models/api-path'], function (can, Model, API) {
    'use strict';

    /**
     * User related
     * @author dorajistyle
     * @namespace user-activate
     */

    /**
     * UserActivate model.
     *  @constructor
     * @type {*}
     * @name user-activate#UserActivate
     */
    var UserActivate = Model({
        update: 'PUT '+API+'/user/activate/{id}'
    }, {
    });

    return UserActivate;
});
