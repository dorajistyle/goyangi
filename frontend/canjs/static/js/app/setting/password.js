define(['can', 'app/models/user/user', 'util', 'validation', 'i18n', 'jquery'],
    function (can, User, util, validation,  i18n, $) {
    'use strict';


    /**
     * Instance of User Contorllers.
     * @private
     */
    var password;


    /**
     * Control for Update Password
     * @private
     * @author dorajistyle
     * @param {string} target
     * @name setting#Password
     * @constructor
     */
    var Password = can.Control.extend({
        init: function () {
            util.logInfo('*User/Password', 'Initialized');
        },
        load: function () {
            password.show();
            util.refreshTitle();
        },
        /**
         * Show Setting view.
         * @memberof setting#Password
         */
        show: function () {
            this.element.html(can.view('views_setting_password_stache', {
                user: this.userData
            }));
        },

        '.update-user-password click': function (el, ev) {
            ev.preventDefault();
            ev.stopPropagation();
            this.performPassword();
            return false;
        },

        validatePassword: function () {
            return validation.minLength('currentPassword', 6, 'validation.password')
                && validation.minLength('newPassword', 6, 'validation.password')
                && validation.minLength('newPasswordConfirm', 6, 'validation.password')
                && validation.maxLength('currentPassword', 20, 'validation.password')
                && validation.maxLength('newPassword', 20, 'validation.password')
                && validation.maxLength('newPasswordConfirm', 20, 'validation.password')
                && validation.isIdentical('newPassword', 'newPasswordConfirm', 'validation.newPasswordConfirm');
        },
        /**
         * perform the Password password action.
         * @memberof setting#Password
         */
        performPassword: function () {
            if (this.validatePassword()) {
                var form = this.element.find('#passwordForm');
                var values = can.deparam(form.serialize());
                var $form = $(form);
                if (!$form.data('submitted')) {
                    $form.data('submitted', true);
                    var updateBtn = this.element.find('.update-user-password');
                    updateBtn.attr('disabled', 'disabled');
                    User.update({id: values.id, type: "password", data: values}, function (result) {
                    // User.findOne({id: values.id}, function (user) {
                    //     user.attr(values);
                    //     user.save(function (result) {
                            updateBtn.removeAttr('disabled');
                            $form.data('submitted', false);
                            // if (result.passwordIncorrect) {
                            //     util.showErrorMsg(i18n.t('setting.password.passwordIncorrect'));
                            //     return false;
                            // }
                            util.showSuccessMsg(i18n.t('setting.password.done'));
                        }, function (xhr) {
                            updateBtn.removeAttr('disabled');
                            $form.data('submitted', false);
                            util.handleStatusWithErrorMsg(xhr, i18n.t('setting.password.fail'));
                        });
                    // });
                }
            } else {
                util.showMessages();
            }
        }
    });

    /**
     * Router for setting password.
     * @author dorajistyle
     * @param {string} target
     * @function Router
     * @name setting#PasswordRouter
     * @constructor
     */
     var Router = can.Control.extend({
        defaults: {}
        }, {
            init: function () {
                password = undefined;
                util.logInfo('*setting/Password/Router', 'Initialized');
            },
            allocate: function () {
                var $app = util.getFreshDiv('setting-password');
                password = new Password($app);
            },
            load: function(user) {
                util.logDebug('*setting/Password/Router', 'loaded');
                util.allocate(this, password);
                password.userData = user;
                password.load();
            }
        });

    return Router;
});
