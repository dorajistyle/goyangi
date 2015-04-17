define(['can', 'app/models/user/user-password', 'app/models/user/user-email', 'refresh', 'util', 'validation', 'i18n', 'jquery'],
    function (can, UserPassword, UserEmail, Refresh, util, validation, i18n, $) {
    'use strict';

    var sendPasswordResetToken, passwordReset;

     /**
     * SendPasswordResetToken
     * @private
     * @author dorajistyle
     * @param {string} target
     * @name users#SendPasswordResetToken
     * @constructor
     */
    var SendPasswordResetToken = can.Control.extend({
        init: function () {
            util.logInfo('*User/SendPasswordResetToken', 'Initialized');
        },
        '.send-password-reset-token click': function (el, ev) {
            ev.preventDefault();
            ev.stopPropagation();
            this.sendToken();
            return false;
        },
        /**
        * Show password reset view.
        * @memberofusers#SendPasswordResetToken
        */
        show: function () {
          this.element.html(can.view('views_user_send-password-reset-token-form_stache', {}));
          util.refreshTitle();
        },
        /**
         * Validate a form.
         * @memberof users#Create
         */
        validate: function () {
            if(validation.validateEmail('email', 'validation.email', false) == false) {
              return false;
            }
              return true;
        },
         sendToken: function () {
            if (this.validate()) {
                util.logDebug('sendToken', 'validated');
                var form = this.element.find('#passwordResetTokenForm');
                var values = can.deparam(form.serialize());
                var $form = $(form);
                if (!$form.data('submitted')) {
                    $form.data('submitted', true);
                    var sendTokenBtn = this.element.find('.send-password-reset-token');
                    sendTokenBtn.attr('disabled', 'disabled');
                    var $registrationEmail = $('#email');
                    UserEmail.findOne({email: $registrationEmail.val()}, function (result) {
                      if (result.user) {
                        can.when(UserPassword.create(values)).then(function (result) {
                            util.logJson('register', result);
                            sendTokenBtn.removeAttr('disabled');
                            $form.data('submitted', false);
                            util.showSuccessMsg(i18n.t('passwordReset.send.sent.done'));
                            Refresh.load({'route': ''});
                        }, function (xhr) {
                            sendTokenBtn.removeAttr('disabled');
                            $form.data('submitted', false);
                            util.handleStatusWithErrorMsg(xhr, i18n.t('passwordReset.send.sent.fail'));
                        });
                      } else {
                        util.showWarningMsg(i18n.t('passwordReset.send.validation.emailNotExist'));
                        sendTokenBtn.removeAttr('disabled');
                        $form.data('submitted', false);
                      }
                    }, function(xhr) {
                      util.handleError(xhr);
                    });
                }
            } else {
                util.showMessages();
            }
        }
    });

    /**
    * PasswordReset
    * @private
    * @author dorajistyle
    * @param {string} target
    * @name users#PasswordReset
    * @constructor
    */
    var PasswordReset = can.Control.extend({
      init: function () {
        util.logInfo('*User/PasswordReset', 'Initialized');
      },
      '.reset-password click': function (el, ev) {
        ev.preventDefault();
        ev.stopPropagation();
        this.resetPassword();
        return false;
      },
      /**
      * Show password reset view.
      * @memberofusers#PasswordReset
      */
      show: function (token) {
        this.token = token;
        this.element.html(can.view('views_user_password-reset-form_stache', {token: this.token}));
        util.refreshTitle();
      },
      /**
      * Validate a form.
      * @memberof users#PasswordReset
      */
      validate: function () {
        return validation.minLength('newPassword', 6, 'validation.password')
        && validation.minLength('newPasswordConfirm', 6, 'validation.password')
        && validation.maxLength('newPassword', 20, 'validation.password')
        && validation.maxLength('newPasswordConfirm', 20, 'validation.password')
        && validation.isIdentical('newPassword', 'newPasswordConfirm', 'validation.newPasswordConfirm');
      },
      resetPassword: function () {
        if (this.validate()) {
          util.logDebug('resetToken', 'validated');
          var form = this.element.find('#passwordResetForm');
          var values = can.deparam(form.serialize());
          util.logDebug('resetToken values', values);
          var $form = $(form);
          if (!$form.data('submitted')) {
            $form.data('submitted', true);
            var sendTokenBtn = this.element.find('.reset-password');
            sendTokenBtn.attr('disabled', 'disabled');
            can.when(UserPassword.update(values)).then(function (result) {
              util.logJson('register', result);
              sendTokenBtn.removeAttr('disabled');
              $form.data('submitted', false);
              util.showSuccessMsg(i18n.t('passwordReset.reset.updated.done'));
              Refresh.load({'route': ''});
            }, function (xhr) {
              sendTokenBtn.removeAttr('disabled');
              $form.data('submitted', false);
              util.handleError(xhr);
              // util.handleStatusWithErrorMsg(xhr, i18n.t('passwordReset.reset.updated.fail'));
            });
          }
        } else {
          util.showMessages();
        }
      }
    });


    /**
     * Router for reset-password.
     * @author dorajistyle
     * @param {string} target
     * @function Router
     * @name reset-password#Router
     * @constructor
     */
    var Router = can.Control.extend({
        defaults: {}
    }, {
        init: function (target) {
            util.logInfo('*ResetPassword/Router', 'Initialized')
        },
        allocateSend: function() {
            var $app = util.getFreshApp();
            sendPasswordResetToken = new SendPasswordResetToken($app);
        },
        allocateReset: function() {
          var $app = util.getFreshApp();
          passwordReset = new PasswordReset($app);
        },
        'send/password/reset/form route': function () {
            this.allocateSend();
            sendPasswordResetToken.show();
        },
        'reset/password/:token route': function (data) {
          this.allocateReset();
          passwordReset.show(data.token);
        }
    });

    return Router;


});
