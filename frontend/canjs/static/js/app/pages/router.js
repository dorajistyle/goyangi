define(['can', 'app/pages/how-it-works'],
    function (can, HowItWorks) {
    'use strict';

    /**
     * Router for pages.
     * @author dorajistyle
     * @param {string} target
     * @function Router
     * @name pages#Routers
     * @constructor
     */

    var Routers = function (target) {
        new HowItWorks();
    };
    return Routers;
});
