package category

import (
	"database/sql"
	"errors"
	"log"
)

type CategoryRepository struct {
	db *sql.DB
}

func NewCategoryRepository(db *sql.DB) *CategoryRepository {
	return &CategoryRepository{db: db}
}

func (r *CategoryRepository) FindById(id int) (*Category, error) {
	var category Category
	query := "SELECT id, name, description FROM categories_tb WHERE id = $1"

	if err := r.db.QueryRow(query, id).Scan(&category.id, &category.name, &category.description); err != nil {
		return nil, ErrCategoryNotFound
	}

	return &category, nil
}

func (r *CategoryRepository) FindAll() (*[]Category, error) {
	rows, err := r.db.Query("SELECT id, name, description FROM categories_tb")
	if err != nil {
		return nil, ErrNoCategoriesFound
	}
	var categories []Category

	for rows.Next() {
		var c Category
		err := rows.Scan(&c.id, &c.name, &c.description)
		if err != nil {
			return nil, ErrCategoryConversionFailed
		}
		categories = append(categories, c)
	}

	if err := rows.Err(); err != nil {
		return nil, errors.New("failed to fetch all categories")
	}

	return &categories, nil
}

func (r *CategoryRepository) Create(cat *Category) (*Category, error) {
	query := "INSERT INTO categories_tb (name, description) VALUES ($1, $2)"

	result, err := r.db.Exec(query, cat.name, cat.description)
	if err != nil {
		return nil, ErrFailedToCreateCategory
	}

	catId, err := result.LastInsertId()
	if err != nil {
		return nil, ErrCategoryConversionFailed
	}

	cat.SetId(catId)

	return cat, nil
}

// Just description can be updated
func (r *CategoryRepository) updateDescription(id int, description string) (*Category, error) {
	query := "UPDATE categories_tb SET description = $1 WHERE id = $2"

	result, err := r.db.Exec(query, description, id)
	if err != nil {
		return nil, ErrCategoryFailedToUpdate
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return nil, sql.ErrNoRows
	}
	if rowsAffected == 0 {
		return nil, ErrCategoryNotFound
	}

	cat, err := r.FindById(id)
	if err != nil {
		return nil, ErrCategoryNotFound
	}

	return cat, nil
}

// Delete Method exists just for complete the CRUD
func (r *CategoryRepository) Delete(id int) error {
	query := "DELETE FROM categories_tb WHERE id = $1"
	result, err := r.db.Exec(query, id)
	if err != nil {
		return ErrCategoryNotFound
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return errors.New("failed to fetch affected rows")
	}

	log.Printf("%d rows affected", rowsAffected)

	return nil
}
