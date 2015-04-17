define(['can',
        'util', 'validation', 'i18n', 'jquery'],
    function (can, util, validation, i18n, $) {
    'use strict';


    /**
     * Instance of comment Contorllers.
     * @private
     */
    var showComment, createComment, destroyComment;
    var parentId, parentData, view, Comment;

    var CreateComment =  can.Control.extend({
        init: function () {
            util.logInfo('*partial/comment/CreateComment', 'Initialized');
        },
        '.submit-comment click': function (el, ev) {
            ev.preventDefault();
            ev.stopPropagation();
            this.performCreate();
            return false;
        },
        initForm: function () {
            var $content = $(this.element.find('#content'));
            $content.val('');
        },
        validate: function (form) {
             var isValid = validation.minLength('userId',1, 'validation.unauthorized', true) &&
                 validation.minLength('parentId',1, 'validation.unknown', true) &&
                 validation.minLengthTextareaOfForm(form, 'content',1, 'validation.content');
             return isValid;
        },
        performCreate: function () {
            var form = this.element.find('#commentForm');
            if(this.validate(form)) {
                var values = can.deparam(form.serialize());
                values.content = validation.replaceBadWords(values.content);
                var $form = $(form);
                if (!$form.data('submitted')) {
                    $form.data('submitted', true);
                    var followBtn = this.element.find('.perform-create');
                    followBtn.attr('disabled', 'disabled');
                    can.when(Comment.create(values)).then(function () {
                        followBtn.removeAttr('disabled');
                        $form.data('submitted', false);
                        util.showSuccessMsg(i18n.t('comment.created.done'));
                        showComment.reload(1);
                        createComment.initForm();
                    }, function (xhr) {
                      util.handleError(xhr);
                        // util.handleStatusWithErrorMsg(xhr, i18n.t('comment.created.fail'));
                    });
                }
            } else {
                util.showMessages();
            }
        }
    });
    var ShowComment = can.Control.extend({
        init: function () {
            util.logInfo('*partial/comment/ShowComment', 'Initialized');
        },
        reload: function(page) {
            util.logDebug('reload','performed');
            Comment.findAll({id: parentId, data: {currentPage: page}}, function (commentsData) {
                    util.logDebug('findall Comment','performed');
                    showComment.allocate(commentsData);
            });
        },
        '.show-more-comments click': function (el, ev) {
            util.logDebug('more-comments','performed');
            util.logDebug('next-page',parentData.commentList.currentPage+1);

            Comment.findAll({id: parentId, data: {currentPage: parentData.commentList.currentPage+1}}, function (commentsData) {
                    util.logDebug('findall Comment','performed');
                    showComment.more(commentsData);
            });
        },
        allocatePage: function (commentsData) {
            parentData.commentList.attr('hasPrev',commentsData.hasPrev);
            parentData.commentList.attr('hasNext',commentsData.hasNext);
            parentData.commentList.attr('currentPage',commentsData.currentPage);
            parentData.commentList.attr('count',commentsData.count);
        },
        // changeCount: function (count) {
        //     parentData.commentList.attr('count',parseInt(parentData.commentList.attr('count'))+count);
        // },
        more: function(commentsData) {
            showComment.allocatePage(commentsData);
            parentData.commentList.comments.attr(parentData.commentList.comments.attr().concat(commentsData.comments.attr()));
        },
        allocate: function(commentsData) {
            util.logDebug('allocate Comment','performed');
            showComment.allocatePage(commentsData);
            util.logDebug('comments', parentData.commentList.comments);
            util.logDebug('new comments', commentsData.comments);
            util.refreshArray(parentData.commentList.comments, commentsData.comments);
        }
    });

    var DestroyComment = can.Control.extend({
        init: function () {
            util.logInfo('*partial/comment/DestroyComment', 'Initialized');
        },
        '.destroy-comment click': function (el, ev) {
            ev.preventDefault();
            ev.stopPropagation();
            destroyComment.commentId = util.getId(ev);
            destroyComment.parentId = util.getData(ev,"parent-id");
            this.performConfirm();
            return false;
        },
         '.destroy-comment-final click': function (el, ev) {
            ev.preventDefault();
            ev.stopPropagation();
            this.isConfirmed = true;
            can.when(destroyComment.modal.hide()).then(this.performDestroy());
            return false;
        },
        '.cancel-confirm click': function (el, ev) {
            ev.preventDefault();
            ev.stopPropagation();
            this.isConfirmed = false;
            return false;
        },

        /**
         * Show confirm modal.
         * @memberof comment#DestroyComment
         */
        performConfirm: function () {
            destroyComment.modal = $.UIkit.modal('#destroyCommentConfirm');
            destroyComment.modal.show();
//            $('#destroyCommentConfirm').modal();
        },
        /**
         * perform Destory action.
         * @memberof comment#DestroyComment
         */
        performDestroy: function () {
            var confirm = this.element.find('#destroyCommentConfirm');
            var $form = $(confirm);
            if (!$form.data('submitted')) {
                $form.data('submitted', true);
                var destroyBtn = this.element.find('.destroy-comment-final');
                destroyBtn.attr('disabled', 'disabled');
                can.when(Comment.destroy(destroyComment.parentId, destroyComment.commentId)).then(function () {
                    util.showSuccessMsg(i18n.t('comment.deleted.done'));
                    destroyBtn.removeAttr('disabled');
                    $form.data('submitted', false);
                    showComment.reload(1);
                }, function (xhr) {
                    destroyBtn.removeAttr('disabled');
                    $form.data('submitted', false);
                    util.handleError(xhr);
                    // if(xhr.status == 409) {
                    //     util.showErrorMsg(i18n.t('comment.deleted.conflict'));
                    //     return false;
                    // }
                    // util.showErrorMsg(i18n.t('comment.deleted.fail'));
                });

            }
        }

    });

    /**
     * Router for comment.
     * @author dorajistyle
     * @param {string} target
     * @function Router
     * @name partial/comment/Router
     * @constructor
     */
    var Router = can.Control.extend({
        defaults: {}
    }, {
        init: function () {
            util.logInfo('*partial/comment/Router', 'Initialized');
            util.logInfo('*partial/comment/Router this', this.element);
            var $app = this.element;
            // var $app = util.getFreshDiv(parentName);
            createComment = new CreateComment($app);
            destroyComment = new DestroyComment($app);
            showComment = new ShowComment($app);
            parentId = this.options.parentId;
            parentData = this.options.parentData;
            view = this.options.view;
            Comment = this.options.model;
        },
        allocate: function (parentName) {
            var $app = util.getFreshDiv(parent);
            createComment = new CreateComment($app);
            destroyComment = new DestroyComment($app);
            showComment = new ShowComment($app);
        }
    });
    return Router;
});
