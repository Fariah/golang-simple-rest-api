package pkg

type User struct {
	Id       int    `json:"-"`
	Name     string `json:"name"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type Social struct {
	Id          int    `json:"-"`
	UserId      int    `json:"userId"`
	SocialTitle string `json:"socialTitle"`
	SocialId    string `json:"socialId"`
}
