define(['can', 'app/upload/uploader'], function (can, Uploader) {
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
        new Uploader(target);
    };


    return Routers;
});
