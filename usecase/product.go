package usecase

import (
	"context"
	"search/models"
	"search/repository/es"
)

type productUsecase struct {
	ElasticRepo es.ProductRepository
}

func NewProductUsecase(esRepo es.ProductRepository) ProductUsecase {
	return &productUsecase{esRepo}
}

func (p *productUsecase) Search(ctx context.Context, req models.SearchRequest) (*[]models.Product, error) {
	return p.ElasticRepo.Search(ctx, req)
}

func (p *productUsecase) AutoComplete(ctx context.Context, keyword string) ([]models.AutoComplete, error) {
	return p.ElasticRepo.Autocomplete(ctx, keyword)
}
