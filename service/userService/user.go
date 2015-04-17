package userService

import (
	"errors"
	"net/http"
	"strconv"

	"golang.org/x/crypto/bcrypt"

	"github.com/dorajistyle/goyangi/config"
	"github.com/dorajistyle/goyangi/db"
	"github.com/dorajistyle/goyangi/model"
	"github.com/dorajistyle/goyangi/service/likingService/likingMeta"
	"github.com/dorajistyle/goyangi/service/likingService/likingRetriever"
	"github.com/dorajistyle/goyangi/util/crypto"
	"github.com/dorajistyle/goyangi/util/log"
	"github.com/dorajistyle/goyangi/util/modelHelper"
	"github.com/dorajistyle/goyangi/util/timeHelper"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

var userFields []string = []string{"name", "email", "createdAt", "updatedAt"}

// SuggestUsername suggest user's name if user's name already occupied.
func SuggestUsername(username string) string {
	var count int
	var usernameCandidate string
	db.ORM.Model(model.User{}).Where(&model.User{Username: username}).Count(&count)
	log.Debugf("count Before : %d", count)
	if count == 0 {
		return username
	} else {
		var postfix int
		for {
			usernameCandidate = username + strconv.Itoa(postfix)
			log.Debugf("usernameCandidate: %s\n", usernameCandidate)
			db.ORM.Model(model.User{}).Where(&model.User{Username: usernameCandidate}).Count(&count)
			log.Debugf("count after : %d\n", count)
			postfix = postfix + 1
			if count == 0 {
				break
			}
		}
	}
	return usernameCandidate
}

// CreateUserFromForm creates a user from a registration form.
func CreateUserFromForm(registrationForm RegistrationForm) (model.User, error) {
	var user model.User
	log.Debugf("registrationForm %+v\n", registrationForm)
	modelHelper.AssignValue(&user, &registrationForm)
	user.Md5 = crypto.GenerateMD5Hash(user.Email)
	token, err := crypto.GenerateRandomToken32()
	if err != nil {
		return user, errors.New("Token not generated.")
	}
	user.Token = token
	user.TokenExpiration = timeHelper.FewDaysLater(config.AuthTokenExpirationDay)
	log.Debugf("user %+v\n", user)
	if db.ORM.Create(&user).Error != nil {
		return user, errors.New("User is not created.")
	}
	return user, nil
}

// CreateUser creates a user.
func CreateUser(c *gin.Context) (int, error) {
	var user model.User
	var registrationForm RegistrationForm
	var status int
	var err error

	c.BindWith(&registrationForm, binding.Form)

	password, err := bcrypt.GenerateFromPassword([]byte(registrationForm.Password), 10)
	if err != nil {
		return http.StatusInternalServerError, err
	}
	registrationForm.Password = string(password)
	user, err = CreateUserFromForm(registrationForm)
	if err != nil {
		return http.StatusInternalServerError, err
	}
	SendVerificationToUser(user)
	status, err = RegisterHandler(c)
	return status, err
}

// RetrieveUser retrieves a user.
func RetrieveUser(c *gin.Context) (*model.PublicUser, bool, int64, int, error) {
	var user model.User
	var currentUserId int64
	var isAuthor bool
	// var publicUser *model.PublicUser
	// publicUser.User = &user
	id := c.Params.ByName("id")
	if db.ORM.Select(config.UserPublicFields).First(&user, id).RecordNotFound() {
		return &model.PublicUser{User: &user}, isAuthor, currentUserId, http.StatusNotFound, errors.New("User is not found.")
	}
	currentUser, err := CurrentUser(c)
	if err == nil {
		currentUserId = currentUser.Id
		isAuthor = currentUser.Id == user.Id
	}
	var likingList model.LikingList
	likings, currentPage, hasPrev, hasNext, _ := likingRetriever.RetrieveLikings(user)
	likingList.Likings = likings
	currentUserlikedCount := db.ORM.Model(&user).Where("id =?", currentUserId).Association("Likings").Count()
	log.Debugf("Current user like count : %d", currentUserlikedCount)
	likingMeta.SetLikingPageMeta(&likingList, currentPage, hasPrev, hasNext, user.LikingCount, currentUserlikedCount)
	user.LikingList = likingList
	var likedList model.LikedList
	liked, currentPage, hasPrev, hasNext, _ := likingRetriever.RetrieveLiked(user)
	likedList.Liked = liked
	likingMeta.SetLikedPageMeta(&likedList, currentPage, hasPrev, hasNext, user.LikedCount)
	user.LikedList = likedList
	return &model.PublicUser{User: &user}, isAuthor, currentUserId, http.StatusOK, nil
}

// RetrieveUsers retrieves users.
func RetrieveUsers(c *gin.Context) []*model.PublicUser {
	var users []*model.User
	var userArr []*model.PublicUser
	db.ORM.Select(config.UserPublicFields).Find(&users)
	for _, user := range users {
		userArr = append(userArr, &model.PublicUser{User: user})
	}
	return userArr
}

// UpdateUserCore updates a user. (Applying the modifed data of user).
func UpdateUserCore(user *model.User) (int, error) {
	user.Md5 = crypto.GenerateMD5Hash(user.Email)
	token, err := crypto.GenerateRandomToken32()
	if err != nil {
		return http.StatusInternalServerError, errors.New("Token not generated.")
	}
	user.Token = token
	user.TokenExpiration = timeHelper.FewDaysLater(config.AuthTokenExpirationDay)
	if db.ORM.Save(user).Error != nil {
		return http.StatusInternalServerError, errors.New("User is not updated.")
	}
	return http.StatusOK, nil
}

// UpdateUser updates a user.
func UpdateUser(c *gin.Context) (*model.User, int, error) {
	id := c.Params.ByName("id")
	var user model.User
	if db.ORM.First(&user, id).RecordNotFound() {
		return &user, http.StatusNotFound, errors.New("User is not found.")
	}
	switch c.Request.FormValue("type") {
	case "password":
		var passwordForm PasswordForm
		c.BindWith(&passwordForm, binding.Form)
		log.Debugf("form %+v\n", passwordForm)
		err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(passwordForm.CurrentPassword))
		if err != nil {
			log.Error("Password Incorrect.")
			return &user, http.StatusInternalServerError, errors.New("User is not updated. Password Incorrect.")
		} else {
			newPassword, err := bcrypt.GenerateFromPassword([]byte(passwordForm.Password), 10)
			if err != nil {
				return &user, http.StatusInternalServerError, errors.New("User is not updated. Password not Generated.")
			} else {
				passwordForm.Password = string(newPassword)
				modelHelper.AssignValue(&user, &passwordForm)
			}
		}
	default:
		var form UserForm
		c.BindWith(&form, binding.Form)
		log.Debugf("form %+v\n", form)
		modelHelper.AssignValue(&user, &form)
	}

	log.Debugf("params %+v\n", c.Params)
	status, err := UpdateUserCore(&user)
	if err != nil {
		return &user, status, err
	}
	status, err = SetCookie(c, user.Token)
	return &user, status, err
}

