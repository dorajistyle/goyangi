package commentService

// CreateCommentForm is used when creating a comment.
type CreateCommentForm struct {
	UserId   uint   `form:"userId" binding:"required"`
	ParentId uint   `form:"parentId" binding:"required"`
	Content  string `form:"content" binding:"required"`
}

// CommentForm is used when updating a comment.
type CommentForm struct {
	CommentId uint   `form:"commentId" binding:"required"`
	ParentId  uint   `form:"parentId" binding:"required"`
	Content   string `form:"content"`
}
