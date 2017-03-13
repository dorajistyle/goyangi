package oauthService

// FacebookUser is a struct that contained facebook user information.
type FacebookUser struct {
	ID         string `json:"id"`
	Username   string `json:"name"`
	ProfileUrl string `json:"link"`
}

// GithubUser is a struct that contained github user information.
type GithubUser struct {
	UserID     int    `json:"id"`
	Email      string `json:"email"`
	Username   string `json:"login"`
	Name       string `json:"name""`
	ImageUrl   string `json:"avatar_url"`
	ProfileUrl string `json:"html_url"`
}

// Email is a struct that contained googleUser's email.
type Email struct {
	Value string `json:"value" binding:"required"`
	Type  string `json:"type" binding:"required"`
}

// Image is a struct that contained googleUser's image meta data.
type Image struct {
	URL       string `json:"url" binding:"required"`
	IsDefault string `json:"isDefault" binding:"required"`
}

// GoogleUser is a struct that contained google user information.
type GoogleUser struct {
	ID         string  `json:"id"`
	Username   string  `json:"nickname"`
	Emails     []Email `json:"emails"`
	Name       string  `json:"displayName"`
	Image      Image   `json:"image"`
	ProfileUrl string  `json:"url"`
}

// KakaoUser is a struct that contained kakao user information.
type KakaoUser struct {
	ID         string `json:"id"`
	Email      string `json:"email"`
	Username   string `json:"login"`
	Name       string `json:"name""`
	ImageUrl   string `json:"avatar_url"`
	ProfileUrl string `json:"html_url"`
}

// NaverUser is a struct that contained naver user information.
type NaverUser struct {
	ID         string `json:"id"`
	Email      string `json:"email"`
	Username   string `json:"login"`
	Name       string `json:"name""`
	ImageUrl   string `json:"avatar_url"`
	ProfileUrl string `json:"html_url"`
}

// TwitterUser is a struct that contained twitter user information.
type TwitterUser struct {
	ID         string `json:"id"`
	Email      string `json:"email"`
	Name       string `json:"name"`
	Username   string `json:"screen_name"`
	ProfileUrl string `json:"url"`
	ImageUrl   string `json:"profile_image_url"`
}

// YahooUser is a struct that contained yahoo user information.
type YahooUser struct {
	ID         string `json:"guid"`
	Email      string `json:"email"`
	Name       string `json:"nickname"`
	Username   string `json:"nickname"`
	ImageUrl   string `json:"imageURL"`
	ProfileUrl string `json:"profileURL"`
}