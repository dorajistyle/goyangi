define(['can', 'app/models/user/authentication', 'app/models/user/user', 'app/models/user/user-email', 'app/models/user/user-username',
        'app/partial/oauth', 'refresh', 'util', 'validation', 'i18n', 'jquery'],
    function (can, Authentication, User, UserEmail, UserUsername,
      Oauth, Refresh, util, validation, i18n, $) {
    'use strict';
    /**
     * Oauth
     * @private
     * @author dorajistyle
     * @param {string} target
     * @name authentications#Create
     * @constructor
     */
    var authentication, registration;

    var UserAuthentication = can.Control.extend({
        init: function () {
            util.logInfo('*Authentications/Create', 'Initialized');

        },
        /**
         * Show login view.
         * @memberof authentications#Create
         */
        show: function () {
            this.element.html(can.view('views_user_login_stache', {registrationMessage : authentication.registrationMessage}));
            util.refreshTitle();
        },

        '.perform-login click': function (el, ev) {
            ev.preventDefault();
            ev.stopPropagation();
            this.performLogin();
            return false;
        },

        /**
         * Validate a form.
         * @memberof autentications#Create
         */
        validate: function () {
          if(validation.validateUsernameOrEmail('loginEmail') == false) {
            // var messages = util.getMessages();
            // util.logDebug("login validate messages", messages);
            // util.showMessages();
            // util.showErrorOnDiv(util.getMessages(), 'loginEmailMssg');
            return false;
          }


            // return validation.validateEmail('loginEmail')
            //     && validation.minLength('loginPassword', 8, 'validation.password')
            return validation.minLength('loginPassword', 6, 'validation.password')
                && validation.maxLength('loginPassword', 20, 'validation.password');
        },
        /**
         * Perform login action.
         * @memberof authentications#Create
         */
        performLogin: function () {
            if (this.validate()) {
                var form = this.element.find('#loginForm');
                var values = can.deparam(form.serialize());
                var $form = $(form);
                if (!$form.data('submitted')) {
                    $form.data('submitted', true);
                    var loginBtn = this.element.find('.perform-login');
                    loginBtn.attr('disabled', 'disabled');
                    util.logJson("performLogin", values);
                    can.when(Authentication.create(values)).then(function (result) {
                        util.logJson("performLogin Response", result);
//                        if (result.status) {
                            util.showSuccessMsg(i18n.t('login.welcome', result.username));
                            Refresh.load({'route': ''});
//                        } else {
//                            loginBtn.removeAttr('disabled');
//                            $form.data('submitted', false);
//                            util.showErrorMsg(i18n.t(result.msg));
//                        }

                    }, function (xhr) {
                         loginBtn.removeAttr('disabled');
                         $form.data('submitted', false);
                         util.handleError(xhr);
                         // legacy
                        //  var status = xhr.responseJSON.status;
                        //  var username = xhr.responseJSON.username;
                        //  if(username == undefined || username.length == 0){
                        //      util.showErrorMsg(i18n.t('login.userNotFound'));
                        //      return false;
                        //  }
                        // //  if(status != undefined) {
                        // //      util.showErrorMsg(i18n.t('login.passwordIncorrect'));
                        // //      return false;
                        // //  }
                        //  util.showErrorMsg(i18n.t('login.passwordIncorrect'));
                         return false;
                    });
                }
            } else {
                util.showMessages();
            }
        },
        /**
         * Perform logout action
         * @memberof authentications#Create
         */
        performLogout: function () {
            can.when(Authentication.destroy()).then(function () {
                util.deleteCookie('session');
                // util.deleteCookie('remember_token');
                Refresh.load({'route': ''});
//                can.when(Authentication.findAll()).then(function () {
//                    util.showErrorMsg(i18n.t('logout.error'));
//                }, function (xhr) {
//                    util.showSuccessMsg(i18n.t('logout.done'));
//                    Refresh.load('login');
//                });
            }, function (xhr) {
                util.handleStatusWithErrorMsg(xhr, i18n.t('logout.alreadyDone'));
            });
        }
    });

     /**
     *
     * @private
     * @author dorajistyle
     * @param {string} target
     * @name users#Create
     * @constructor
     */
    var Registration = can.Control.extend({
        init: function () {
            util.logInfo('*User/Create', 'Initialized');
        },
        '.register-user click': function (el, ev) {
            ev.preventDefault();
            ev.stopPropagation();
            this.performRegister();
            return false;
        },
        '#registrationUsername focusout': function (el, ev) {
          ev.preventDefault();
          ev.stopPropagation();
          this.checkUsernameAlreadyTaken(false);
          return false;
        },

        '#registrationEmail focusout': function (el, ev) {
            ev.preventDefault();
            ev.stopPropagation();
            //TODO Check Email already taken or not.
            this.checkEmailAlreadyTaken();
            return false;
        },
        /**
         * Validate a form.
         * @memberof users#Create
         */
        validate: function () {
            // return validation.validateEmail('registrationEmail')
            //     && validation.minLength('registrationPassword', 8, 'validation.password')
            //     && validation.maxLength('registrationPassword', 20, 'validation.password');
            if(validation.validateUsername('registrationUsername') == false) {
              // util.showErrorOnDiv(util.getMessages(), 'usernameMssg');
              return false;
            }
            if(validation.validateEmail('registrationEmail', 'validation.email', false) == false) {
              // util.showErrorOnDiv(util.getMessages(), 'emailMssg');
              return false;
            }

            if(!(validation.minLengthWithSpace('registrationPassword', 6, 'validation.password', false, true) &&
              validation.maxLength('registrationPassword', 20, 'validation.password'))) {
                // util.showErrorOnDiv(util.getMessages(), 'passwordMssg');
                return false;
              }
              if(!validation.validateIdentical('registrationPassword','registrationVerifyPassword','validation.verifyPassword')) {
                // util.showErrorOnDiv(i18n.t('validation.verifyPassword'), 'passwordMssg');
                return false;
              }

              return true;



        },
        /**
         * perform Register action.
         * @memberof users#Create
         */
        performRegister: function () {
            if (this.validate()) {
                util.logDebug('register-user', 'validated');
                var form = this.element.find('#registrationForm');
                var values = can.deparam(form.serialize());
                var $form = $(form);
                if (!$form.data('submitted')) {
                    $form.data('submitted', true);
                    var registerBtn = this.element.find('.register-user');
                    registerBtn.attr('disabled', 'disabled');
                    if(registration.providerId != undefined) values.providerId = registration.providerId;
//                    TODO User registration
                    can.when(User.create(values)).then(function (result) {
                        util.logJson('register', result);
                        registerBtn.removeAttr('disabled');
                        $form.data('submitted', false);
//                        if (result.status) {
                            util.showSuccessMsg(i18n.t('registration.welcome'));
                            Refresh.load({'route': ''});
//                        } else {
//                            util.showErrorMsg(i18n.t('registration.fail'));
//                        }

                    }, function (xhr) {
                        registerBtn.removeAttr('disabled');
                        $form.data('submitted', false);
                        util.handleError(xhr);
                    });
                }
            } else {
                util.showMessages();
            }
        },
        /**
         * Check a email that is already taken.
         * @memberof users#Create
         */
        checkEmailAlreadyTaken: function () {
            if (validation.validateEmail('registrationEmail', true)) {
                var $registrationEmail = $('#registrationEmail');
                UserEmail.findOne({email: $registrationEmail.val()}, function (result) {
                        if (result.user) {
                            util.showWarningMsg(i18n.t('registration.already', result.email));
                            $registrationEmail.val('');
                            return false;
                        } else {
                            util.showSuccessMsg(i18n.t('registration.available', result.email));
                        }
                        return true;
                    }
                );
            } else {
                util.clearMessages();
                return true;
            }
        },
        /**
        * Check a username that is already taken.
        * @memberof users#Create
        */
        checkUsernameAlreadyTaken: function (skipBlank) {
          if(validation.validateUsername('registrationUsername')) {
            var $registrationUsername = $('#registrationUsername');
            var $username = $registrationUsername.val();
            $username = util.trim($username);
            if($username != "") {
              UserUsername.findOne({username: $username}, function (result) {
                util.logDebug("checkUsernameAlreadyTaken result", result);
                if (result.user) {
                  util.showErrorMsg(i18n.t('registration.already', result.username), 'usernameMssg');
                  if (!skipBlank) {
                    $registrationUsername.val('');
                  }
                  //$registrationUsername.focus();
                  return false;
                } else {
                  util.showSuccessMsg(i18n.t('registration.available', result.username), 'usernameMssg');
                }
                return true;
              });
            }
          } else {
            util.showMessages();
            // util.showErrorOnDiv(util.getMessages(), 'usernameMssg');
            return true;
          }
        }

    });

    /**
     * Router for authentication.
     * @author dorajistyle
     * @param {string} target
     * @function Router
     * @name authentications#Router
     * @constructor
     */
    var Router = can.Control.extend({
        defaults: {}
    }, {
        init: function (target) {
            util.logInfo('*Authentications/Router', 'Initialized')
        },
        allocate: function() {
            var $app = util.getFreshApp();
            authentication = new UserAuthentication($app);
            registration = new Registration($app);
            new Oauth($app);
        },
        'login route': function () {
            this.allocate();
            authentication.show();
        },
        'logout route': function () {
            this.allocate();
            authentication.performLogout();
        }
    });

    return Router;


});
