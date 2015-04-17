define(['can', 'app/admin/roles', 'app/admin/users', 'app/models/user/user-current',
    'app/partial/tab', 'util', 'i18n', 'jquery'],
    function (can, Roles, Users, UserCurrent, Tab, util, i18n, $) {
    'use strict';


    /**
     * Instance of Admin Contorllers.
     * @private
     */
    var admin, tab, roles, users;

    /**
     * Control for Admin
     * @private
     * @author dorajistyle
     * @param {string} target
     * @name admin#Admin
     * @constructor
     */
    var Admin = can.Control.extend({
        init: function () {
            util.logInfo('*admin/Admin', 'Initialized');
        },
        load: function (tab, page) {
            UserCurrent.findOne({}, function (result) {
                if(!result.hasAdmin) {
                    can.route.attr({'route': ''}, true);
                }  else {
                    if(!admin.isReload) {
                      admin.show();
                      admin.isReload = true;
                      util.refreshTitle();
                    }
                    admin.selectTab(tab, page);
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
            this.element.html(can.view('views_admin_admin_stache', {
            }));
        },
        /**
         * Select tab
         * @param tabName
         */
        selectTab: function (tabName, page) {
            if (!tabName) tabName = 'basic';
            tab.activeTab(tabName);
            switch (tabName) {
                case 'role':
                    roles.load(page);
                    tab.showTab(tabName);
                    break;
                case 'user':
                    users.load(page);
                    tab.showTab(tabName);
                    break;
                default:
                    tab.showTab('role');
                    break;
            }
        }
    });

   /**
     * Router for user.
     * @author dorajistyle
     * @param {string} target
     * @function Router
     * @name users#Router
     * @constructor
     */
   var Router = can.Control.extend({
        defaults: {}
   }, {
       init: function () {
           util.logInfo('*admin/Router', 'Initialized')
       },
       allocate: function () {
           var $app = util.getFreshApp();
           tab = new Tab($app,{route: 'admin'});
           admin = new Admin($app);
           roles = new Roles();
           users = new Users();
       },
       'admin route': function () {
           can.route.attr({'route':'admin/:tab', 'tab':'role'}, true);
       },
       'admin/:tab/:page route': function (data) {

          if(admin === undefined || admin.isReload === undefined) {
             util.allocate(this, admin);
           }
           admin.load(data.tab, data.page);
       },
       'admin/:tab route': function (data) {
            util.allocate(this, admin);
           admin.isReload = false;
           admin.load(data.tab);
       }
   });

   return Router;
});
