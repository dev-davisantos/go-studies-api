package category

import "errors"

var (
	ErrCategoryNotFound           = errors.New("category not found")
	ErrInvalidCategoryName        = errors.New("invalid category name")
	ErrInvalidCategoryDescription = errors.New("invalid category description")
	ErrNoCategoriesFound          = errors.New("no categories found")
	ErrCategoryConversionFailed   = errors.New("category conversion failed")
	ErrFailedToCreateCategory     = errors.New("failed to create category")
	ErrCategoryFailedToUpdate     = errors.New("failed to update category")
)
