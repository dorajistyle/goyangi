define(['can', 'app/models/user/user',
    'util', 'i18n', 'jquery'],
    function (can, User, util, i18n, $) {
    'use strict';

    /**
     * Instance of User Contorllers.
     * @private
     */
    var basic;

    /**
     * Control for User Profile
     * @private
     * @author dorajistyle
     * @param {string} target
     * @name users#Basic
     * @constructor
     */
    var Basic = can.Control.extend({
        init: function () {
            util.logInfo('*User/Basic', 'Initialized');
        },
        load: function (page) {
            basic.show();
            basic.initForm();
            util.refreshTitle();
        },
        initForm: function () {
            var gender = basic.userData.gender;
            $('#gender').val(gender);
            switch(gender) {
                case 0:
                    $('.female').addClass('active');
                    break;
                case 1:
                    $('.male').addClass('active');
                    break;
                case 2:
                    $('.unknown').addClass('active');
                    break;
            }
        },
        /**
         * Show Setting view.
         * @memberof users#Basic
         */
        show: function () {
            this.element.html(can.view('views_setting_basic_stache', {
                user: basic.userData
            }));
        },
        '.update-user-basic click': function (el, ev) {
            ev.preventDefault();
            ev.stopPropagation();
            this.performUpdateBasic();
            return false;
        },
        '.select-gender click': function (el, ev) {
            $('.select-gender').removeClass('active');
            $(ev.target).addClass('active');
            $('#gender').val(util.getData(ev, 'value'));
            return false;
        },
        /**
         * Validate a form.
         * @memberof users#Basic
         */
        validateBasic: function () {
            return validation.minLength('firstName', 1, 'validation.firstName')
               validation.min.minLength('lastName', 1, 'validation.lastName');
        },
        /**
         * perform Basic action.
         * @memberof users#Basic
         */
        performUpdateBasic: function () {
            if (this.validateBasic()) {
                var form = this.element.find('#updateBasicForm');
                var values = can.deparam(form.serialize());
                var $form = $(form);
                if (!$form.data('submitted')) {
                    $form.data('submitted', true);
                    var updateBtn = this.element.find('.update-user-basic');
                    updateBtn.attr('disabled', 'disabled');
                    User.findOne({id: values.id}, function (user) {
                        user.attr(values);
                        user.save(function (result) {
                            util.logJson('performUpdateBasic',result);
                            updateBtn.removeAttr('disabled');
                            $form.data('submitted', false);
                            util.showSuccessMsg(i18n.t('setting.updateBasic.done'));
                        }, function (xhr) {
                            updateBtn.removeAttr('disabled');
                            $form.data('submitted', false);
                            util.handleStatusWithErrorMsg(xhr, i18n.t('setting.updateBasic.fail'));
                        });
                    });
                }
            } else {
                util.showMessages();
            }
        }
    });


    /**
     * Router for basic.
     * @author dorajistyle
     * @param {string} target
     * @function Router
     * @name setting#BasicRouter
     * @constructor
     */
     var Router = can.Control.extend({
        defaults: {}
        }, {
            init: function () {
                basic = undefined;
                util.logInfo('*setting/Basic/Router', 'Initialized');
            },
            allocate: function () {
                var $app = util.getFreshDiv('setting-basic');
                basic = new Basic($app);
            },
            load: function(user) {
                util.logDebug('*setting/Basic/Router', 'loaded');
                util.allocate(this, basic);
                basic.userData = user;
                basic.load();
            }
        });

    return Router;
});
