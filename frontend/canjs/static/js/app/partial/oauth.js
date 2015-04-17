define(['can', 'app/models/oauth/google', 'app/models/oauth/github',  'app/models/oauth/yahoo', 'app/models/oauth/facebook', 'app/models/oauth/twitter',
'app/models/oauth/linkedin', 'app/models/oauth/kakao', 'app/models/oauth/naver',
'util', 'i18n', 'jquery'],
function (can, Google, Github, Yahoo, Facebook, Twitter, Linkedin, Kakao, Naver,
     util, i18n, $) {
    'use strict';
    /**
    * Oauth
    * @private
    * @author dorajistyle
    * @param {string} target
    * @name partial#Oauth
    * @constructor
    */
    var Oauth = can.Control.extend({
      init: function () {
        util.logInfo('*partial/Oauth', 'Initialized');
        this.view = this.options.view;

      },
      '.oauth-google click': function (el, ev) {
        this.hideModal(ev);
        can.when(Google.findAll()).then(function (result) {
          window.location.replace(result.url);
        });
        return false;
      },
      '.oauth-github click': function (el, ev) {
        this.hideModal(ev);
        can.when(Github.findAll()).then(function (result) {
          window.location.replace(result.url);
        });
        return false;
      },
      // '.oauth-yahoo click': function (el, ev) {
      //   this.hideModal(ev);
      //   can.when(Yahoo.findAll()).then(function (result) {
      //     window.location.replace(result.url);
      //   });
      //   return false;
      // },
      '.oauth-facebook click': function (el, ev) {
        this.hideModal(ev);
        can.when(Facebook.findAll()).then(function (result) {
          window.location.replace(result.url);
        });
        return false;
      },
      // '.oauth-twitter click': function (el, ev) {
      //   this.hideModal(ev);
      //   can.when(Twitter.findAll()).then(function (result) {
      //     window.location.replace(result.url);
      //   });
      //   return false;
      // },
      '.oauth-linkedin click': function (el, ev) {
        this.hideModal(ev);
        can.when(Linkedin.findAll()).then(function (result) {
          window.location.replace(result.url);
        });
        return false;
      },
      // '.oauth-kakao click': function (el, ev) {
      //   this.hideModal(ev);
      //   can.when(Kakao.findAll()).then(function (result) {
      //     window.location.replace(result.url);
      //   });
      //   return false;
      // },
      // '.oauth-naver click': function (el, ev) {
      //   this.hideModal(ev);
      //   can.when(Naver.findAll()).then(function (result) {
      //     window.location.replace(result.url);
      //   });
      //   return false;
      // },
      '.close-oauth-modal click': function (el, ev) {
        ev.preventDefault();
        ev.stopPropagation();
        this.isConfirmed = false;
        return false;
      },
      '.oauth-connect click': function (el, ev) {
        ev.preventDefault();
        ev.stopPropagation();
        this.showModal();
        return false;
      },
      showModal: function () {
        if (this.modal == undefined) {
          this.modal = $.UIkit.modal('#ConnectOauth');
        }
        this.modal.show();
      },
      hideModal: function (ev) {
        ev.preventDefault();
        ev.stopPropagation();
        this.isConfirmed = true;
        if(this.modal != undefined) {
          this.modal.hide();
        }
        return false;
      },
      '.revoke-oauth-google click': function (el, ev) {
        this.revokeContent = i18n.t('user.view.oauth.provider.google');
        this.revokeModel = Google;
        this.showRevokeModal(ev);
        return false;
      },
      '.revoke-oauth-github click': function (el, ev) {
        this.revokeContent = i18n.t('user.view.oauth.provider.github');
        this.revokeModel = Github;
        this.showRevokeModal(ev);
        return false;
      },
      // '.revoke-oauth-yahoo click': function (el, ev) {
        // this.revokeContent = i18n.t('user.view.oauth.provider.yahoo');
        // this.revokeModel = Yahoo;
        // this.showRevokeModal(ev);
      //   return false;
      // },
      '.revoke-oauth-facebook click': function (el, ev) {
        this.revokeContent = i18n.t('user.view.oauth.provider.facebook');
        this.revokeModel = Facebook;
        this.showRevokeModal(ev);
        return false;
      },
      // '.revoke-oauth-twitter click': function (el, ev) {
      // this.revokeContent = i18n.t('user.view.oauth.provider.twitter');
      // this.revokeModel = Twitter;
      // this.showRevokeModal(ev);
      //   return false;
      // },
      '.revoke-oauth-linkedin click': function (el, ev) {
        this.revokeContent = i18n.t('user.view.oauth.provider.linkedin');
        this.revokeModel = Linkedin;
        this.showRevokeModal(ev);
        return false;
      },
      // '.revoke-oauth-kakao click': function (el, ev) {
      //   this.revokeContent = i18n.t('user.view.oauth.provider.kakao');
      //   this.revokeModel = Kakao;
      //   this.showRevokeModal(ev);
      //   return false;
      // },
      // '.revoke-oauth-naver click': function (el, ev) {
      //   this.revokeContent = i18n.t('user.view.oauth.provider.naver');
      //   this.revokeModel = Naver;
      //   this.showRevokeModal(ev);
      //   return false;
      // },
      showRevokeModal: function () {
        if (this.revokeModal == undefined) {
          this.revokeModal = $.UIkit.modal('#revokeOauth');
        }
        $("#revokeContent").html(i18n.t('user.view.oauth.revoke.content.prefix')+this.revokeContent+i18n.t('user.view.oauth.revoke.content.postfix'));
        this.revokeModal.show();
      },
      hideRevokeModal: function (ev) {
        ev.preventDefault();
        ev.stopPropagation();
        this.isConfirmed = true;
        if(this.revokeModal != undefined) {
          this.revokeModal.hide();
        }
        return false;
      },
      '.revoke-oauth-final click': function (el, ev) {
        this.hideRevokeModal(ev);
        // this.revokeModal.destroy();
        var oauth = this;
        can.when(this.revokeModel.destroy()).then(function (result) {
          if(oauth.view != undefined) {
            oauth.view.reload(result.oauthStatus);
          }
        },function (xhr) {
          util.handleStatus(xhr);
        });
        // can.when(Naver.findAll()).then(function (result) {
        //   window.location.replace(result.url);
        // });
        return false;
      },


    });
    return Oauth;
  });
