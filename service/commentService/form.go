package commentService

// CreateCommentForm is used when creating a comment.
type CreateCommentForm struct {
	Content  string `form:"content" binding:"required"`
	UserId   uint   `form:"userId"`
}

// CommentForm is used when updating a comment.
type CommentForm struct {
	CommentId uint   `form:"commentId" binding:"required"`
	Content   string `form:"content"`
	UserId   uint   `form:"userId" binding:"required"`
}
