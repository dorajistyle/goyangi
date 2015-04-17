define(['can', 'app/models/user/user', 'refresh', 'util', 'i18n', 'jquery'],
    function (can, User, Refresh, util, i18n, $) {
    'use strict';


    /**
     * Instance of User Contorllers.
     * @private
     */
    var destroy;


       /**
     *
     * @private
     * @author dorajistyle
     * @param {string} target
     * @name users#Destroy
     * @constructor
     */
    var Destroy = can.Control.extend({
        init: function () {
            this.isConfirmed = false;
            util.logInfo('*User/Destroy', 'Initialized');
        },

        load: function () {
            destroy.show();
            util.refreshTitle();
        },
        /**
         * Show Setting view.
         * @memberof admin#Admin
         */
        show: function () {
            this.element.html(can.view('views_setting_leave-our-service_stache', {
                user: this.userData
            }));
        },


        '.leave-our-service-confirm click': function (el, ev) {
            util.logInfo('*User/Destroy', 'Confirm Clicked in users');
            ev.preventDefault();
            ev.stopPropagation();
            this.performConfirm();
        },
        '.leave-our-service-final click': function (el, ev) {
            ev.preventDefault();
            ev.stopPropagation();
            this.isConfirmed = true;
            can.when(destroy.modal.hide()).then(this.performDestroy());
        },
        '.cancel-confirm click': function (el, ev) {
            ev.preventDefault();
            ev.stopPropagation();
            this.isConfirmed = false;
            this.performCancel();
        },

        /**
         * Show confirm modal.
         * @memberof users#Destroy
         */
        performCancel: function () {
//            $('#leaveOurServiceConfirm').addClass('uk-hidden');
        },

        /**
         * Show confirm modal.
         * @memberof users#Destroy
         */
        performConfirm: function () {
//            $('#leaveOurServiceConfirm').removeClass('uk-hidden');
            destroy.modal = $.UIkit.modal('#leaveOurServiceConfirm');
            destroy.modal.show();
        },
        /**
         * perform Destory action.
         * @memberof users#Destroy
         */
        performDestroy: function () {
            var form = this.element.find('#leaveOurServiceForm');
            var values = can.deparam(form.serialize());
            var $form = $(form);
            if (!$form.data('submitted')) {
                $form.data('submitted', true);
                var destroyBtn = this.element.find('.leave-our-service-final');
                destroyBtn.attr('disabled', 'disabled');
                can.when(User.destroy(values.id)).then(function () {
                    util.showSuccessMsg(i18n.t('setting.leaveOurService.done'));
                    Refresh.load({'route': ''});
                }, function (xhr) {
                    destroyBtn.removeAttr('disabled');
                    $form.data('submitted', false);
                    util.handleError(xhr);
                    // util.showErrorMsg(i18n.t('setting.leaveOurService.fail'));
                });

            }
        }
    });


 /**
     * Router for leave our service.
     * @author dorajistyle
     * @param {string} target
     * @function Router
     * @name setting#DestoryRouter
     * @constructor
     */
     var Router = can.Control.extend({
        defaults: {}
        }, {
            init: function () {
                destroy = undefined;
                util.logInfo('*setting/Destory/Router', 'Initialized');
            },
            allocate: function () {
                var $app = util.getFreshDiv('setting-leave-our-service');
                destroy = new Destroy($app);
            },
            load: function(user) {
                util.logDebug('*setting/Destory/Router', 'loaded');
                util.allocate(this, destroy);
                destroy.userData = user;
                destroy.load();
            }
        });

return Router;
});
