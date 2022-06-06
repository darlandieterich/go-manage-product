package repo

import (
	"context"
	"product_manager/domain/model"
	"product_manager/domain/repository"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ProductRepository interface {
	repository.ProductRepository
}

type ProductConn struct {
	conn *gorm.DB
}

func NewProductRepository(conn *gorm.DB) ProductRepository {
	return &ProductConn{conn: conn}
}

func (t *ProductConn) Delete(ctx context.Context, uuid uuid.UUID) error {
	if t.conn.Delete(&model.Product{ID: uuid}).Error != nil {
		return t.conn.Error
	}

	return nil
}

func (t *ProductConn) FindAll(ctx context.Context) (products []*model.Product, err error) {
	if finded := t.conn.Find(&products); finded.Error != nil {
		return nil, finded.Error
	}

	return
}

func (t *ProductConn) FindByCode(ctx context.Context, code string) (product *model.Product, err error) {
	if finded := t.conn.First(&product, "code = ?", code); finded.Error != nil {
		return product, finded.Error
	}

	return
}

func (t *ProductConn) Create(ctx context.Context, product *model.Product) (uuid uuid.UUID, err error) {
	if created := t.conn.Create(&product); created.Error != nil {
		return uuid, created.Error
	}

	return product.ID, nil
}

func (t *ProductConn) Update(ctx context.Context, product *model.Product) error {
	if done := t.conn.Model(&model.Product{}).Where("id = ?", product.ID).Update("done", true); done.Error != nil {
		return done.Error
	}

	return nil
}
