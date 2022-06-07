package repository

import (
	"context"
	"product_manager/domain/model"

	"github.com/google/uuid"
)

type StockRepository interface {
	Create(ctx context.Context, stock *model.Stock) (uuid.UUID, error)
	Update(ctx context.Context, stock *model.Stock) error
	Delete(ctx context.Context, uuid uuid.UUID) error
}
