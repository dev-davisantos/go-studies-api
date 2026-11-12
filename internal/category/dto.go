package category

type CategoryRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type CategoryResponse struct {
	Id          int64  `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}
