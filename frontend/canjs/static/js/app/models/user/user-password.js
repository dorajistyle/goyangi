/*global define*/
define(['can', 'app/models/model', 'app/models/api-path'], function (can, Model, API) {
    'use strict';

    /**
     * User related
     * @author dorajistyle
     * @namespace user-password
     */

    /**
     * UserPassword model.
     *  @constructor
     * @type {*}
     * @name user-password#UserPassword
     */
    var UserPassword = Model({
        create: 'POST '+API+'/user/send/password/reset/token',
        update : function(data){
          return $.ajax({
            type: 'PUT',
            url: API+'/user/reset/password',
            data: data
          });
        },
    }, {
    });

    return UserPassword;
});
