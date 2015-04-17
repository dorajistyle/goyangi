define(['can', 'app/models/user/role', 'app/partial/destroy', 'util', 'validation', 'i18n', 'jquery'],
    function (can, Role, Destroy, util, validation, i18n, $) {
    'use strict';


    /**
     * Instance of Admin Contorllers.
     * @private
     */
    var createRole, updateRole, showRole;
    var modalDestroyId =  "deleteRoleConfirm";
    var modalDestroyName =  "role";
    var ShowRole = can.Control.extend({
        init: function () {
            util.logInfo('*admin.roles/ShowRole', 'Initialized');
        },
        load: function(page) {
            Role.findAll({'page': page}, function (rolesData) {
                util.logDebug('roles data', rolesData.roles);
                var roles = [];
                if(rolesData.roles != null) {
                    roles = rolesData.roles;
                }
                var templateData = new can.Observe({
                    roles: roles,
                    hasPrev: rolesData.hasPrev,
                    hasNext: rolesData.hasNext,
                    prevPage: rolesData.currentPage-1,
                    nextPage: rolesData.currentPage+1,
                    modalDestroyId: modalDestroyId,
                    modalDestroyName: modalDestroyName
                });
                showRole.rolesData = templateData;
                $("#admin-role").html(
                    can.view('views_admin_admin-role_stache', showRole.rolesData)
                );
                }
            );
        },
        reload: function(page) {
            util.logDebug('reload','performed');
            if(page == undefined) {
                can.route.attr({'route': 'admin/:tab/:page', 'tab': 'role', 'page': 1});
            }
            Role.findAll({'page': page}, function (rolesData) {
                    showRole.rolesData.attr('hasPrev',rolesData.hasPrev);
                    showRole.rolesData.attr('hasNext',rolesData.hasNext);
                    showRole.rolesData.attr('prevPage',rolesData.currentPage-1);
                    showRole.rolesData.attr('nextPage',rolesData.currentPage+1);
                    util.refreshArray(showRole.rolesData.roles, rolesData.roles);
                    util.logDebug('roles','reloaded');
                    createRole.initForm();
                }
            );
        }
    });
    /**
     * Control for CreateRole
     * @private
     * @author dorajistyle
     * @param {string} target
     * @name admin#CreateRole
     * @constructor
     */
    var CreateRole = can.Control.extend({
        init: function () {
            util.logInfo('*admin.roles/CreateRole', 'Initialized');
        },
        '.create-role click': function (el, ev) {
            ev.preventDefault();
            ev.stopPropagation();
            this.createRole();
            return false;
        },

        '.create-new-role-form click': function (el, ev) {
            ev.preventDefault();
            ev.stopPropagation();
            createRole.initForm();
//            return false;
        },
        initForm: function () {
            var $createBtn = $(createRole.element.find('.create-role'));
            var $updateBtn = $(createRole.element.find('.update-role'));
            var name = createRole.element.find('#name');
            var description = createRole.element.find('#description');
            name.val('');
            description.val('');
            $createBtn.removeClass('uk-hidden');
            $updateBtn.addClass('uk-hidden');
        },

       /**
         * Validate a form.
         * @memberof admin#CreateRole
         */
        validate: function () {
            return validation.minLength('name', 1, 'admin.role.nameValidation')
        },
        /**
         * Perform create role action.
         * @memberof admin#CreateRole
         */
        createRole: function () {
            if (this.validate()) {
                var form = this.element.find('#adminRoleForm');
                var values = can.deparam(form.serialize());
                var $form = $(form);
                if (!$form.data('submitted')) {
                    $form.data('submitted', true);
                    var createBtn = this.element.find('.create-role');
                    createBtn.attr('disabled', 'disabled');
                    can.when(Role.create(values)).then(function (result) {
                        util.logJson('register', result);
                        createBtn.removeAttr('disabled');
                        $form.data('submitted', false);
//                        if(util.isHashNow('admin/role/1')) showRole.reload();
                        showRole.reload();
                        util.showSuccessMsg(i18n.t('admin.role.create.done'));
                    }, function (xhr) {
                        createBtn.removeAttr('disabled');
                        $form.data('submitted', false);
                        util.handleError(xhr);
                        // util.showErrorMsg(i18n.t('admin.role.create.fail'));
                    });
                }
            } else {
                util.showMessages();
            }
        }


    });
    /**
     * Control for CreateRole
     * @private
     * @author dorajistyle
     * @param {string} target
     * @name admin#UpdateRole
     * @constructor
     */
    var UpdateRole = can.Control.extend({
        init: function () {
            util.logInfo('*admin.roles/UpdateRole', 'Initialized');
        },
        '.update-role click': function (el, ev) {
            ev.preventDefault();
            ev.stopPropagation();
            this.updateRole();
            return false;
        },
        '.select-role click': function (el, ev) {
            this.roleId = util.getId(ev);
            this.showForm();
        },

       /**
         * Validate a form.
         * @memberof admin#UPdateRole
         */
        validate: function () {
            return validation.minLength('name', 1, 'admin.role.nameValidation')
        },

        /**
         * Show update form of selected role.
         * @memberof admin#UPdateRole
         */
        showForm: function() {
                var form = this.element.find('#adminRoleForm');
                var name = this.element.find('#name');
                var description = this.element.find('#description');
                var $createBtn = $(this.element.find('.create-role'));
                var $updateBtn = $(this.element.find('.update-role'));
                var $form = $(form);
                $createBtn.addClass('uk-hidden');
                $updateBtn.removeClass('uk-hidden');
                can.when(Role.findOne({id: updateRole.roleId})).then(function (result) {
                    util.logJson('load',result);
                    name.val(result.role.name);
                    description.val(result.role.description);
                });
        },

        /**
         * perform Update role action.
         * @memberof admin#UpdateRole
         */
        updateRole: function () {
            if (this.validate()) {
                var form = this.element.find('#adminRoleForm');
                var values = can.deparam(form.serialize());
                var $form = $(form);
                if (!$form.data('submitted')) {
                    $form.data('submitted', true);
                    var updateBtn = this.element.find('.create-role');
                    updateBtn.attr('disabled', 'disabled');
                    util.logJson('updateRole', values);
                    can.when(Role.update(updateRole.roleId, {name: values.name, description: values.description})).then(function (result) {
                        util.logJson('register', result);
                        updateBtn.removeAttr('disabled');
                        $form.data('submitted', false);
//                        if(util.isHashNow('admin/role/1')) showRole.reload();
                        showRole.reload();
                        util.showSuccessMsg(i18n.t('admin.role.update.done'));
                    }, function (xhr) {
                        updateBtn.removeAttr('disabled');
                        $form.data('submitted', false);
                        util.handleError(xhr);
                        // util.showErrorMsg(i18n.t('admin.role.update.fail'));
                    });
                }
            } else {
                util.showMessages();
            }
        }


    });

    /**
     * Router for Roles.
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
            showRole = undefined;
            createRole = undefined;
            updateRole = undefined;
            util.logInfo('*admin/Router', 'Initialized')

        },
        allocate: function () {
            var $app = util.getFreshDiv('admin-role');
            showRole = new ShowRole($app);
            createRole = new CreateRole($app);
            updateRole = new UpdateRole($app);
            new Destroy($app, {modalDestroyId: modalDestroyId, name: modalDestroyName, view: showRole, model: Role});
        },
        load: function(page) {
            util.allocate(this, showRole);
            if(showRole.rolesData == undefined || page == undefined){
                showRole.load(page);
                return false;
            }
            showRole.reload(page);
        }
    });

    return Router;
});
