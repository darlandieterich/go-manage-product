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

func (t *ProductConn) ListAll(ctx context.Context) (products []*model.Product, err error) {
	if finded := t.conn.Find(&products); finded.Error != nil {
		return nil, finded.Error
	}

	return
}

func (t *ProductConn) FindByCode(ctx context.Context, code string) (product *model.Product, err error) {
	if finded := t.conn.Preload("Stock").First(&product, "code = ?", code); finded.Error != nil {
		return product, finded.Error
	}

	return
}

func (t *ProductConn) FindById(ctx context.Context, uuid uuid.UUID) (product *model.Product, err error) {
	if finded := t.conn.Preload("Stock").First(&product, "id = ?", uuid); finded.Error != nil {
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
	updated := t.conn.Session(&gorm.Session{FullSaveAssociations: true}).Save(&product)
	if updated.Error != nil {
		return updated.Error
	}

	return nil
}
