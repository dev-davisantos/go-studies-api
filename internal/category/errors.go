package category

import "errors"

var (
	ErrCategoryNotFound           = errors.New("category not found")
	ErrInvalidCategoryName        = errors.New("invalid category name")
	ErrInvalidCategoryDescription = errors.New("invalid category description")
)
