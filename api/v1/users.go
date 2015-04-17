package v1

import (
	"github.com/gin-gonic/gin"

	"github.com/dorajistyle/goyangi/api/response"
	"github.com/dorajistyle/goyangi/service/userService"
	"github.com/dorajistyle/goyangi/service/userService/userLiking"
	"github.com/dorajistyle/goyangi/service/userService/userPermission"
	"github.com/dorajistyle/goyangi/util/email"
	"github.com/dorajistyle/goyangi/util/log"
)

// @Title Users
// @Description Users's router group.
func Users(parentRoute *gin.RouterGroup) {
	route := parentRoute.Group("/users")
	route.POST("/", createUser)
	route.GET("/:id", retrieveUser)
	route.GET("/", retrieveUsers)
	route.PUT("/:id", userPermission.AuthRequired(updateUser))
	route.DELETE("/:id", userPermission.AuthRequired(deleteUser))

	route.POST("/roles", userPermission.AdminRequired(addRoleToUser))
	route.DELETE(":id/roles/:roleId", userPermission.AdminRequired(removeRoleFromUser))

	route.POST("/likings", userPermission.AuthRequired(createLikingOnUser))
	route.GET("/:id/likings", retrieveLikingsOnUsers)
	route.DELETE("/:id/likings/:userId", userPermission.AuthRequired(deleteLikingOnUser))
	route.GET("/:id/liked", retrieveLikedOnUsers)

	route = parentRoute.Group("/user")
	route.GET("/current", retrieveCurrentUser)
	route.POST("/send/password/reset/token", sendPasswordResetToken)
	route.PUT("/reset/password", resetPassword)
	route.POST("/send/email/verification/token", sendEmailVerificationToken)
	route.PUT("/verify/email", verifyEmail)
	route.GET("/email/:email", retrieveUserByEmail)
	route.GET("/email/:email/list", retrieveUsersByEmail)
	route.GET("/username/:username", retrieveUserByUsername)
	route.GET("/admin/:id", userPermission.AdminRequired(retrieveUserForAdmin))
	route.GET("/admin", userPermission.AdminRequired(retrieveUsersForAdmin))
	route.PUT("/activate/:id", userPermission.AdminRequired(activateUser))
	route.GET("/test/send/email", sendTestEmail)
}

// @Title createUser
// @Description Create a user.
// @Accept  json
// @Param   registrationEmail        form   string     true        "User Email."
// @Param   registrationPassword        form   string  true        "User Password."
// @Success 201 {object} response.BasicResponse "User is registered successfully"
// @Failure 401 {object} response.BasicResponse "Authentication required"
// @Failure 404 {object} response.BasicResponse "User not logged in."
// @Failure 500 {object} response.BasicResponse "User is not created."
// @Resource /users
// @Router /users [post]
func createUser(c *gin.Context) {
	status, err := userService.CreateUser(c)
	messageTypes := &response.MessageTypes{
		OK:                  "registration.done",
		Unauthorized:        "login.error.fail",
		NotFound:            "registration.error.fail",
		InternalServerError: "registration.error.fail",
	}
	messages := &response.Messages{OK: "User is registered successfully."}
	response.JSON(c, status, messageTypes, messages, err)
}

// @Title retrieveUser
// @Description Retrieve a user.
// @Accept  json
// @Param   id        path    int     true        "User ID"
// @Success 200 {object} model.PublicUser "OK"
// @Failure 404 {object} response.BasicResponse "Not found"
// @Resource /users
// @Router /users/{id} [get]
func retrieveUser(c *gin.Context) {
	user, isAuthor, currentUserId, status, err := userService.RetrieveUser(c)
	if err == nil {
		c.JSON(status, gin.H{"user": user, "isAuthor": isAuthor, "currentUserId": currentUserId})
	} else {
		messageTypes := &response.MessageTypes{
			NotFound: "user.error.notFound",
		}
		response.ErrorJSON(c, status, messageTypes, err)
	}

}

