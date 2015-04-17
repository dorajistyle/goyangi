define(['can', 'app/models/article/item', 'refresh',
    'util', 'validation',  'info', 'i18n', 'jquery', 'jquery.trumbowyg'],
    function (can, Article, Refresh, util, validation, info, i18n, $) {
    'use strict';


    /**
     * Instance of wrtieArticle Contorllers.
     * @private
     */
    var writeArticle;

    var WriteArticle = can.Control.extend({
        init: function () {
            util.logInfo('*article/WriteArticle', 'Initialized');
        },
        load: function(articleId, categoryId) {
//             var article-categories = info.getArticleCategories();
//             util.logJson("article categories", article-categories);
             if(articleId == undefined || articleId.length <= 0 || articleId <= 0) {
               var templateData = new can.Observe({
                 article: null,
                 categoryId : categoryId,
               });
               writeArticle.articleData = templateData;
               writeArticle.articleId = articleId;
               writeArticle.show();
               }
             else {
               Article.findOne({id: articleId}, function (articleData) {
                 var templateData = new can.Observe({
                   article: articleData.article,
                   //                    categories: article-categories.categories,
                   categoryId : categoryId,
                 });
                 writeArticle.articleData = templateData;
                 writeArticle.articleId = articleId;
                 writeArticle.show();
               }, function (xhr) {

               });
             }

        },
        /**
         * Refresh will be performed after image uploaded.
         */
        refresh: function() {
            var articleId = $('#id').val();
            var articleTitle = $('#articleTitle').val();
                if(articleId != undefined && articleId.length > 0 && articleTitle != undefined && articleTitle.length > 0) {
                    can.route.attr({'route': 'article/:title/:id', 'title': util.slashToBlank(util.truncate40(articleTitle)), 'id': articleId}, true);
                }
                return true;
            if(writeArticle.articleData.article != undefined) {
                if(articleId != undefined && articleId.length > 0) can.route.attr({'route': 'article/write/:id', 'id': articleId}, true);
                return false;
            }
            can.route.attr({'route': 'article/write'}, true);
            writeArticle.load();
            return false;
        },
        reload: function() {
            var articleId = $('#id').val();
            if(articleId != undefined && articleId.length > 0) {
                can.route.attr({'route': 'article/write/:id', 'id': articleId}, true);
                writeArticle.load(articleId);
            }
        },
        validate: function () {
             var  isValid = validation.minLength('title', 1,  'article.view.write.validation.title', false) &&
                    validation.minLengthTextarea('content', 1, 'article.view.write.validation.content');
//                    validation.minLength('tags', 1, 'article.view.write.validation.tags', false);
             return isValid;
        },
        '.write-article click': function (el, ev) {
           util.logDebug('write-article', writeArticle.articleId);
            ev.preventDefault();
            ev.stopPropagation();
            if(writeArticle.articleId == -1) {
                writeArticle.performCreate();
                return false;
            }
            writeArticle.performUpdate();

        },

        performCreate: function () {
            if (this.validate()) {
                var form = this.element.find('#articleWriteForm');
                var values = can.deparam(form.serialize());
                var $form = $(form);
                values.title = validation.replaceBadWords(values.title);
                values.content = validation.replaceBadWords(values.content);
                if (!$form.data('submitted')) {
                    $form.data('submitted', true);
                    var updateBtn = this.element.find('.change-status');
                    updateBtn.attr('disabled', 'disabled');
                    can.when(Article.create(values)).then(function (result) {
                        var id = writeArticle.element.find('#id');
                        var $document = $(document);
                        $document.ajaxError(function() {
                            can.route.attr({'route': ''}, true);
                        });
                        id.val(result.article.id);
                        util.showSuccessMsg(i18n.t('article.view.created.done'));
                        if(result.article.categoryId == 100) {
                            Refresh.load({'route': 'article/:title/:id', 'title': values.title, 'id': result.article.id});
                        }
                        writeArticle.refresh();
                    }, function (xhr) {
                        updateBtn.removeAttr('disabled');
                        $form.data('submitted', false);
                        util.handleError(xhr);
                    });
                }
            } else {
                util.showMessages();
            }
        },
        performUpdate: function () {
            if (this.validate()) {
                var form = this.element.find('#articleWriteForm');
                var values = can.deparam(form.serialize());
                var $form = $(form);
                values.title = validation.replaceBadWords(values.title);
                values.content = validation.replaceBadWords(values.content);
                 if (!$form.data('submitted')) {
                    $form.data('submitted', true);
                    var updateBtn = this.element.find('.change-status');
                    updateBtn.attr('disabled', 'disabled');
                    can.when(Article.update(writeArticle.articleId, values)).then(function (result) {
                        util.showSuccessMsg(i18n.t('article.view.updated.done'));
                        $form.data('submitted', false);
                        if(result.article.categoryId == 100) {
                            Refresh.load({'route': 'article/:title/:id', 'title': values.title, 'id': result.article.id});
                        }
                        writeArticle.refresh();
                    }, function (xhr) {
                        updateBtn.removeAttr('disabled');
                        $form.data('submitted', false);
                        util.handleError(xhr);
                    });
                }
            } else {
                util.showMessages();
            }
        },

        show: function () {
          this.element.html(can.view('views_article_write_stache', writeArticle.articleData));
          $('#articleContent').trumbowyg({
            mobile: true,
            tablet: true
          });
        }
    });

    /**
     * Router for article write.
     * @author dorajiarticle
     * @param {string} target
     * @function Router
     * @name article#ArticleWriteRouter
     * @constructor
     */
    var Router = can.Control.extend({
        defaults: {}
    }, {
        init: function () {
            writeArticle = undefined;
            util.logInfo('*article/WrtieArticle/Router', 'Initialized');
        },
        allocate: function () {
            var $app = util.getFreshApp();
            writeArticle = new WriteArticle($app);
        },
        load: function (articleId, categoryId) {
            util.allocate(this, writeArticle);
            writeArticle.load(articleId, categoryId);
        },
        'articles/write route': function () {
            this.load();
        },
        'articles/notice/write route': function () {
            this.load(-1,100);
        },
        'articles/general/write route': function () {
            this.load(-1,200);
        },
        'articles/etc/write route': function () {
            this.load(-1,300);
        },
        'articles/write/:id route': function (data) {
            this.load(data.id);
        }
    });
    return Router;
});