// DeleteUser deletes a user.
func DeleteUser(c *gin.Context) (int, error) {
	id := c.Params.ByName("id")
	var user model.User
	if db.ORM.First(&user, id).RecordNotFound() {
		return http.StatusNotFound, errors.New("User is not found.")
	}
	if db.ORM.Delete(&user).Error != nil {
		return http.StatusInternalServerError, errors.New("User is not deleted.")
	}
	status, err := ClearCookie(c)
	return status, err
}

// AddRoleToUser adds a role to a user.
func AddRoleToUser(c *gin.Context) (int, error) {
	var form UserRoleForm
	var user model.User
	var role model.Role
	var roles []model.Role
	c.BindWith(&form, binding.Form)

	if db.ORM.First(&user, form.UserId).RecordNotFound() {
		return http.StatusNotFound, errors.New("User is not found.")
	}
	if db.ORM.First(&role, form.RoleId).RecordNotFound() {
		return http.StatusNotFound, errors.New("Role is not found.")
	}
	log.Debugf("user email : %s", user.Email)
	log.Debugf("Role name : %s", role.Name)
	db.ORM.Model(&user).Association("Roles").Append(role)
	db.ORM.Model(&user).Association("Roles").Find(&roles)
	if db.ORM.Save(&user).Error != nil {
		return http.StatusInternalServerError, errors.New("Role not appended to user.")
	}
	return http.StatusOK, nil
}