// @Title retrieveUsers
// @Description Retrieve user array.
// @Accept  json
// @Success 200 {array} model.PublicUser "OK"
// @Resource /users
// @Router /users [get]
func retrieveUsers(c *gin.Context) {
	users := userService.RetrieveUsers(c)
	c.JSON(200, gin.H{"users": users})
}

// @Title updateUser
// @Description Update a user.
// @Accept  json
// @Param   id        path    int     true        "User ID"
// @Success 200 {object} model.User "OK"
// @Failure 401 {object} response.BasicResponse "Authentication required"
// @Failure 404 {object} response.BasicResponse "Not found"
// @Failure 500 {object} response.BasicResponse "User is not updated."
// @Resource /users
// @Router /users/{id} [put]
func updateUser(c *gin.Context) {
	user, status, err := userService.UpdateUser(c)
	if err == nil {
		c.JSON(status, gin.H{"user": user})
	} else {
		messageTypes := &response.MessageTypes{
			Unauthorized:        "user.error.unauthorized",
			NotFound:            "user.error.notFound",
			InternalServerError: "user.error.internalServerError",
		}
		response.ErrorJSON(c, status, messageTypes, err)
	}
}

// @Title deleteUser
// @Description Delete a user.
// @Accept  json
// @Param   id        path    int     true        "User ID"
// @Success 200 {object} response.BasicResponse
// @Failure 401 {object} response.BasicResponse "Authentication required"
// @Failure 404 {object} response.BasicResponse "Not found"
// @Failure 500 {object} response.BasicResponse "User is not deleted."
// @Resource /users
// @Router /users/{id} [delete]
func deleteUser(c *gin.Context) {
	status, err := userService.DeleteUser(c)
	if err == nil {
		c.JSON(status, response.BasicResponse{})
	} else {
		messageTypes := &response.MessageTypes{
			Unauthorized:        "user.error.unauthorized",
			NotFound:            "user.error.notFound",
			InternalServerError: "setting.leaveOurService.fail",
		}
		response.ErrorJSON(c, status, messageTypes, err)
	}

}

// @Title addRoleToUser
// @Description Add a role to user.
// @Accept  json
// @Param   userId        form   int  true        "User ID."
// @Param   roleId        form   int  true        "Role ID."
// @Success 201 {object} response.BasicResponse "Created"
// @Failure 401 {object} response.BasicResponse "Authentication required"
// @Failure 404 {object} response.BasicResponse "User or role is not found"
// @Failure 500 {object} response.BasicResponse "Role is not added to a user"
// @Resource /users
// @Router /users/roles [post]
func addRoleToUser(c *gin.Context) {
	status, err := userService.AddRoleToUser(c)
	messageTypes := &response.MessageTypes{
		OK:                  "admin.user.role.add.done",
		Unauthorized:        "user.error.unauthorized",
		NotFound:            "user.error.notFound",
		InternalServerError: "admin.user.role.add.fail",
	}
	messages := &response.Messages{OK: "Role is added to a user successfully."}
	response.JSON(c, status, messageTypes, messages, err)
}

// @Title removeRoleFromUser
// @Description Remove a role from user.
// @Accept  json
// @Param   id        path   int  true        "User ID."
// @Param   roleId        form   int  true        "Role ID."
// @Success 200 {object} response.BasicResponse "OK"
// @Failure 401 {object} response.BasicResponse "Authentication required"
// @Failure 404 {object} response.BasicResponse "Not found"
// @Failure 500 {object} response.BasicResponse "Role is not deleted from a user"
// @Resource /users
// @Router /users/{id}/roles/{roleId} [delete]
func removeRoleFromUser(c *gin.Context) {
	status, err := userService.RemoveRoleFromUser(c)
	messageTypes := &response.MessageTypes{
		OK:                  "admin.user.role.delete.done",
		Unauthorized:        "user.error.unauthorized",
		NotFound:            "user.error.notFound",
		InternalServerError: "admin.user.role.delete.fail",
	}
	messages := &response.Messages{OK: "Role is deleted from a user successfully."}
	response.JSON(c, status, messageTypes, messages, err)
}

