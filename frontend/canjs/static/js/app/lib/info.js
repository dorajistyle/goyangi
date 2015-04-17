define(['i18n', 'text!app/preload/article-categories.json'],
    function (i18n, articleCategories) {
    'use strict';
    var articleCategoryList;
    articleCategoryList = [];

    return {


        /**
         * Get Count Array for mustache iteration.
         * @param count
         * @returns {Array}
         */
        getCountArray: function (count) {
            var max = 10;
            var limit = count < max ? count : max;
            var arr=[];
            for(var i = 1; i<=limit; i++) {arr.push({'count': i});}
            return arr;
        },

         /**
         * get a list of article categories.
         * @returns {Array}
         */
        getArticleCategories: function () {
            if(articleCategoryList.length == 0) {
                articleCategoryList = $.parseJSON(articleCategories);
                articleCategoryList.categoriesIdx = {};
                $.each(articleCategoryList, function(key, articleCategory){
                   switch(key) {
                       case 'categories':
                           $.each(articleCategory, function(idx, category) {
                               articleCategoryList.categoriesIdx[category.id]= idx;
                           });
                           break;
                   }
                });

            }
            return articleCategoryList;
        },
          /**
         * Get name from filter.
         * @param status
         * @returns {r.t|*|u.t|t|L.t}
         */
        getArticleCategoryName: function (kind, id) {
            var articleCategories = this.getArticleCategories();
            var name='';
            switch(kind) {
                case 'category':
                    var category = articleCategories.categories[articleCategories.categoriesIdx[id]];
                    if(category == undefined) {return false;}
                    name = category.name;
                    break;
            }
            return i18n.t('article.'+kind+'.'+name);
        },

          /**
         * Get route from filter.
         * @param status
         * @returns {r.t|*|u.t|t|L.t}
         */
        getArticleCategoryRoute: function (id) {
            var articleCategories = this.getArticleCategories();
            var name='';
            var category = articleCategories.categories[articleCategories.categoriesIdx[id]];
            if(category == undefined) {return false;}
            name = category.name;

            return name;
        }
    }
});
