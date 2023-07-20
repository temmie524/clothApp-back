package repositories

import (
	"clothApp-backend/model"
	"fmt"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type IItemsRepository interface {
	GetAllItems(items *[]model.Item, userId uint) error
	GetItemById(item *model.Item, UserId uint, ItemId uint) error
	CreateItem(item *model.Item) error
	UpdateItem(item *model.Item) error
	DeleteItem(userId uint, itemId uint) error
}

type itemRepository struct {
	db *gorm.DB
}

func NewItemRepository(db *gorm.DB) IItemsRepository {
	return &itemRepository{db}
}

func (ir *itemRepository) GetAllItems(items *[]model.Item, userId uint) error {
	if err := ir.db.Joins("User").Where("user_id = ?", userId).Order("created_at").Find(items).Error; err != nil {
		return err
	}
	return nil
}

func (ir *itemRepository) getItemsById(item *model.Item, userId uint, itemId uint) error {
	if err := ir.db.Joins("User").Where("user_id = ?", userId).First(item, itemId).Error; err != nil {
		return err
	}
	return nil
}

func (ir *itemRepository) CreateItem(item *model.Item) error {
	if err := ir.db.Create(item).Error; err != nil {
		return err
	}
	return nil
}

func (ir *itemRepository) UpdateItem(item *model.Item, userId uint, itemId uint) error {
	result := ir.db.Model(item).Clauses(clause.Returning{}).Where("id = ? AND user_id = ?", itemId, userId).Update("name", item.Name).Update("price", item.Price).Update("category_id", item.CategoryId).Update("brand_id", item.BrandId).Update("season_id", item.SeasonId).Update("buy_date", item.BuyDate).Update("picture", item.Picture)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected < 1 {
		return fmt.Errorf("no record found")
	}
	return nil
}

func (ir *itemRepository) DeleteItem(userId uint, itemId uint) error {
	result := ir.db.Where("id = ? AND user_id = ?", itemId, userId).Delete(&model.Item{})
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected < 1 {
		return fmt.Errorf("no record found")
	}
	return nil
}
