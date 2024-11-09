package repository

import (
	"database/sql"
	_ "github.com/lib/pq"
	"inventaris/model"
)

type CategoryRepository interface {
	GetAll() ([]model.Category, error)
	Create(category *model.Category) error
	GetByID(id int) (model.Category, error)
	Update(category *model.Category) error
	Delete(id int) error
}

type categoryRepository struct {
	db *sql.DB
}

func NewCategoryRepository(db *sql.DB) CategoryRepository {
	return &categoryRepository{db: db}
}

func (r *categoryRepository) GetAll() ([]model.Category, error) {
	rows, err := r.db.Query(`SELECT id, name, description FROM "Categories"`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var categories []model.Category
	for rows.Next() {
		var category model.Category
		if err := rows.Scan(&category.ID, &category.Name, &category.Description); err != nil {
			return nil, err
		}
		categories = append(categories, category)
	}

	return categories, nil
}

func (r *categoryRepository) Create(category *model.Category) error {
	query := `INSERT INTO "Categories"(name, description) VALUES($1, $2) RETURNING id`
	err := r.db.QueryRow(query, category.Name, category.Description).Scan(&category.ID)
	return err
}

func (r *categoryRepository) GetByID(id int) (model.Category, error) {
	var category model.Category
	query := `SELECT id, name, description FROM "Categories" WHERE id = $1`
	err := r.db.QueryRow(query, id).Scan(&category.ID, &category.Name, &category.Description)
	if err != nil {
		return model.Category{}, err
	}
	return category, nil
}

func (r *categoryRepository) Update(category *model.Category) error {
	query := `UPDATE "Categories" SET name=$1, description=$2 WHERE id=$3`
	_, err := r.db.Exec(query, category.Name, category.Description, category.ID)

	return err
}

func (r *categoryRepository) Delete(id int) error {
	query := `DELETE FROM "Categories" WHERE id = $1`
	_, err := r.db.Exec(query, id)
	return err
}
