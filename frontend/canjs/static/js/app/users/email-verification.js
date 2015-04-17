define(['can', 'app/models/user/user-email-verification', 'refresh', 'util', 'validation', 'i18n', 'jquery'],
    function (can, UserEmailVerification, Refresh, util, validation, i18n, $) {
    'use strict';

    var sendEmailVerificationToken, emailVerification;

    /**
    * SendEmailVerificationToken
    * @private
    * @author dorajistyle
    * @param {string} target
    * @name users#SendEmailVerificationToken
    * @constructor
    */
    var SendEmailVerificationToken = can.Control.extend({
      init: function () {
        util.logInfo('*User/SendEmailVerificationToken', 'Initialized');
      },
      '.send-email-verification-token click': function (el, ev) {
        ev.preventDefault();
        ev.stopPropagation();
        this.sendToken();
        return false;
      },
      show: function () {
        this.element.html(can.view('views_user_send-email-verification-token-form_stache', {}));
        util.refreshTitle();
      },
      sendToken: function () {
          util.logDebug('sendToken', 'validated');
          var form = this.element.find('#emailVerificationTokenForm');
          var values = can.deparam(form.serialize());
          var $form = $(form);
          if (!$form.data('submitted')) {
            $form.data('submitted', true);
            var sendTokenBtn = this.element.find('.send-email-verification-token');
            sendTokenBtn.attr('disabled', 'disabled');
            can.when(UserEmailVerification.create(values)).then(function (result) {
              util.logJson('register', result);
              sendTokenBtn.removeAttr('disabled');
              $form.data('submitted', false);
              util.showSuccessMsg(i18n.t('emailVerification.send.sent.done'));
              Refresh.load({'route': ''});
            }, function (xhr) {
              sendTokenBtn.removeAttr('disabled');
              $form.data('submitted', false);
              util.handleError(xhr);
              // util.showErrorMsg(i18n.t('emailVerification.send.sent.fail'));
            });
          }
      }
    });


    /**
    * EmailVerification
    * @private
    * @author dorajistyle
    * @param {string} target
    * @name users#EmailVerification
    * @constructor
    */
    var EmailVerification = can.Control.extend({
      init: function () {
        util.logInfo('*User/EmailVerification', 'Initialized');
      },
      show: function () {
        this.element.html(can.view('views_user_email-verification_stache', {verified: this.verified}));
        util.refreshTitle();
      },
      verifyEmail: function (token) {
          util.logDebug("verifyEmail token : ", token);
          can.when(UserEmailVerification.update({"token": token})).then(function () {
            emailVerification.verified = true;
            emailVerification.show();
          }, function (xhr) {
            emailVerification.verified = false;
            emailVerification.show();
          });
      }
    });


    /**
     * Router for emailVerification.
     * @author dorajistyle
     * @param {string} target
     * @function Router
     * @name emailVerification#Router
     * @constructor
     */
    var Router = can.Control.extend({
        defaults: {}
    }, {
        init: function (target) {
            util.logInfo('*EmailVerification/Router', 'Initialized')
        },
        allocateSend: function() {
          var $app = util.getFreshApp();
          sendEmailVerificationToken = new SendEmailVerificationToken($app);
        },
        allocateVerification: function() {
          var $app = util.getFreshApp();
          emailVerification = new EmailVerification($app);
        },
        'send/email/verification/form route': function () {
          this.allocateSend();
          sendEmailVerificationToken.show();
        },
        'verify/email/:token route': function (data) {
          this.allocateVerification();
          emailVerification.verifyEmail(data.token);
        }
    });

    return Router;


});
