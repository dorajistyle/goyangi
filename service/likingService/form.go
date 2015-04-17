package likingService

// CreateLikingForm is used when creating a liking.
type CreateLikingForm struct {
	UserId   int64 `form:"userId" binding:"required"`
	ParentId int64 `form:"parentId" binding:"required"`
}
