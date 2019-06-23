package articleService

// ArticleFilter is a filter for retriving articles.
type ArticleFilter struct {
	UserId         int   `form:"userId" json:"userId"`
	Categories     []int `form:"categories" json:"categories"`
	CurrentPage    int   `form:"currentPage" json:"currentPage"`
	ArticlePerPage int   `form:"articlePerPage" json:"articlePerPage"`
}

// ArticleForm is a form of article.
type ArticleForm struct {
	Id            uint   `form:"id"`
	UserId        uint   `form:"userId"`
	CategoryId    int    `form:"categoryId"`
	Title         string `form:"title" binding:"required"`
	Content       string `form:"content"  binding:"required"`
	Url           string `form:"url"`
	ImageName     string `form:"imageName"`
	ThumbnailName string `form:"thumbnailName"`
}

// ArticlesForm is used when creating or updating multiple Articles.
type ArticlesForm struct {
	Articles []ArticleForm `form:"json" binding:"required"`
}
