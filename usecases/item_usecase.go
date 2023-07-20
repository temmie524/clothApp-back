package usecases

import (
	"clothApp-backend/model"
	"clothApp-backend/repositories"
)

type IItemUsecase interface {
	GetAllItems(userId uint) ([]model.ItemResponse, error)
	GetItemById(userId uint, itemId uint) (model.ItemResponse, error)
	CreateItem(item model.Item) (model.ItemResponse, error)
	UpdateItem(item model.Item, userId uint, itemId uint) (model.ItemResponse, error)
	DeleteItem(userId uint, itemId uint) error
}

type itemUsecase struct {
	ir repositories.IItemsRepository
}

func NewItemUsecase(ir repositories.IItemsRepository) IItemUsecase {
	return &itemUsecase{ir}
}

func (iu *itemUsecase) GetAllItems(userId uint) ([]model.ItemResponse, error) {
	items := []model.ItemResponse{}
	if err := iu.ir.GetAllItems(&items, userId); err != nil {
		return nil, err
	}
	resItems := []model.ItemResponse{}
	for _, item := range items {
		t := model.ItemResponse{
			Id:         item.Id,
			Name:       item.Name,
			Price:      item.Price,
			CategoryId: item.CategoryId,
			BrandId:    item.BrandId,
			SeasonId:   item.SeasonId,
			BuyDate:    item.BuyDate,
			Picture:    item.Picture,
			CreatedAt:  item.CreatedAt,
			UpdatedAt:  item.UpdatedAt,
		}
		resItems = append(resItems, t)
	}
	return items, nil
}

func (iu *itemUsecase) GetItemById(userId uint, itemId uint) (model.ItemResponse, error) {
	item := model.Item{}
	if err := iu.ir.GetItemById(&item, userId, itemId); err != nil {
		return model.ItemResponse{}, err
	}
	resItem := model.ItemResponse{
		Id:         item.Id,
		Name:       item.Name,
		Price:      item.Price,
		CategoryId: item.CategoryId,
		BrandId:    item.BrandId,
		SeasonId:   item.SeasonId,
		BuyDate:    item.BuyDate,
		Picture:    item.Picture,
		CreatedAt:  item.CreatedAt,
		UpdatedAt:  item.UpdatedAt,
	}
	return resItem, nil
}

func (iu *itemUsecase) CreateItem(item model.Item) (model.ItemResponse, error) {
	if err := iu.ir.CreateItem(&item); err != nil {
		return model.ItemResponse{}, err
	}
	resItem := model.ItemResponse{
		Id:         item.ID,
		Name:       item.Name,
		Price:      item.Price,
		CategoryId: item.CategoryId,
		BrandId:    item.BrandId,
		SeasonId:   item.SeasonId,
		BuyDate:    item.BuyDate,
		Picture:    item.Picture,
		CreatedAt:  item.CreatedAt,
		UpdatedAt:  item.UpdatedAt,
	}
	return resItem, nil
}

func (iu *itemUsecase) UpdateItem(item model.Item, userId uint, itemId uint) (model.ItemResponse, error) {
	if err := iu.ir.UpdateItem(&item, userId, itemId); err != nil {
		return model.ItemResponse{}, err
	}
	resItem := model.ItemResponse{
		Id:         item.ID,
		Name:       item.Name,
		Price:      item.Price,
		CategoryId: item.CategoryId,
		BrandId:    item.BrandId,
		SeasonId:   item.SeasonId,
		BuyDate:    item.BuyDate,
		Picture:    item.Picture,
		CreatedAt:  item.CreatedAt,
		UpdatedAt:  item.UpdatedAt,
	}
	return resItem, nil
}

func (iu *itemUsecase) DeleteItem(userId uint, itemId uint) error {
	if err := iu.ir.DeleteItem(userId, itemId); err != nil {
		return err
	}
	return nil
}
