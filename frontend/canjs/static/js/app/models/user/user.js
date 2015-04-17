/*global define*/
define(['can', 'app/models/model', 'app/models/api-path'], function (can, Model, API) {
    'use strict';

    /**
     * User related
     * @author dorajistyle
     * @namespace users
     */

    /**
     * User model.
     * @constructor
     * @type {*}
     * @name users#User
     */
    var User = Model({
        findAll: 'GET '+API+'/users',
        findOne: 'GET '+API+'/users/{id}',
        create: 'POST '+API+'/users',
        // update: 'PUT '+API+'/users/{id}',
        update : function(params){
          console.log("params",params);
          return $.ajax({
            type: 'PUT',
            url: API+'/users/'+params.id+'?type='+params.type,
            data: params.data
          });
        },
        destroy: 'DELETE '+API+'/users/{id}'
    }, {
    });
    return User;
});
