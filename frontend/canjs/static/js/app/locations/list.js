define(['can', 'app/models/location/item', 'util', 'i18n', 'jquery'],
    function (can, Location, util, i18n, $) {
    'use strict';

    var List = can.Control.extend({
        init: function ($wrapper) {
            util.logInfo('*location/List', 'Initialized');
            this.id = $wrapper.attr('id');
            this.filterList = {};
            this.filterList.currentPage =1;
        },
        load: function () {
             util.logDebug('*Location/List load Filter', 'Before load');
             var list = this;
             util.logDebug('*Location/List load Filter', 'Already loaded');
             list.loadList();
        },
        loadList: function () {
            var list = this;
            Location.findAll({'filter': util.stringify(list.filterList)},
                function (locationsData) {
                    util.logJson('List',locationsData);
                    list.setListData(locationsData);
                }
            );
        },

        setListData: function (locationsData) {
            var list = this;
            var templateData = new can.Observe({
                locations: locationsData.locations,
                hasPrev: locationsData.hasPrev,
                hasNext: locationsData.hasNext,
                prevPage: locationsData.currentPage-1,
                nextPage: locationsData.currentPage+1,
                category: locationsData.category,
                canWrite: locationsData.canWrite,
                staticPath: util.getStaticPath()
            });
            list.locationData = templateData;
            list.show();
        },
//        '{window} scroll': function() {
//            if(util.distanceToBottom() < 50) {
//                util.logDebug('Location List', 'scrolled down');
//                this.nextPage();
//            }
//        },
        '.more click': function() {
            var list = this;
            list.nextPage();
        },

//        scroll: function() {
//            if(util.distanceToBottom() < 50) {
//                this.nextPage();
//            }
//        },
        nextPage: function() {
            var list = this;
            if(list.locationData && list.locationData.hasNext &&  list.filterList.currentPage != list.locationData.nextPage) {
                list.filterList.currentPage = list.locationData.nextPage;
                util.logDebug("list.filterList ", list.filterList);
                Location.findAll({'filter': util.stringify(list.filterList)}, function (locationsData) {
                    util.logJson('List reload', locationsData.locations);
                    util.concatArray(list.locationData.locations, locationsData.locations);
                    util.refreshPaginate(list.locationData, locationsData);
                    }
                );
            }
        },
        // filter: function(formName) {
        //     var list = this;
        //     var categoryIds = [];
        //     var form = list.element.find('#'+formName);
        //     var filterBtn = list.element.find('.get-filtered-product-list');
        //     var $categories = form.find('input[name=category]:checked');
        //     $.each($categories, function(i, selectedCategory) {
        //        var id = $(selectedCategory).val();
        //        categoryIds.push(id);
        //     });
        //     var fl = list.filterList;
        //     fl.categories = categoryIds;
        //     if(fl.currentPage == undefined || fl.currentPage == 0) {
        //       fl.currentPage = 1;
        //     }
        //     list.filterList = fl;
        //     filterBtn.attr('disabled', 'disabled');
        //     list.load();
        // },
        // '.apply-filter click': function(el, ev) {
        //     var list = this;
        //     ev.preventDefault();
        //     ev.stopPropagation();
        //     list.filter('filterForm');
        // },
        // '.sort-new click': function (el, ev) {
        //     var list = this;
        //     list.load();
        // }
        /**
         * Show location view.
         * @memberof location#List
         */
        show: function () {
            var list = this;
            list.element.html(can.view('views_location_list_stache', list.locationData));
        }
    });
    return List;
});
