package commentService

// CreateCommentForm is used when creating a comment.
type CreateCommentForm struct {
	UserId   int64  `form:"userId" binding:"required"`
	ParentId int64  `form:"parentId" binding:"required"`
	Content  string `form:"content" binding:"required"`
}

// CommentForm is used when updating a comment.
type CommentForm struct {
	CommentId int64  `form:"commentId" binding:"required"`
	ParentId  int64  `form:"parentId" binding:"required"`
	Content   string `form:"content"`
}
