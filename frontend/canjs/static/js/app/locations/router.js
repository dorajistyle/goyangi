define(['can', 'app/locations/list-router', 'app/locations/item', 'app/locations/write', 'app/locations/search'],
    function (can, List, Item, Write, Search) {
    'use strict';

    /**
     * Router for upload.
     * @author dorajilocation
     * @param {string} target
     * @function Router
     * @name upload#Routers
     * @constructor
     */

    var Routers = function (target) {
        new List(target);
        new Item(target);
        new Write(target);
        new Search(target);
    };

    return Routers;
});
