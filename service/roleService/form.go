package roleService

// RoleForm is used when creating or updating a role.
type RoleForm struct {
	Name        string `form:"name" binding:"required"`
	Description string `form:"description" binding:"required"`
}
