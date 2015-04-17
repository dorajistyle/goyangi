package roleService

import (
	"errors"
	"net/http"

	"github.com/dorajistyle/goyangi/db"
	"github.com/dorajistyle/goyangi/model"
	"github.com/dorajistyle/goyangi/util/log"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

// CreateRole creates a role.
func CreateRole(c *gin.Context) (model.Role, int, error) {
	var form RoleForm
	c.BindWith(&form, binding.Form)
	name := form.Name
	description := form.Description
	role := model.Role{Name: name, Description: description}
	if db.ORM.Create(&role).Error != nil {
		return role, http.StatusInternalServerError, errors.New("Role is not created.")
	}
	return role, http.StatusCreated, nil
}

// RetrieveRole retrieves a role.
func RetrieveRole(c *gin.Context) (model.Role, int, error) {
	var role model.Role
	id := c.Params.ByName("id")
	if db.ORM.First(&role, id).RecordNotFound() {
		return role, http.StatusNotFound, errors.New("Role is not found.")
	}
	return role, http.StatusOK, nil
}

// RetrieveRoles retrieves roles.
func RetrieveRoles(c *gin.Context) []model.Role {
	var roles []model.Role
	db.ORM.Find(&roles)
	return roles
}

// UpdateRole updates a role.
func UpdateRole(c *gin.Context) (model.Role, int, error) {
	var role model.Role
	var form RoleForm
	id := c.Params.ByName("id")
	c.BindWith(&form, binding.Form)
	if db.ORM.First(&role, id).RecordNotFound() {
		return role, http.StatusNotFound, errors.New("Role is not found.")
	}
	role.Name = form.Name
	role.Description = form.Description
	if db.ORM.Save(&role).Error != nil {
		return role, http.StatusInternalServerError, errors.New("Role is not updated.")
	}
	return role, http.StatusOK, nil
}

// DeleteRole deletes a role.
func DeleteRole(c *gin.Context) (int, error) {
	log.Debug("deleteRole performed")
	var role model.Role
	id := c.Params.ByName("id")
	if db.ORM.First(&role, id).RecordNotFound() {
		return http.StatusNotFound, errors.New("Role is not found.")
	}
	if db.ORM.Delete(&role).Delete(role).Error != nil {
		return http.StatusInternalServerError, errors.New("Role is not deleted.")
	}
	return http.StatusOK, nil
}
