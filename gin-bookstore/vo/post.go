package vo

type CreaterPostRequest struct {

	CategoryId uint `json:"Category_id" binding:"required"`
	Title string `json:"title" binding:"required,max=20`
	HeadImg string `json:"head_img"`
	Content string `json:"content" binding:"required"`
}

