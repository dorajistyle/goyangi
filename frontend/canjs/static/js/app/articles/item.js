define(['can', 'app/models/article/item', 'app/partial/comment', 'app/models/article/comment', 'app/partial/destroy',
        'app/partial/liking', 'app/models/article/liking', 'app/partial/sharing',
        'facebook', 'settings', 'util', 'validation', 'i18n', 'jquery', 'jquery.fileupload', 'jquery.magnific-popup'],
    function (can, Article, Comment, CommentModel, Destroy, Liking, LikingModel, Sharing, FB, settings, util, validation, i18n, $) {
    'use strict';


    /**
     * Instance of ShowArticle Contorllers.
     * @private
     */
    var showArticle, updateArticle;
    var modalDestroyId =  "destroyArticleConfirm";
    var modalDestroyName =  "article";
    /**
     * Control for show Article
     * @private
     * @author dorajistyle
     * @param {string} target
     * @name article#ShowArticle
     * @constructor
     */

    var ShowArticle = can.Control.extend({
        init: function () {
            util.logInfo('*share/ShowArticle', 'Initialized');
        },
        load: function (articleId) {
            var articleItem = this;
            util.logInfo('*share/ShowArticle', 'loaded');
            Article.findOne({id: articleId}, function (articleData) {
                articleItem.articleId = articleId;
                articleItem.loadArticleData(articleData);
                articleItem.element.hammer({drag: false, hold: true, transform: false,tapAlways: false, swipeVelocityX: 0.3});
             });
        },
        loadArticleData: function(articleData) {
            var articleItem = this;
            util.logJson('ShowArticle articleData',articleData);
            var templateData = new can.Observe({
                article: articleData.article,
                isAuthor: articleData.isAuthor,
                currentUserId: articleData.currentUserId,
                staticPath: util.getStaticPath(),
                modalDestroyId: modalDestroyId,
                modalDestroyName: modalDestroyName,
                parentId: articleData.article.id,
                parentTitle: util.truncate(articleData.article.title),
                parentDescription: util.truncate(articleData.article.content),
                parentTagNames: articleData.article.tagNames,
                commentList: articleData.article.commentList,
                likingList: articleData.article.likingList
            });
            articleItem.articleData = templateData;

            can.when(articleItem.show()).then(function () {
                       util.refreshMeta();
                       articleItem.initControl();
                }
            );

        },
        'tab': function (el, ev) {
            util.logDebug('article', 'tab');
        },
        'hold': function (el, ev) {
            util.logDebug('article', 'hold');
        },

        'swipeleft': function (el, ev) {
            util.logDebug('article', 'swipeleft');
            this.element.find('.right').click();
        },
        'swiperight': function (el, ev) {
            util.logDebug('article', 'swiperight');
            this.element.find('.left').click();
        },


        initFacebook: function () {
            var facebookScript = document.createElement('script');
            facebookScript.type = 'text/javascript';
            facebookScript.text = '(function(d, s, id) {var js, fjs = d.getElementsByTagName(s)[0];if (d.getElementById(id)) return;js = d.createElement(s); js.id = id;js.src = "//connect.facebook.net/ko_KR/sdk.js#xfbml=1&appId='+settings.facebookAppId+'&version=v2.0";fjs.parentNode.insertBefore(js, fjs);}(document, "script", "facebook-jssdk"));';
            util.logDebug('tourl', facebookScript.text);
            $('body').append(facebookScript);
            FB.XFBML.parse();
        },
        initControl: function (){
            var $articleWrapper = $(this.element.find('#articleWrapper'));
            var $headerDiv = $(this.element.find('#articleHeader'));
            var $commentDiv = $(this.element.find('#commentWrapper'));
            updateArticle = new UpdateArticle($headerDiv);
            new Liking($headerDiv, {parentId: showArticle.articleId, parentName: i18n.t("article.modelName"), view: showArticle, model: LikingModel, likePrefix: "liking.like", unlikePrefix: "liking.unlike"});
            new Destroy($headerDiv, {modalDestroyId: modalDestroyId, name: modalDestroyName, view: showArticle, model: Article});
            new Comment($commentDiv, {parentId: showArticle.articleId, parentData: showArticle.articleData.article, view: showArticle, model: CommentModel});
            new Sharing($articleWrapper);
        },

        initGestures: function (){

        },
        pageLikeOrUnlikeCallback: function(url, htmlElement) {
          util.fetchFBLikingCount(url).always(function(res){
                $('#fbCount').text(res);
            });

          util.logDebug("pageLikeOrUnlikeCallback", 'yeah!');
          util.logDebug('callback url', url);
          util.logDebug('callback elem', htmlElement);
        },

        loadImage: function (){
             var articleItem = this;
        },

        reload: function () {
          can.route.attr({'route': 'articles'}, true);
        },
        show: function () {
            this.element.html(can.view('views_article_item_stache', this.articleData));
        }

    });

    var UpdateArticle = can.Control.extend({
      init: function () {
        util.logInfo('*article/UpdateArticle', 'Initialized');
      },
      '.update-article click': function (el, ev) {
        ev.preventDefault();
        ev.stopPropagation();
        updateArticle.id = util.getId(ev);
        this.performUpdate();
        return false;
      },
      performUpdate: function () {
        can.route.attr({'route': 'articles/write/:id', 'id': updateArticle.id}, true);
      }

    });

    /**
     * Router for profile.
     * @author dorajistyle
     * @param {string} target
     * @function Router
     * @name article#ShowArticle
     * @constructor
     */
    var Router = can.Control.extend({
        defaults: {}
    }, {
        init: function () {
            util.logInfo('*article/ShowArticle/Router', 'Initialized');
        },
        allocate: function () {
            var $app = util.getFreshApp();
            showArticle = new ShowArticle($app);
        },
        show: function(articleId) {
            util.allocate(this, showArticle);
            showArticle.load(articleId);
        },
        'article/:title/:id route': function (data) {
            this.show(data.id);
        }
    });
    return Router;
});