// RemoveRoleFromUser removes a role from a user.
func RemoveRoleFromUser(c *gin.Context) (int, error) {
	var user model.User
	var role model.Role
	log.Debugf("params : %v\n", c.Params)
	userId := c.Params.ByName("id")
	roleId := c.Params.ByName("roleId")
	if db.ORM.First(&user, userId).RecordNotFound() {
		return http.StatusNotFound, errors.New("User is not found.")
	}
	if db.ORM.First(&role, roleId).RecordNotFound() {
		return http.StatusNotFound, errors.New("Role is not found.")
	}

	log.Debugf("user : %v\n", user)
	log.Debugf("role : %v\n", role)
	if db.ORM.Model(&user).Association("Roles").Delete(role).Error != nil {
		return http.StatusInternalServerError, errors.New("Role is not deleted from user.")
	}
	return http.StatusOK, nil
}

// RetrieveCurrentUser retrieves a current user.
func RetrieveCurrentUser(c *gin.Context) (model.User, int, error) {
	user, err := CurrentUser(c)
	if err != nil {
		return user, http.StatusInternalServerError, err
	}
	return user, http.StatusOK, nil
}

// RetrieveUserByEmail retrieves a user by an email
func RetrieveUserByEmail(c *gin.Context) (*model.PublicUser, string, int, error) {
	email := c.Params.ByName("email")
	var user model.User
	if db.ORM.Unscoped().Select(config.UserPublicFields).Where("email like ?", "%"+email+"%").First(&user).RecordNotFound() {
		return &model.PublicUser{User: &user}, email, http.StatusNotFound, errors.New("User is not found.")
	}
	return &model.PublicUser{User: &user}, email, http.StatusOK, nil
}

// RetrieveUsersByEmail retrieves users by an email
func RetrieveUsersByEmail(c *gin.Context) []*model.PublicUser {
	var users []*model.User
	var userArr []*model.PublicUser
	email := c.Params.ByName("email")
	db.ORM.Select(config.UserPublicFields).Where("email like ?", "%"+email+"%").Find(&users)
	for _, user := range users {
		userArr = append(userArr, &model.PublicUser{User: user})
	}
	return userArr
}

// RetrieveUserByUsername retrieves a user by username.
func RetrieveUserByUsername(c *gin.Context) (*model.PublicUser, string, int, error) {
	username := c.Params.ByName("username")
	var user model.User
	if db.ORM.Unscoped().Select(config.UserPublicFields).Where("username like ?", "%"+username+"%").First(&user).RecordNotFound() {
		return &model.PublicUser{User: &user}, username, http.StatusNotFound, errors.New("User is not found.")
	}
	return &model.PublicUser{User: &user}, username, http.StatusOK, nil
}

// RetrieveUserForAdmin retrieves a user for an administrator.
func RetrieveUserForAdmin(c *gin.Context) (model.User, int, error) {
	id := c.Params.ByName("id")
	var user model.User
	if db.ORM.First(&user, id).RecordNotFound() {
		return user, http.StatusNotFound, errors.New("User is not found.")
	}
	db.ORM.Model(&user).Association("Languages").Find(&user.Languages)
	db.ORM.Model(&user).Association("Roles").Find(&user.Roles)
	return user, http.StatusOK, nil
}

// RetrieveUsersForAdmin retrieves users for an administrator.
func RetrieveUsersForAdmin(c *gin.Context) []model.User {
	var users []model.User
	var userArr []model.User
	db.ORM.Find(&users)
	for _, user := range users {
		db.ORM.Model(&user).Association("Languages").Find(&user.Languages)
		db.ORM.Model(&user).Association("Roles").Find(&user.Roles)
		userArr = append(userArr, user)
	}
	return userArr
}

// ActivateUser toggle activation of a user.
func ActivateUser(c *gin.Context) (model.User, int, error) {
	id := c.Params.ByName("id")
	var user model.User
	var form ActivateForm
	c.BindWith(&form, binding.Form)
	if db.ORM.First(&user, id).RecordNotFound() {
		return user, http.StatusNotFound, errors.New("User is not found.")
	}
	user.Activation = form.Activation
	if db.ORM.Save(&user).Error != nil {
		return user, http.StatusInternalServerError, errors.New("User not activated.")
	}
	return user, http.StatusOK, nil
}
