define(['can', 'util', 'validation', 'i18n', 'jquery'],
    function (can, util, validation, i18n, $) {
    'use strict';
    var parentId, parentName, view, Liking, Liked, likePrefix, unlikePrefix;

    var LikingControl =  can.Control.extend({
      init: function () {
        util.logInfo('*share/ShowArticle/LikingArticle', 'Initialized');
        parentId = this.options.parentId;
        parentName = this.options.parentName;
        view = this.options.view;
        Liking = this.options.model;
        Liked = this.options.modelLiked;
        likePrefix = this.options.likePrefix;
        unlikePrefix = this.options.unlikePrefix;
      },
      load: function () {
        var liking = this;
        if(liking.usersDisplayed === undefined) {
          liking.usersDisplayed = false;
        }
        liking.handleLikingdUsersView();
      },
      loadLikedUsers: function () {
        var liking = this;
        if(liking.likedUsersDisplayed === undefined) {
          liking.likedUsersDisplayed = false;
        }
        liking.handleLikedUsersView();
      },
      handleLikingdUsersView: function () {
        var liking = this;
        //          var $likedUsersWrapper = $(this.element.find('#likedUsersWrapper'));
        liking.modal = $.UIkit.modal('#likedUsersWrapper');
        if(liking.modal.isActive()) {
          // liking.usersDisplayed = false;
          liking.modal.hide();
          //                $likedUsersWrapper.modal('hide');
          //                $likedUsersWrapper.addClass('uk-hidden');

        } else {

          //                $likedUsersWrapper.modal({backdrop: 'false'});
          //                $likedUsersWrapper.removeClass('uk-hidden');
          if(!liking.usersLoaded) {
            Liking.findAll({id: parentId, data:{currentPage: 1}}, function(likingsData) {
              var templateData = new can.Observe({
                likings: likingsData.likings,
                currentPage: likingsData.currentPage,
                hasNext: likingsData.hasNext
              });
              liking.likingsData = templateData;
              var moreLikedUsers = liking.element.find('.show-more-liked-users');
              if(liking.likingsData.hasNext) { moreLikedUsers.removeClass('uk-hidden'); }
              liking.show();
              liking.usersDisplayed = true;
              liking.usersLoaded = true;
              liking.modal.show();
            });
          } else {
            liking.usersDisplayed = true;
            liking.modal.show();
          }

        }
      },
      handleLikedUsersView: function () {
        var liking = this;
        //          var $likedUsersWrapper = $(this.element.find('#likedUsersWrapper'));
        liking.modalLiked = $.UIkit.modal('#likingUsersWrapper');
        if(liking.modalLiked.isActive()) {
          // liking.usersDisplayed = false;
          liking.modalLiked.hide();
          //                $likedUsersWrapper.modal('hide');
          //                $likedUsersWrapper.addClass('uk-hidden');

        } else {

          //                $likedUsersWrapper.modal({backdrop: 'false'});
          //                $likedUsersWrapper.removeClass('uk-hidden');
          if(!liking.likedUsersLoaded) {
            Liked.findAll({id: parentId, data:{currentPage: 1}}, function(likedData) {
              var templateData = new can.Observe({
                liked: likedData.liked,
                currentPage: likedData.currentPage,
                hasNext: likedData.hasNext
              });
              liking.likedData = templateData;
              var moreLikingUsers = liking.element.find('.show-more-liking-users');
              util.logDebug("liking.likedData.hasNext", liking.likedData.hasNext);
              if(liking.likedData.hasNext) { moreLikingUsers.removeClass('uk-hidden'); }
              liking.show();
              liking.likedUsersDisplayed = true;
              liking.likedUsersLoaded = true;
              liking.modalLiked.show();
            });
          } else {
            liking.likedUsersDisplayed = true;
            liking.modalLiked.show();
          }

        }
      },
      '.perform-like click': function (el, ev) {
        ev.preventDefault();
        ev.stopPropagation();
        this.performLiking();
        return false;
      },
      '.perform-unlike click': function (el, ev) {
        ev.preventDefault();
        ev.stopPropagation();
        this.performUnlike();
        return false;
      },
      '.show-liked-users click': function (el, ev) {
        var liking = this;
        ev.preventDefault();
        ev.stopPropagation();
        if(util.getData(ev,'count') > 0) {
          liking.load();
        }
        return false;
      },
      '.show-liking-users click': function (el, ev) {
        var liking = this;
        ev.preventDefault();
        ev.stopPropagation();
        if(util.getData(ev,'count') > 0) {
          liking.loadLikedUsers();
        }
        return false;
      },

      '.show-more-liked-users click': function (el, ev) {
        ev.preventDefault();
        ev.stopPropagation();
        this.showMoreLikings();
        return false;
      },
      '.show-more-liking-users click': function (el, ev) {
        ev.preventDefault();
        ev.stopPropagation();
        this.showMoreLiked();
        return false;
      },

      initStatus: function () {
        var liking = this;
        liking.usersDisplayed = false;
        liking.usersLoaded = false;
        liking.likedUsersDisplayed = false;
        liking.likedUsersLoaded = false;
      },
      validate: function () {
        var isValid = validation.minLength('userId',1, 'validation.unauthorized', true) &&
        validation.minLength('parentId',1, 'validation.unknown', true);
        return isValid;
      },
      performLiking: function () {
        var liking = this;
        if(liking.validate()) {
          util.logDebug('performLiking', 'validation success');
          var form = liking.element.find('#likingsForm');
          var values = can.deparam(form.serialize());
          var $form = $(form);
          if (!$form.data('submitted')) {
            $form.data('submitted', true);
            var likeBtn = liking.element.find('.perform-like');
            likeBtn.attr('disabled', 'disabled');
            can.when(Liking.create(values)).then(function () {
              likeBtn.removeAttr('disabled');
              $form.data('submitted', false);
              util.showSuccessMsg(i18n.t(likePrefix+'.done', parentName));
              view.load(parentId);
              liking.initStatus();
            }, function (xhr) {
              util.handleError(xhr);
              // util.handleStatusWithErrorMsg(xhr, i18n.t(likePrefix+'.fail', parentName));
              likeBtn.removeAttr('disabled');
              $form.data('submitted', false);
            });
          }
        } else {
          util.showMessages();
        }
      },
      performUnlike: function () {
        var liking = this;
        if(this.validate()) {
          var form = liking.element.find('#likingsForm');
          var values = can.deparam(form.serialize());
          var $form = $(form);
          if (!$form.data('submitted')) {
            $form.data('submitted', true);
            var likeBtn = liking.element.find('.perform-unlike');
            likeBtn.attr('disabled', 'disabled');
            can.when(Liking.destroy(values.parentId,values.userId)).then(function () {
              likeBtn.removeAttr('disabled');
              $form.data('submitted', false);
              util.showSuccessMsg(i18n.t(unlikePrefix+'.done', parentName));
              view.load(parentId);
              liking.initStatus();
            }, function (xhr) {
              likeBtn.removeAttr('disabled');
              $form.data('submitted', false);
              util.handleError(xhr);
              // util.handleStatusWithErrorMsg(xhr, i18n.t(unlikePrefix+'.fail', parentName));
            });
          }
        } else {
          util.showMessages();
        }
      },
      show: function () {
        var liking = this;
        liking.element.find('#likedUsers').html(can.view('views_partial_liked-user-list_stache', liking.likingsData));
        liking.element.find('#likingUsers').html(can.view('views_partial_liking-user-list_stache', liking.likedData));
      },
      showMoreLikings: function () {
        var liking = this;
        var moreLikedUsers = liking.element.find('.show-more-liked-users');
        moreLikedUsers.attr('disabled', 'disabled');
        Liking.findAll({id: parentId, data:{currentPage: liking.likingsData.currentPage + 1}}, function (likingsData) {
          likingsData.hasNext ? moreLikedUsers.removeAttr('disabled') : moreLikedUsers.addClass('uk-hidden');
          liking.likingsData.attr('currentPage', likingsData.currentPage);
          liking.likingsData.attr('hasNext', likingsData.hasNext);
          liking.likingsData.likings.attr(liking.likingsData.likings.attr().concat(likingsData.likings.attr()));
          }
        );
      },
      showMoreLiked: function () {
        var liking = this;
        var moreLikingUsers = liking.element.find('.show-more-liking-users');
        moreLikingUsers.attr('disabled', 'disabled');
        Liked.findAll({id: parentId, data:{currentPage: liking.likedData.currentPage + 1}}, function (likedData) {
          likedData.hasNext ? moreLikingUsers.removeAttr('disabled') : moreLikingUsers.addClass('uk-hidden');
          liking.likedData.attr('currentPage', likedData.currentPage);
          liking.likedData.attr('hasNext', likedData.hasNext);
          liking.likedData.liked.attr(liking.likedData.liked.attr().concat(likedData.liked.attr()));
        }
      );
    }

  });
  return LikingControl;

});