// @Title createLikingOnUser
// @Description Create a liking on a user.
// @Accept  json
// @Param   parentId       form   int     true        "Parent item id."
// @Param   userId        form   int     true        "User id."
// @Success 201 {object} model.PublicUser "Created"
// @Failure 401 {object} response.BasicResponse "Authentication required"
// @Failure 404 {object} response.BasicResponse "Not found"
// @Failure 500 {object} response.BasicResponse "Liking is not created"
// @Resource /users
// @Router /users/likings [post]
func createLikingOnUser(c *gin.Context) {
	status, err := userLiking.CreateLikingOnUser(c)

	messageTypes := &response.MessageTypes{
		OK:                  "liking.like.done",
		Unauthorized:        "liking.error.unauthorized",
		NotFound:            "liking.error.notFound",
		InternalServerError: "liking.like.fail",
	}
	messages := &response.Messages{OK: "Liking is created successfully."}
	response.JSON(c, status, messageTypes, messages, err)
}

// @Title retrieveLikingsOnUsers
// @Description Retrieve likings on a user.
// @Accept  json
// @Param   userId        path    int     true        "User ID"
// @Success 200 {array} model.PublicUser "OK"
// @Failure 404 {object} response.BasicResponse "Not found"
// @Resource /users
// @Router /users/{id}/likings [get]
func retrieveLikingsOnUsers(c *gin.Context) {
	likings, currentPage, hasPrev, hasNext, status, err := userLiking.RetrieveLikingsOnUser(c)
	if err == nil {
		c.JSON(status, gin.H{"likings": likings, "currentPage": currentPage,
			"hasPrev": hasPrev, "hasNext": hasNext})
	} else {
		messageTypes := &response.MessageTypes{
			Unauthorized: "liking.error.unauthorized",
			NotFound:     "liking.error.notFound",
		}
		response.ErrorJSON(c, status, messageTypes, err)
	}
}

// @Title retrieveLikedOnUsers
// @Description Retrieve likings on a user.
// @Accept  json
// @Param   userId        path    int     true        "User ID"
// @Success 200 {array} model.PublicUser "OK"
// @Failure 401 {object} response.BasicResponse "Authentication required"
// @Failure 404 {object} response.BasicResponse "Not found"
// @Resource /users
// @Router /users/{id}/liked [get]
func retrieveLikedOnUsers(c *gin.Context) {
	liked, currentPage, hasPrev, hasNext, status, err := userLiking.RetrieveLikedOnUser(c)
	if err == nil {
		c.JSON(status, gin.H{"liked": liked, "currentPage": currentPage,
			"hasPrev": hasPrev, "hasNext": hasNext})
	} else {
		messageTypes := &response.MessageTypes{
			Unauthorized: "liking.error.unauthorized",
			NotFound:     "liking.error.notFound",
		}
		response.ErrorJSON(c, status, messageTypes, err)
	}
}

// @Title deleteLikingOnUser
// @Description Delete a liking on user.
// @Accept  json
// @Param   userId        path    int     true        "User ID"
// @Param   id                path    int     true        "Liking ID"
// @Success 200 {object} response.BasicResponse
// @Failure 401 {object} response.BasicResponse "Authentication required"
// @Failure 404 {object} response.BasicResponse "Not found"
// @Failure 500 {object} response.BasicResponse "Liking is not deleted"
// @Resource /users
// @Router /users/{id}/likings/{likingId} [delete]
func deleteLikingOnUser(c *gin.Context) {
	status, err := userLiking.DeleteLikingOnUser(c)
	messageTypes := &response.MessageTypes{
		OK:                  "liking.unlike.done",
		Unauthorized:        "liking.error.unauthorized",
		NotFound:            "liking.error.notFound",
		InternalServerError: "liking.unlike.fail",
	}
	messages := &response.Messages{OK: "Liking is deleted successfully."}
	response.JSON(c, status, messageTypes, messages, err)
}

