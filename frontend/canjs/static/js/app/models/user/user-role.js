/*global define*/
define(['can', 'app/models/model', 'app/models/api-path'], function (can, Model, API) {
    'use strict';

    /**
     * Role model.
     * @constructor
     * @type {*}
     * @name admin#UserRole
     */
    var UserRole = Model({
        create: 'POST '+API+'/users/roles',
        destroy : function(userId,roleId){
          return $.ajax({
              type: 'DELETE',
              url: API+'/users/'+userId+'/roles/'+roleId,
              data: {}
          });
        }
    }, {
    });
    return UserRole;
});
