package repositories

import (
	"errors"

	"gin-fleamarket/models"
)

type IItemRepository interface {
	FindAll() (*[]models.Item, error)
	FindById(itemid uint) (*models.Item, error)
}

type ItemMemoryRepository struct {
	items []models.Item
}

func NewItemMemoryRepository(items []models.Item) IItemRepository {
	return &ItemMemoryRepository{items: items}
}

func (r *ItemMemoryRepository) FindAll() (*[]models.Item, error) {
	return &r.items, nil
}

func (r *ItemMemoryRepository) FindById(itemid uint) (*models.Item, error) {
	for _, item := range r.items {
		if item.ID == itemid {
			return &item, nil
		}
	}
	return nil, errors.New("Item not found")
}
