define(['can', 'app/articles/list',
    'util', 'i18n', 'jquery'],
    function (can, List, util, i18n, $) {
    'use strict';


    /**
     * Instance of List Contorllers.
     * @private
     */
    var list, overlay;

    /**
     * Router for ArticleList.
     * @author dorajiarticle
     * @param {string} target
     * @function Router
     * @name article#ArticleListRouter
     * @constructor
     */
    var Router = can.Control.extend({
        defaults: {}
    }, {
        init: function (target) {
            util.logInfo('*Articles/Router', 'Initialized');
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
            list.filterList.articlePerPage=1;
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
//        'profile/:userId/article route': function (data) {
//            this.filterUserId(data);
//        },
        'articles route': function () {
            this.filterNone();
        },
        'notice route': function () {
            this.filterCategory(100);
        },
        'general route': function () {
            this.filterCategory(200);
        },
        'etc route': function () {
            this.filterCategory(300);
        },
        'noticeOne route': function () {
            this.filterCategoryOne(100);
        },
        'articles/user/:username route': function (data) {
            this.filterUserName(data);
        },
        'articles/tag/:name route': function (data) {
            this.filterTag(data);
        }
    });
    return Router;
});
