package validation

import (
	"fmt"
	"inventaris/model"
	"strings"
)

func ValidateItem(item *model.Item) error {

	if strings.TrimSpace(item.Name) == "" {
		return fmt.Errorf("item name cannot be empty")
	}

	if strings.TrimSpace(item.Photo) == "" {
		return fmt.Errorf("item photo cannot be empty")
	}

	if item.Price <= 0 {
		return fmt.Errorf("item price must be greater than zero")
	}


	if strings.TrimSpace(item.PurchaseDate) == "" {
		return fmt.Errorf("purchase date cannot be empty")
	}

	if item.CategoryID <= 0 {
		return fmt.Errorf("category ID cannot empty or must be greater than zero")
	}

	if item.UsageDays < 0 {
		return fmt.Errorf("usage days cannot be negative")
	}

	return nil
}
