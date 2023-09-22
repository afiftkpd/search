package usecase

import (
	"context"
	"search/models"
)

type ProductUsecase interface {
	Search(ctx context.Context, req models.SearchRequest) (*[]models.Product, error)
	AutoComplete(ctx context.Context, keyword string) ([]models.AutoComplete, error)
}
