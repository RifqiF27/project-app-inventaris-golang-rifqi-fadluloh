package repository

import (
	"database/sql"
	"fmt"
	"inventaris/model"
)

type ItemRepository interface {
	GetAll(name, categoryName string, maxUsageDays bool, limit, page int) ([]model.Item, int, int, error)
	Create(item *model.Item) error
	GetByID(id int) (model.Item, error)
	Update(item *model.Item) error
	Delete(id int) error
	GetAllInvestment() ([]model.Item, error)
	GetAllReplacementNeeded() ([]model.Item, error)
}

type itemRepository struct {
	db *sql.DB
}

func NewItemRepository(db *sql.DB) ItemRepository {
	return &itemRepository{db: db}
}

func (r *itemRepository) GetAllInvestment() ([]model.Item, error) {
	query := `SELECT id, name, photo_url, price, purchase_date, usage_days, depreciation_rate, category_id FROM "Items"`
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var items []model.Item
	for rows.Next() {
		var item model.Item
		if err := rows.Scan(&item.ID, &item.Name, &item.Photo, &item.Price, &item.PurchaseDate, &item.UsageDays, &item.DepreciationRate, &item.CategoryID); err != nil {
			return nil, err
		}
		items = append(items, item)
	}

	return items, nil
}
func (r *itemRepository) GetAllReplacementNeeded() ([]model.Item, error) {
	query := `SELECT i.id, i.name,  c.name AS category, i.purchase_date, i.usage_days
              FROM "Items" i
              JOIN "Categories" c ON i.category_id = c.id
              where i.usage_days > 100`
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var items []model.Item
	for rows.Next() {
		var item model.Item
		if err := rows.Scan(&item.ID, &item.Name, &item.CategoryName, &item.PurchaseDate,
			&item.UsageDays); err != nil {
			return nil, err
		}
		items = append(items, item)
	}

	return items, nil
}

func (r *itemRepository) GetAll(name, categoryName string, maxUsageDays bool, limit, page int) ([]model.Item, int, int, error) {

	query := `SELECT i.id, i.name,  c.name AS category, i.photo_url, i.price, i.purchase_date, i.usage_days
              FROM "Items" i
              JOIN "Categories" c ON i.category_id = c.id
              WHERE 1=1`

	countQuery := `SELECT COUNT(*) FROM "Items" i JOIN "Categories" c ON i.category_id = c.id WHERE 1=1`

	var params []interface{}
	paramIndex := 1

	if name != "" {
		query += ` AND i.name ILIKE $` + fmt.Sprint(paramIndex)
		countQuery += ` AND i.name ILIKE $` + fmt.Sprint(paramIndex)
		params = append(params, name+"%")
		paramIndex++
	}
	if categoryName != "" {
		query += ` AND c.name ILIKE $` + fmt.Sprint(paramIndex)
		countQuery += ` AND c.name ILIKE $` + fmt.Sprint(paramIndex)
		params = append(params, categoryName+"%")
		paramIndex++
	}

	if maxUsageDays {
		query += ` AND i.usage_days > 100`
		countQuery += ` AND i.usage_days > 100`
	}

	var totalItems int
	err := r.db.QueryRow(countQuery, params...).Scan(&totalItems)
	if err != nil {
		return nil, 0, 0, err
	}

	totalPages := (totalItems + limit - 1) / limit

	offset := (page - 1) * limit
	query += ` ORDER BY i.id LIMIT $` + fmt.Sprint(paramIndex) + ` OFFSET $` + fmt.Sprint(paramIndex+1)
	params = append(params, limit, offset)

	rows, err := r.db.Query(query, params...)
	if err != nil {
		return nil, 0, 0, err
	}
	defer rows.Close()

	var items []model.Item
	for rows.Next() {
		var item model.Item
		if err := rows.Scan(&item.ID, &item.Name, &item.CategoryName, &item.Photo, &item.Price, &item.PurchaseDate,
			&item.UsageDays); err != nil {
			return nil, 0, 0, err
		}
		items = append(items, item)
	}

	return items, totalItems, totalPages, nil
}

func (r *itemRepository) Create(item *model.Item) error {

	query := `
		WITH new_item AS (
			INSERT INTO "Items" (name, photo_url, price, purchase_date, usage_days, depreciation_rate, category_id)
			VALUES ($1, $2, $3, $4, $5, $6, $7)
			RETURNING id, category_id
		)
		SELECT i.id, c.name
		FROM new_item i
		JOIN "Categories" c ON i.category_id = c.id
	`
	var categoryName string
	err := r.db.QueryRow(query,
		item.Name, item.Photo, item.Price, item.PurchaseDate, item.UsageDays, item.DepreciationRate, item.CategoryID).
		Scan(&item.ID, &categoryName)
	if err != nil {
		return err
	}

	item.CategoryName = categoryName
	return nil
}

func (r *itemRepository) GetByID(id int) (model.Item, error) {
	var item model.Item
	query := `SELECT i.id, i.name, c.name AS category, i.photo_url, i.price, i.purchase_date, i.usage_days FROM "Items" i JOIN "Categories" c ON i.category_id = c.id WHERE i.id = $1`
	err := r.db.QueryRow(query, id).
		Scan(&item.ID, &item.Name, &item.CategoryName, &item.Photo, &item.Price, &item.PurchaseDate, &item.UsageDays)
	return item, err
}

func (r *itemRepository) Update(item *model.Item) error {
	query := `
		WITH update_item AS (
			UPDATE "Items" 
			SET name=$1, photo_url=$2, price=$3, purchase_date=$4, usage_days=$5, depreciation_rate=$6, category_id=$7 
			WHERE id=$8 
			RETURNING id, category_id
		) 
		SELECT i.id, c.name
		FROM update_item i
		JOIN "Categories" c ON i.category_id = c.id`

	var categoryName string

	err := r.db.QueryRow(query,
		item.Name, item.Photo, item.Price, item.PurchaseDate, item.UsageDays, item.DepreciationRate, item.CategoryID, item.ID).
		Scan(&item.ID, &categoryName)

	if err != nil {
		return err
	}

	item.CategoryName = categoryName
	return nil
}

func (r *itemRepository) Delete(id int) error {
	query := `DELETE FROM "Items" WHERE id = $1`
	_, err := r.db.Exec(query, id)
	return err
}
