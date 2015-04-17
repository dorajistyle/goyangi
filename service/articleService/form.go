package articleService

// ArticleFilter is a filter for retriving articles.
type ArticleFilter struct {
	UserId         int   `json:"userId"`
	Categories     []int `json:"categories"`
	CurrentPage    int   `json:"currentPage"`
	ArticlePerPage int   `json:"articlePerPage"`
}

// ArticleForm is a form of article.
type ArticleForm struct {
	Id            int64  `form:"id"`
	UserId        int64  `form:"userId"`
	CategoryId    int    `form:"categoryId"`
	Title         string `form:"title"`
	Url           string `form:"url" binding:"required"`
	ImageName     string `form:"imageName" binding:"required"`
	ThumbnailName string `form:"thumbnailName"`
	Content       string `form:"content"`
}

// ArticlesForm is used when creating or updating multiple Articles.
type ArticlesForm struct {
	Articles []ArticleForm `form:"json" binding:"required"`
}
