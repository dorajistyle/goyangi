define(['jquery', 'can', 'app/models/user/user', 'app/models/user/user-activate',
        'app/models/user/user-admin', 'app/models/user/role', 'app/models/user/user-role', 'app/partial/destroy',
    'settings', 'util', 'validation', 'i18n', 'bloodhound', 'jquery.typeahead'],
    function ($, can, User, UserActivate, UserAdmin, Role, UserRole, Destroy, settings, util, validation, i18n, Bloodhound) {
    'use strict';

    /**
     * Instance of Admin Contorllers.
     * @private
     */
    var createUserRole, showUser, destroyUserRole, toggleActivateUser;
    var modalDestroyId =  "deleteUserConfirm";
    var modalDestroyName =  "user";
    var ShowUser = can.Control.extend({
        init: function () {
            util.logInfo('*admin/ShowUser', 'Initialized');

        },
        searchFilter: function(data) {
            var users = [];
                $.each(users, function(i, obj) {
                    users.push({email: obj.email});
                });
            return users;
        },
        searchInit: function() {
            var $searchUser = $('#searchUser');
            var usersByEmail = new Bloodhound({
              datumTokenizer: Bloodhound.tokenizers.obj.whitespace('email'),
              queryTokenizer: Bloodhound.tokenizers.whitespace,
//              prefetch: '../data/users.json',
              remote: {url: settings.apiPath+'/user/email/%QUERY/list',
                       filter: function(resp, status, jqXhr) {
                           var users = [];
                           $.each(resp.users, function(i, obj) {
                               users.push({id: obj.id, email: obj.email});
                           });
                           return users;
                       }
              }
            });
            usersByEmail.initialize();
            $searchUser.typeahead('destroy');
            $searchUser.typeahead({
                hint: true,
                highlight: true,
                minLength: 1
            },{
              name: 'users',
              displayKey: 'email',
              source: usersByEmail.ttAdapter()

            });
            $searchUser.bind('typeahead:selected', function(obj, datum) {
                createUserRole.userId = datum.id;
                createUserRole.showForm();
            });

        },
        load: function(page) {
            UserAdmin.findAll({page: page}, function (usersData) {
                Role.findAll({}, function (rolesData) {
                    var roles = {};
                    if(rolesData.roles != undefined) {
                        roles = rolesData.roles;
                    }
                    var templateData = new can.Observe({
                        users: usersData.users,
                        hasPrev: usersData.hasPrev,
                        hasNext: usersData.hasNext,
                        prevPage: usersData.currentPage-1,
                        nextPage: usersData.currentPage+1,
                        roles: roles,
                        currentUser: {},
                        modalDestroyId: modalDestroyId,
                        modalDestroyName: modalDestroyName
                    });
                    showUser.usersData = templateData;
                    $("#admin-user").html(
                        can.view('views_admin_admin-user_stache', showUser.usersData)
                    );
                    util.logDebug('*admin/ShowUser', 'loaded');
                    showUser.searchInit();
                });
            });
        },
        reload: function(page) {
          util.logDebug("page", page);
            if(page == undefined) {
                can.route.attr({'route': 'admin/:tab/:page', 'tab': 'user', 'page': 1});
            }
            UserAdmin.findAll({page: page}, function (usersData) {
                    showUser.usersData.attr('hasPrev',usersData.hasPrev);
                    showUser.usersData.attr('hasNext',usersData.hasNext);
                    showUser.usersData.attr('prevPage',usersData.currentPage-1);
                    showUser.usersData.attr('nextPage',usersData.currentPage+1);
                    util.refreshArray(showUser.usersData.users, usersData.users);
                    util.logDebug("reload", usersData);
                    showUser.searchInit();
            });
        }
    });
    /**
     * Control for CreateUser
     * @private
     * @author dorajistyle
     * @param {string} target
     * @name admin#CreateUser
     * @constructor
     */
    var CreateUserRole = can.Control.extend({
        init: function () {
            util.logInfo('*admin/CreateUser', 'Initialized');
        },



        '.select-user click': function (el, ev) {
            this.userId = util.getId(ev);
            this.showForm();
        },

        '.add-role-to-user click': function (el, ev) {
            ev.preventDefault();
            ev.stopPropagation();
            this.addRoleToUser();
            return false;
        },

       /**
         * Validate a form.
         * @memberof admin#Users
         */
        validate: function() {
            return validation.minLength('roleName', 1, 'admin.role.nameValidation')
        },
        hideForm: function() {
            $(this.element.find('#adminUserForm')).addClass('uk-hidden');
        },
        changeTypeahead: function(obj, datum) {
            $('#roleId').val(datum.id);
            createUserRole.roleId = datum.id;
            util.logDebug('role changed obj', obj);
            util.logDebug('role changed datum', datum);
        },
        showForm: function() {
                var form = this.element.find('#adminUserForm');
                var email = this.element.find('#adminUserEmail');
                var $form = $(form);
                var roles = [];
                var $roleName = $('#roleName');
                var $select = $(this.element.find('#adminUserRoles'));
                var options = '';
                $.each(showUser.usersData.roles, function(i, obj) {
                    roles.push({id: obj.id, name: obj.name, description: obj.description});
                });
                var roleList = new Bloodhound({
              datumTokenizer: Bloodhound.tokenizers.obj.whitespace('name'),
              queryTokenizer: Bloodhound.tokenizers.whitespace,
              local: roles
            });
            roleList.initialize();


                $roleName.typeahead('destroy');
                $roleName.typeahead({
                    hint: true,
                    highlight: true,
                    minLength: 1
                },{
                  name: 'roles',
                  displayKey: 'name',
                  templates: {
                    empty: [
                      '<div class="empty-message">',
                      i18n.t('typeahead.emptyMessage'),
                      '</div>'
                    ].join('\n'),
                    suggestion: util.compileTemplate('<p><strong>{{name}}</strong> - {{description}}</p>')
                  },
                  source: roleList.ttAdapter()
                }).bind('typeahead:selected', function(obj, datum) {
                    createUserRole.changeTypeahead(obj, datum);
                }).bind('typeahead:autocompleted', function(obj, datum) {
                    createUserRole.changeTypeahead(obj, datum);
                });

                can.when(UserAdmin.findOne({id: createUserRole.userId})).then(function (result) {
                    util.logJson('load',result);
                    util.logJson("current user", showUser.usersData.currentUser);
                    util.refreshObject(showUser.usersData.currentUser, result.user);
                    util.refreshArray(showUser.usersData.currentUser.Roles, result.user.Roles);
                    util.logJson("current user", showUser.usersData.currentUser);
                    $form.removeClass('uk-hidden');
                });
        },
        /**
         * perform Admin action.
         * @memberof admin#Users
         */
        addRoleToUser: function () {
            if (this.validate()) {
                var form = this.element.find('#adminUserForm');
                var values = can.deparam(form.serialize());
                util.logDebug('addRoleToUser',values);
                var $form = $(form);
                if (!$form.data('submitted')) {
                    $form.data('submitted', true);
                    var addRoleBtn = this.element.find('.add-role-to-user');
                    addRoleBtn.attr('disabled', 'disabled');
                    can.when(UserRole.create({userId: createUserRole.userId, roleId: createUserRole.roleId})).then(function (result) {
                        util.logDebug('role added', result);
                        util.logDebug('role added roleId', createUserRole.roleId);
                        addRoleBtn.removeAttr('disabled');
                        $form.data('submitted', false);
                        showUser.reload();
                        createUserRole.showForm();
                        util.showSuccessMsg(i18n.t('admin.user.role.add.done'));
                    }, function (xhr) {
                        addRoleBtn.removeAttr('disabled');
                        $form.data('submitted', false);
                        util.handleError(xhr);
                        // util.showErrorMsg(i18n.t('admin.user.role.add.fail'));
                    });
                }
            } else {
                 util.showMessages();
            }
        }


    });

//
    var DestroyUserRole = can.Control.extend({
        init: function () {
            this.isConfirmed = false;
            util.logInfo('*admin/DestroyRole', 'Initialized');
        },
        '.delete-role-from-user-confirm click': function (el, ev) {
            ev.preventDefault();
            ev.stopPropagation();
            this.roleId = util.getId(ev);
            this.userId = $(ev.target.parentNode.parentNode.parentNode).data('id');
            this.performConfirm();
            return false;
        },
        '.delete-role-from-user-final click': function (el, ev) {
            ev.preventDefault();
            ev.stopPropagation();
            this.isConfirmed = true;
            createUserRole.hideForm();
            can.when(destroyUserRole.modal.hide()).then(this.performDestroy());
            return false;
        },
        '.cancel-confirm click': function (el, ev) {
            ev.preventDefault();
            ev.stopPropagation();
            this.isConfirmed = false;
            return false;
        },

        /**
         * Show confirm modal.
         * @memberof admin#DestroyRole
         */
        performConfirm: function () {
            destroyUserRole.modal = $.UIkit.modal('#deleteRoleFromUserConfirm');
            destroyUserRole.modal.show();
        },
        /**
         * perform Destory action.
         * @memberof admin#DestroyRole
         */
        performDestroy: function () {
            var item = this.element.find('#deleteRoleConfirm');
            var $form = $(item);
            if (!$form.data('submitted')) {
                $form.data('submitted', true);
                var destroyBtn = this.element.find('.delete-role-from-user');
                destroyBtn.attr('disabled', 'disabled');
                util.logDebug('userId', destroyUserRole.userId);
                util.logDebug('roleId', destroyUserRole.roleId);
                can.when(UserRole.destroy(destroyUserRole.userId, destroyUserRole.roleId)).then(function () {
                    util.showSuccessMsg(i18n.t('admin.user.role.delete.done'));
                    destroyBtn.removeAttr('disabled');
                    $form.data('submitted', false);
                    showUser.reload();
                }, function (xhr) {
                    destroyBtn.removeAttr('disabled');
                    $form.data('submitted', false);
                    util.handleError(xhr);
                    // util.showErrorMsg(i18n.t('admin.user.role.delete.fail'));
                });

            }
        }
    });

    /**
     * Control for ToggleActivateUser
     * @private
     * @author dorajistyle
     * @param {string} target
     * @name admin#DestroyUser
     * @constructor
     */
    var ToggleActivateUser = can.Control.extend({
        init: function () {
            this.isConfirmed = false;
            util.logInfo('*admin/ToggleActivateUser', 'Initialized');
        },
        '.toggle-activation click': function (el, ev) {
            ev.preventDefault();
            ev.stopPropagation();
            this.userId = util.getId(ev);
            this.activation = ev.target.checked ? 1 : 0;
            this.performToggleActivate();
            return false;
        },

        /**
         * perform Destory action.
         * @memberof admin#DestroyUser
         */
        performToggleActivate: function () {
                util.logDebug('performToggleActivate',toggleActivateUser.activation);
                can.when(UserActivate.update(toggleActivateUser.userId,{activation: toggleActivateUser.activation})).then(function (result) {
                    util.logJson('performToggleActivate',result);
                    util.showSuccessMsg(i18n.t('admin.user.toggleActivate.done'));
                    showUser.reload();
                }, function (xhr) {
                  util.handleError(xhr);
                    // util.showErrorMsg(i18n.t('admin.user.toggleActivate.fail'));
                });

        }
    });

    /**
     * Router for user.
     * @author dorajistyle
     * @param {string} target
     * @function Router
     * @name admin#Router
     * @constructor
     */
    var Router = can.Control.extend({
        defaults: {}
    }, {
        init: function () {
            showUser = undefined;
            createUserRole = undefined;
            destroyUserRole = undefined;
            toggleActivateUser = undefined;
            util.logInfo('*admin/Router', 'Initialized')
        },
         allocate: function () {
            var $app = util.getFreshDiv('admin-user');
            showUser = new ShowUser($app);
            createUserRole = new CreateUserRole($app);
            destroyUserRole = new DestroyUserRole($app);
            new Destroy($app, {modalDestroyId: modalDestroyId, name: modalDestroyName, view: showUser, model: User});
            toggleActivateUser = new ToggleActivateUser($app);
        },
        load: function(page) {
            util.allocate(this, showUser);
            if(showUser.usersData == undefined || page == undefined){
                showUser.load(page);
                return false;
            }
            showUser.reload(page);
        }
    });

    return Router;
});
