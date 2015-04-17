define(['can', 'app/admin/admin'], function (can, Admin) {
    'use strict';

    /**
     * Router for users.
     * @author dorajistyle
     * @param {string} target
     * @function Router
     * @name admin#Routers
     * @constructor
     */

    var Routers = function (target) {
        new Admin(target);
    };


    return Routers;
});
