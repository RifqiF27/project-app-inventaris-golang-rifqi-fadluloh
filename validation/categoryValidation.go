package validation

import (
    "inventaris/model"
    "errors"
)

func ValidateCategory(category *model.Category) error {
    if category.Name == "" {
        return errors.New("category name cannot be empty")
    }
    if category.Description == "" {
        return errors.New("category description cannot be empty")
    }
    return nil
}