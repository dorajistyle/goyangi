package locationService

// LocationFilter is a filter for retriving locations.
type LocationFilter struct {
	UserId          int   `json:"userId"`
	Categories      []int `json:"categories"`
	CurrentPage     int   `json:"currentPage"`
	LocationPerPage int   `json:"locationPerPage"`
}

// LocationForm is used when creating or updating a location.
type LocationForm struct {
	Id        int64   `form:"id"`
	UserId    int64   `form:"userId"`
	Latitude  float64 `form:"latitude" binding:"required"`
	Longitude float64 `form:"longitude" binding:"required"`
	Type      string  `form:"type"`
	Name      string  `form:"name" binding:"required"`
	Url       string  `form:"url" binding:"required"`
	Content   string  `form:"content"`
	Address   string  `form:"address"`
}
