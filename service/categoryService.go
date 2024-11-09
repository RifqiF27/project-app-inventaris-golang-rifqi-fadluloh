package service

import (
    "inventaris/model"
    "inventaris/repository"
)

type CategoryService struct {
    Repo repository.CategoryRepository
}

func NewCategoryService(repo repository.CategoryRepository) *CategoryService {
    return &CategoryService{Repo: repo}
}

func (s *CategoryService) GetAllCategories() ([]model.Category, error) {
    return s.Repo.GetAll()
}

func (s *CategoryService) CreateCategory(category *model.Category) error {
    return s.Repo.Create(category)
}

func (s *CategoryService) GetCategoryByID(id int) (model.Category, error) {
    return s.Repo.GetByID(id)
}

func (s *CategoryService) UpdateCategory(category *model.Category) error {
    return s.Repo.Update(category)
}

func (s *CategoryService) DeleteCategory(id int) error {
    return s.Repo.Delete(id)
}
