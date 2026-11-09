package category

type CategoryRequest struct {
	name        string
	description string
}

type CategoryResponse struct {
	id          string
	name        string
	description string
}
