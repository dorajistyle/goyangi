define(['can', 'app/articles/list-router', 'app/articles/item', 'app/articles/write'],
    function (can, List, Item, Write) {
    'use strict';

    /**
     * Router for Article.
     * @author dorajiarticle
     * @param {string} target
     * @function Router
     * @name article#Routers
     * @constructor
     */

    var Routers = function (target) {
        new List(target);
        new Item(target);
        new Write(target);
    };

    return Routers;
});
