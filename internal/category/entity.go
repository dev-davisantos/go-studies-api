package category

import "errors"

type Category struct {
	id          int64  `db:"id" json:"id"`
	name        string `db:"name" json:"name"`
	description string `db:"description" json:"description"`
}

func NewCategory(name string, description string) (*Category, error) {
	if name == "" {
		return nil, errors.New("name is required")
	}
	if description == "" {
		return nil, errors.New("description is required")
	}

	return &Category{
		id:          0,
		name:        name,
		description: description,
	}, nil
}

func (c *Category) ID() int64 {
	return c.id
}

func (c *Category) Name() string {
	return c.name
}

func (c *Category) Description() string {
	return c.description
}

func (c *Category) SetId(id int64) error {
	if id <= 0 {
		return errors.New("Id must be grater than 0")
	}
	c.id = id
	return nil
}

func (c *Category) SetName(name string) error {
	if name == "" {
		return errors.New("name is required")
	}
	c.name = name
	return nil
}

func (c *Category) SetDescription(description string) error {
	if description == "" {
		return errors.New("Description is required")
	}
	c.description = description
	return nil
}
