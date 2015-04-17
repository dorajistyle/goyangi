define(['can', 'app/setting/basic', 'app/setting/password', 'app/setting/connection', 'app/setting/leave-our-service'
    , 'app/models/user/user-current', 'app/partial/tab', 'util', 'i18n', 'jquery'],
    function (can, Basic, Password, Connection, LeaveOurService, UserCurrent, Tab, util, i18n, $) {
    'use strict';


    /**
     * Instance of User Contorllers.
     * @private
     */
    var tab, setting, basic, password, connection, leaveOurService;


 /**
     * Control for Setting
     * @private
     * @author dorajistyle
     * @param {string} target
     * @name setting#Setting
     * @constructor
     */
    var Setting = can.Control.extend({
        init: function () {
            util.logInfo('*setting/Setting', 'Initialized');
        },
        load: function (tab, page) {
            UserCurrent.findOne({}, function (result) {
                if(!result.user) {
                    can.route.attr({'route': ''}, true);
                } else {
                    if(!setting.isReload) {
                      setting.show();
                      util.logDebug('Setting load', 'it is performed');
                      setting.isReload = true;
                      util.refreshTitle();
                      util.setWrapperInfo(setting, 'settingWrapper');
                    }
                    setting.user = result.user;
                    util.replaceWrapper(setting, 'settingWrapper');
                    setting.selectTab(tab, page);
                }
            },function (xhr) {
                can.route.attr({'route': ''}, true);
            });
        },
        /**
         * Show Setting view.
         * @memberof admin#Admin
         */
        show: function () {
            util.logDebug('setting element', this);
            this.element.html(can.view('views_setting_setting_stache', {
            }));
        },
        /**
         * Select tab
         * @param tabName
         */
        selectTab: function (tabName) {
            if (!tabName) tabName = 'basic';
            util.logDebug('selectTab', 'performed');
            tab.activeTab(tabName);
            switch (tabName) {
                case 'basic':
                    basic.load(setting.user);
                    tab.showTab(tabName);
                    break;
                case 'password':
                    password.load(setting.user);
                    tab.showTab(tabName);
                    break;
                case 'connection':
                    connection.load();
                    tab.showTab(tabName);
                    break;
                case 'leave-our-service':
                    leaveOurService.load(setting.user);
                    tab.showTab(tabName);
                    break;
                default:
                    tab.showTab('basic');
                    break;
            }
        }


    });

    /**
     * Router for update.
     * @author dorajistyle
     * @param {string} target
     * @function Router
     * @name users#UpdateRouter
     * @constructor
     */
    var Router = can.Control.extend({
        defaults: {}
    }, {
        init: function () {
            util.logInfo('*Users/Setting/Router', 'Initialized');
        },
         allocate: function () {
            var $app = util.getFreshApp();
            tab = new Tab($app,{route: 'setting'});
            setting = new Setting($app);
            basic = new Basic();
            password = new Password();
            connection = new Connection();
            leaveOurService = new LeaveOurService();
        },
        'setting route': function () {
            can.route.attr({'route': 'setting/:tab', 'tab': 'basic'}, true);
        },
        'setting/:tab route': function (data) {
            util.allocate(this, setting);
            setting.load(data.tab);
        }
    });

    return Router;
});
