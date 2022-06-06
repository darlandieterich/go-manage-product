package repo

import (
	"context"
	"product_manager/domain/model"
	"product_manager/domain/repository"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type StockRepository interface {
	repository.StockRepository
}

type StockConn struct {
	conn *gorm.DB
}

func NewStockRepository(conn *gorm.DB) StockRepository {
	return &StockConn{conn: conn}
}

func (t *StockConn) Delete(ctx context.Context, uuid uuid.UUID) error {
	if t.conn.Delete(&model.Stock{ID: uuid}).Error != nil {
		return t.conn.Error
	}

	return nil
}

func (t *StockConn) FindAll(ctx context.Context, productID uuid.UUID) (stocks []*model.Stock, err error) {
	if delete := t.conn.Find(&stocks, "product_id = ?", productID); delete.Error != nil {
		return nil, delete.Error
	}

	return
}

func (t *StockConn) FindByProductID(ctx context.Context, productID uuid.UUID) (stock *model.Stock, err error) {
	if finded := t.conn.First(&stock, "code = ?", productID); finded.Error != nil {
		return stock, finded.Error
	}

	return
}

func (t *StockConn) Create(ctx context.Context, stock *model.Stock) (uuid uuid.UUID, err error) {
	if created := t.conn.Create(&stock); created.Error != nil {
		return uuid, created.Error
	}

	return stock.ID, nil
}

func (t *StockConn) Update(ctx context.Context, stock *model.Stock) error {
	finded := t.conn.Model(&model.Stock{}).Where("id = ?", stock.ID)

	if finded.Error != nil {
		return finded.Error
	}

	updated := finded.Updates(&model.Stock{
		Total:     stock.Total,
		Cute:      stock.Cute,
		Available: stock.Available,
	})

	if updated.Error != nil {
		return updated.Error
	}

	return nil
}
