package pkg

type Box struct {
	Id    int    `json:"-"`
	Title string `json:"title"`
	Image string `json:"image"`
}
