define(['can', 'app/setting/setting'], function (can, Setting) {
    'use strict';

    /**
     * Router for users.
     * @author dorajistyle
     * @param {string} target
     * @function Router
     * @name users#Routers
     * @constructor
     */

    var Routers = function (target) {
        new Setting(target);
    };


    return Routers;
});
