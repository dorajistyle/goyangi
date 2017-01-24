package likingService

// CreateLikingForm is used when creating a liking.
type CreateLikingForm struct {
	UserId   uint `form:"userId" binding:"required"`
	ParentId uint `form:"parentId" binding:"required"`
}
