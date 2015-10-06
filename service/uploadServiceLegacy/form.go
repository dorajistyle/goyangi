package uploadService

// FileForm is used when creating or updating a file.
type FileForm struct {
	Id     int64  `form:"id"`
	UserId int64  `form:"userId"`
	Name   string `form:"name" binding:"required"`
	Size   int    `form:"size" binding:"required"`
}

// FilesForm is used when creating or updating multiple files.
type FilesForm struct {
	Files []FileForm `form:"json" binding:"required"`
}
