package form

// RetrieveListForm used when retrieving any kind of list.
type RetrieveListForm struct {
	CurrentPage int `form:"currentPage" binding:"required"`
}