// @Title retrieveCurrentUser
// @Description Retrieve the current user.
// @Accept  json
// @Success 200 {object} model.User "OK"
// @Failure 401 {object} response.BasicResponse "Authentication required"
// @Failure 404 {object} response.BasicResponse "Not found"
// @Failure 500 {object} response.BasicResponse "Liking is not deleted"
// @Resource /user
// @Router /user/current [get]
func retrieveCurrentUser(c *gin.Context) {
	user, status, err := userService.RetrieveCurrentUser(c)
	// if hasUser := log.CheckError(err); hasUser {
	if err == nil {
		c.JSON(status, gin.H{"hasAdmin": userPermission.HasAdmin(&user), "user": user})
	} else {
		c.JSON(200, gin.H{"hasAdmin": false, "user": nil})
	}

}

// @Title sendPasswordResetToken
// @Description Create a user session.
// @Accept  json
// @Param   email        form   string     true        "User email."
// @Success 200 {object} response.BasicResponse "Sent"
// @Failure 401 {object} response.BasicResponse "Authentication required"
// @Failure 500 {object} response.BasicResponse "Password reset token not sent"
// @Resource /user
// @Router /user/send/password/reset/token [post]
func sendPasswordResetToken(c *gin.Context) {
	status, err := userService.SendPasswordResetToken(c)
	messageTypes := &response.MessageTypes{
		OK:                  "passwordReset.send.sent.done",
		Unauthorized:        "error.unauthorized",
		NotFound:            "error.notFound",
		InternalServerError: "passwordReset.send.sent.fail",
	}
	messages := &response.Messages{OK: "Password reset token sent successfully."}
	response.JSON(c, status, messageTypes, messages, err)
}

// @Title resetPassword
// @Description Create a user session.
// @Accept  json
// @Param   token        form   string     true        "User password reset token"
// @Param   newPassword        form   string     true        "New password"
// @Success 200 {object} response.BasicResponse "User password updated"
// @Failure 400 {object} response.BasicResponse "User password is not updated."
// @Failure 401 {object} response.BasicResponse "Authentication required"
// @Failure 500 {object} response.BasicResponse "Password reset token not sent"
// @Resource /user
// @Router /user/reset/password [put]
func resetPassword(c *gin.Context) {
	status, err := userService.ResetPassword(c)
	messageTypes := &response.MessageTypes{
		OK:                  "passwordReset.reset.updated.done",
		Forbidden:           "passwordReset.error.tokenExpired",
		NotFound:            "error.notFound",
		InternalServerError: "passwordReset.reset.updated.fail",
	}
	messages := &response.Messages{OK: "Password reset successfully."}
	response.JSON(c, status, messageTypes, messages, err)
}

// @Title sendEmailVerificationToken
// @Description Create a user session.
// @Accept  json
// @Success 201 {object} response.BasicResponse "Created"
// @Failure 401 {object} response.BasicResponse "Authentication required"
// @Failure 500 {object} response.BasicResponse "Email verification token not sent"
// @Resource /user
// @Router /user/send/email/verification/token [post]
func sendEmailVerificationToken(c *gin.Context) {
	status, err := userService.SendVerification(c)
	messageTypes := &response.MessageTypes{
		OK:                  "emailVerification.send.sent.done",
		Unauthorized:        "error.unauthorized",
		NotFound:            "error.notFound",
		InternalServerError: "emailVerification.send.sent.fail",
	}
	messages := &response.Messages{OK: "Email verification token sent successfully."}
	response.JSON(c, status, messageTypes, messages, err)
}

// @Title verifyEmail
// @Description Create a user session.
// @Accept  json
// @Param   token        form   string     true        "User email validation token"
// @Success 200 {object} response.BasicResponse "User email verified."
// @Failure 400 {object} response.BasicResponse "User email not verified."
// @Resource /user
// @Router /user/verify/email [put]
func verifyEmail(c *gin.Context) {
	status, err := userService.EmailVerification(c)
	messageTypes := &response.MessageTypes{
		OK:                  "emailVerification.done",
		Forbidden:           "emailVerification.error.tokenExpired",
		NotFound:            "error.notFound",
		InternalServerError: "emailVerification.fail",
	}
	messages := &response.Messages{OK: "Email verified successfully."}
	response.JSON(c, status, messageTypes, messages, err)
}

