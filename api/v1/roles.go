package v1

import (
	"github.com/gin-gonic/gin"

	"github.com/dorajistyle/goyangi/api/response"
	"github.com/dorajistyle/goyangi/service/roleService"
	"github.com/dorajistyle/goyangi/service/userService/userPermission"
	"github.com/dorajistyle/goyangi/util/log"
)

// @Title Roles
// @Description Roles's router group.
func Roles(parentRoute *gin.RouterGroup) {
	route := parentRoute.Group("/roles")
	route.POST("/", userPermission.AdminRequired(createRole))
	route.GET("/:id", retrieveRole)
	route.GET("/", retrieveRoles)
	route.PUT("/:id", userPermission.AdminRequired(updateRole))
	route.DELETE("/:id", userPermission.AdminRequired(deleteRole))
}

// @Title createRole
// @Description Create a role.
// @Accept  json
// @Param   name        form   string     true        "Name of Role."
// @Param   description        form   string  true        "Description of Role."
// @Success 201 {object} model.Role "Created"
// @Failure 401 {object} response.BasicResponse "Authentication required"
// @Failure 500 {object} response.BasicResponse "Role is not created"
// @Resource /roles
// @Router /roles [post]
func createRole(c *gin.Context) {
	role, status, err := roleService.CreateRole(c)
	if err == nil {
		c.JSON(status, gin.H{"role": role})
	} else {
		messageTypes := &response.MessageTypes{Unauthorized: "role.error.unauthorized",
			InternalServerError: "admin.role.create.fail"}
		response.ErrorJSON(c, status, messageTypes, err)
		// c.JSON(400, gin.H{"role": nil})
	}
}

// @Title retrieveRole
// @Description Retrieve a role.
// @Accept  json
// @Param   id        path    int     true        "Role ID"
// @Success 200 {object} model.Role "OK"
// @Failure 404 {object} response.BasicResponse "Not found"
// @Resource /roles
// @Router /roles/{id} [get]
func retrieveRole(c *gin.Context) {
	role, status, err := roleService.RetrieveRole(c)
	if err == nil {
		c.JSON(status, gin.H{"role": role})
	} else {
		messageTypes := &response.MessageTypes{NotFound: "role.error.notFound"}
		response.ErrorJSON(c, status, messageTypes, err)
	}
}

// @Title retrieveRoles
// @Description Retrieve role array.
// @Accept  json
// @Success 200 {array} model.Role "OK"
// @Resource /roles
// @Router /roles [get]
func retrieveRoles(c *gin.Context) {
	roles := roleService.RetrieveRoles(c)
	c.JSON(200, gin.H{"roles": roles})
}

// @Title updateRole
// @Description Update a role.
// @Accept  json
// @Param   id        path    int     true        "Role ID"
// @Success 200 {object} model.Role "OK"
// @Failure 401 {object} response.BasicResponse "Authentication required"
// @Failure 404 {object} response.BasicResponse "Not found"
// @Failure 500 {object} response.BasicResponse "Role is not updated"
// @Resource /roles
// @Router /roles/{id} [put]
func updateRole(c *gin.Context) {
	role, status, err := roleService.UpdateRole(c)
	if err == nil {
		c.JSON(status, gin.H{"role": role})
	} else {
		messageTypes := &response.MessageTypes{Unauthorized: "role.error.unauthorized",
			NotFound:            "role.error.notFound",
			InternalServerError: "admin.role.update.fail"}
		response.ErrorJSON(c, status, messageTypes, err)
		// c.JSON(400, gin.H{"role": nil})
	}

}

// @Title deleteRole
// @Description Delete a role.
// @Accept  json
// @Param   id        path    int     true        "Role ID"
// @Success 200 {object} response.BasicResponse
// @Failure 401 {object} response.BasicResponse "Authentication required"
// @Failure 404 {object} response.BasicResponse "Not found"
// @Resource /roles
// @Router /roles/{id} [delete]
func deleteRole(c *gin.Context) {
	log.Debug("deleteRole performed")
	status, err := roleService.DeleteRole(c)
	messageTypes := &response.MessageTypes{
		OK:           "destroy.done",
		BadRequest:   "destroy.fail",
		Unauthorized: "role.error.unauthorized",
		NotFound:     "role.error.notFound"}
	messages := &response.Messages{OK: "Role is deleted successfully."}
	response.JSON(c, status, messageTypes, messages, err)
	//
	// if err == nil {
	// 	c.JSON(status, response.BasicResponse{})
	// } else {
	// 	c.JSON(400, response.BasicResponse{})
	// }
}
