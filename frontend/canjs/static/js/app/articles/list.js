define(['can', 'app/models/article/item', 'util', 'i18n', 'jquery'],
    function (can, Article, util, i18n, $) {
    'use strict';

    var List = can.Control.extend({
        init: function ($wrapper) {
            util.logInfo('*article/List', 'Initialized');
            this.id = $wrapper.attr('id');
            this.filterList = {};
            this.filterList.currentPage =1;
        },
        load: function () {
             util.logDebug('*Article/List load Filter', 'Before load');
             var list = this;
             util.logDebug('*Article/List load Filter', 'Already loaded');
             list.loadList();
        },
        loadList: function () {
            var list = this;
            Article.findAll({'filter': util.stringify(list.filterList)},
                function (articlesData) {
                    util.logJson('List',articlesData);
                    list.setListData(articlesData);
                }
            );
        },

        setListData: function (articlesData) {
            var list = this;
            var templateData = new can.Observe({
                articles: articlesData.articles,
                hasPrev: articlesData.hasPrev,
                hasNext: articlesData.hasNext,
                prevPage: articlesData.currentPage-1,
                nextPage: articlesData.currentPage+1,
                category: articlesData.category,
                canWrite: articlesData.canWrite,
                staticPath: util.getStaticPath()
            });
            list.articleData = templateData;
            list.show();
        },
//        '{window} scroll': function() {
//            if(util.distanceToBottom() < 50) {
//                util.logDebug('Article List', 'scrolled down');
//                this.nextPage();
//            }
//        },
        '.more click': function() {
            var list = this;
            list.nextPage();
        },
        nextPage: function() {
            var list = this;
            if(list.articleData && list.articleData.hasNext &&  list.filterList.currentPage != list.articleData.nextPage) {
                list.filterList.currentPage = list.articleData.nextPage;
                util.logDebug("list.filterList ", list.filterList);
                Article.findAll({'filter': util.stringify(list.filterList)}, function (articlesData) {
                    util.logJson('List reload', articlesData.articles);
                    util.concatArray(list.articleData.articles, articlesData.articles);
                    util.refreshPaginate(list.articleData, articlesData);
                    }
                );
            }
        },
        /**
         * Show article view.
         * @memberof article#List
         */
        show: function () {
            var list = this;
            list.element.html(can.view('views_article_list_stache', list.articleData));
        }
    });
    return List;
});
