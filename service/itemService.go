package service

import (
    "inventaris/model"
    "inventaris/repository"
)

type ItemService struct {
    Repo repository.ItemRepository
}

func NewItemService(repo repository.ItemRepository) *ItemService {
    return &ItemService{Repo: repo}
}

func (s *ItemService) GetAllItemsReplacementNeeded() ([]model.Item, error) {
    return s.Repo.GetAllReplacementNeeded()
}
func (s *ItemService) GetAllItems(name, categoryName string, maxUsageDays bool, limit, page int) ([]model.Item, int, int, error) {
    return s.Repo.GetAll(name, categoryName, maxUsageDays, limit, page)
}

func (s *ItemService) CreateItem(category *model.Item) error {
    return s.Repo.Create(category)
}

func (s *ItemService) GetItemByID(id int) (model.Item, error) {
    return s.Repo.GetByID(id)
}

func (s *ItemService) UpdateItem(category *model.Item) error {
    return s.Repo.Update(category)
}

func (s *ItemService) DeleteItem(id int) error {
    return s.Repo.Delete(id)
}


