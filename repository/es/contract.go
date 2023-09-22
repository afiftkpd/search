package es

import (
	"context"
	"search/models"
)

type ProductRepository interface {
	Search(ctx context.Context, req models.SearchRequest) (*[]models.Product, error)
	Autocomplete(ctx context.Context, keyword string) ([]models.AutoComplete, error)
}