// @Title retrieveUserByEmail
// @Description Retrieve user by email.
// @Accept  json
// @Param   email        path    string     true        "User email"
// @Success 200 {object} model.PublicUser "OK"
// @Failure 404 {object} response.BasicResponse "Not found"
// @Resource /user
// @Router /user/email/{email} [get]
func retrieveUserByEmail(c *gin.Context) {
	user, email, status, err := userService.RetrieveUserByEmail(c)
	if err == nil {
		c.JSON(status, gin.H{"user": user, "email": email})
	} else {
		c.JSON(status, gin.H{"user": nil, "email": email})
	}
}

// @Title retrieveUsersByEmail
// @Description Retrieve user array by email.
// @Accept  json
// @Param   email        path    string     true        "User email"
// @Success 200 {array} model.PublicUser "OK"
// @Resource /user
// @Router /user/email/{email}/list [get]
func retrieveUsersByEmail(c *gin.Context) {
	users := userService.RetrieveUsersByEmail(c)
	c.JSON(200, gin.H{"users": users})
}

// @Title retrieveUserByUsername
// @Description Retrieve user by username.
// @Accept  json
// @Param   username        path    string     true        "User email"
// @Success 200 {object} model.PublicUser "OK"
// @Failure 404 {object} response.BasicResponse "Not found"
// @Resource /user
// @Router /user/username/{username} [get]
func retrieveUserByUsername(c *gin.Context) {
	user, username, status, err := userService.RetrieveUserByUsername(c)
	if err == nil {
		c.JSON(status, gin.H{"user": user, "username": username})
	} else {
		c.JSON(status, gin.H{"user": nil, "username": username})
	}
}

// @Title retrieveCurrentUser
// @Description Retrieve a user for admin. It contains more information than normal query.
// @Accept  json
// @Param   id        path    int     true        "User ID"
// @Success 200 {object} model.User "OK"
// @Failure 401 {object} response.BasicResponse "Authentication required"
// @Failure 404 {object} response.BasicResponse "Not found"
// @Resource /user
// @Router /user/admin/{id} [get]
func retrieveUserForAdmin(c *gin.Context) {
	user, status, err := userService.RetrieveUserForAdmin(c)
	if err == nil {
		c.JSON(status, gin.H{"user": user})
	} else {
		messageTypes := &response.MessageTypes{
			Unauthorized: "user.error.unauthorized",
			NotFound:     "user.error.notFound",
		}
		response.ErrorJSON(c, status, messageTypes, err)
	}

}

// @Title retrieveUsersForAdmin
// @Description Retrieve user array for admin. It contains more information than normal query.
// @Accept  json
// @Success 200 {array} model.User "OK"
// @Failure 401 {object} response.BasicResponse "Authentication required"
// @Resource /user
// @Router /user/admin [get]
func retrieveUsersForAdmin(c *gin.Context) {
	users := userService.RetrieveUsersForAdmin(c)
	c.JSON(200, gin.H{"users": users})
}

// @Title activateUser
// @Description Activate/Deactivate a user.
// @Accept  json
// @Param   id        path    int     true        "User ID"
// @Success 200 {object} model.User "OK"
// @Failure 401 {object} response.BasicResponse "Authentication required"
// @Failure 404 {object} response.BasicResponse "Not found"
// @Resource /user
// @Router /users/{id} [put]
func activateUser(c *gin.Context) {
	user, status, err := userService.ActivateUser(c)
	if err == nil {
		c.JSON(status, gin.H{"user": user})
	} else {
		messageTypes := &response.MessageTypes{
			Unauthorized:        "user.error.unauthorized",
			NotFound:            "user.error.notFound",
			InternalServerError: "admin.user.toggleActivate.fail",
		}
		response.ErrorJSON(c, status, messageTypes, err)
	}
}

// @Title sendTestEmail
// @Description send a test email.
// @Accept  json
// @Success 200 {object} response.BasicResponse "OK"
// @Resource /user
// @Router /user/test/send/email [get]
func sendTestEmail(c *gin.Context) {
	log.Debugf("sendmail %v", c.Params)
	email.SendTestEmail()
	c.JSON(200, response.BasicResponse{})
}
