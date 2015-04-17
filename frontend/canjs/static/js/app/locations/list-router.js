define(['can', 'app/locations/list',
    'util', 'i18n', 'jquery'],
    function (can, List, util, i18n, $) {
    'use strict';


    /**
     * Instance of List Contorllers.
     * @private
     */
    var list, overlay;

    /**
     * Router for profile.
     * @author dorajilocation
     * @param {string} target
     * @function Router
     * @name users#LocationListRouter
     * @constructor
     */
    var Router = can.Control.extend({
        defaults: {}
    }, {
        init: function (target) {
            util.logInfo('*Locations/Router', 'Initialized');
            this.target = target;
        },
        allocate: function () {
            var $app = util.getFreshApp();
            list = new List($app);
        },
        initFilter: function () {
            if(list != undefined && list.id != 'app'){
                list = undefined;
            }
            util.logDebug('filter',list);
            util.allocate(this, list);
        },
        filterNone: function () {
            this.initFilter();
            list.load();
        },
        filterCategory: function (categoryId) {
            this.initFilter();
            list.filterList.categories=[categoryId];
            list.load();
        },
        filterCategoryOne: function (categoryId) {
            this.initFilter();
            list.filterList.categories=[categoryId];
            list.filterList.locationPerPage=1;
            list.load();
        },
        filterUserName: function (data) {
            this.initFilter();
            list.filterList.username=data.username;
            list.load();
        },
        filterTag: function (data) {
            this.initFilter();
            list.filterList.tags = data.name;
            list.load();
        },
//        'profile/:userId/location route': function (data) {
//            this.filterUserId(data);
//        },
        'locations route': function () {
            this.filterNone();
        },
        'locations/user/:username route': function (data) {
            this.filterUserName(data);
        },
        'locations/tag/:name route': function (data) {
            this.filterTag(data);
        }
    });
    return Router;
});
