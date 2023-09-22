package es

import (
	"context"
	"encoding/json"
	"fmt"
	"search/models"
	"strconv"
	"strings"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/esapi"
	"github.com/elastic/go-elasticsearch/v8/typedapi/core/search"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
)

type productRepository struct {
	DB *elasticsearch.TypedClient
}

func NewProductRepository(db *elasticsearch.TypedClient) ProductRepository {
	return &productRepository{db}
}

func (p *productRepository) Search(ctx context.Context, req models.SearchRequest) (*[]models.Product, error) {
	fmt.Println("keyword: ", req.Keyword)
	fmt.Println(req.Limit, (req.Page-1)*req.Limit)
	var searchQuery esapi.SearchRequest
	result := []models.Product{}
	// searchQuery := p.DB
	// newKeyword := fmt.Sprintf("*%v*", req.Keyword)
	fmt.Println(req.Keyword)
	if len(req.Keyword) > 0 {
		// searchQuery = searchQuery.Query(&types.Query{MatchAll: &types.MatchAllQuery{}})
		searchQuery = esapi.SearchRequest{
			Index: []string{"products"},
			Body: strings.NewReader(fmt.Sprintf(`
			{
				"query": {
					"match": {
						"name": "%s"
					}
				},
				"size": %v,
				"from": %v
			}`, req.Keyword, req.Limit, (req.Page-1)*req.Limit)),
		}
	} else {
		searchQuery = esapi.SearchRequest{
			Index: []string{"products"},
			Body: strings.NewReader(fmt.Sprintf(`
			{
				"query": {
					"match_all": {}
				},
				"size": %v,
				"from": %v
			}`, req.Limit, (req.Page-1)*req.Limit)),
		}
	}

	res, err := searchQuery.Do(ctx, p.DB)
	if err != nil {
		return &result, err
	}

	// Parse and process the hits (search results)
	var response map[string]interface{}
	if err := json.NewDecoder(res.Body).Decode(&response); err != nil {
		fmt.Println("Error parsing search response: " + err.Error())
	}

	hits, found := response["hits"].(map[string]interface{})["hits"].([]interface{})
	if !found {
		fmt.Println("No hits found in the response.")
	}

	for _, hit := range hits {
		source := hit.(map[string]interface{})["_source"]
		fmt.Printf("Document ID: %s\n", hit.(map[string]interface{})["_id"])
		fmt.Printf("Document Source: %v\n", source)

		res := source.(map[string]interface{})
		rawID := hit.(map[string]interface{})["_id"].(string)
		convID, err := strconv.Atoi(rawID)
		if err != nil {
			return &result, err
		}

		result = append(result, models.Product{
			ID:          int64(convID),
			Description: res["description"].(string),
			Name:        res["name"].(string),
			Price:       int64(res["price"].(float64)),
			ImageURL:    res["image_url"].(string),
			Rating:      int(res["rating"].(float64)),
			Stock:       int(res["stock"].(float64)),
		})
	}

	return &result, nil
}

func (p *productRepository) Autocomplete(ctx context.Context, keyword string) ([]models.AutoComplete, error) {
	fmt.Println("keyword: ", keyword)
	result := []models.AutoComplete{}
	searchQuery := p.DB.Search().Index("products")
	if len(keyword) > 0 {
		searchQuery = searchQuery.Request(&search.Request{
			Query: &types.Query{
				Prefix: map[string]types.PrefixQuery{
					"name": {Value: keyword},
				},
			},
			Fields: []types.FieldAndFormat{
				{
					Field: "name",
				},
			},
		})
	}

	res, err := searchQuery.Do(ctx)
	if err != nil {
		return result, err
	}

	for _, h := range res.Hits.Hits {
		product := models.AutoComplete{}

		x, err := h.Source_.MarshalJSON()
		if err != nil {
			fmt.Println(err.Error())
		}

		err = json.Unmarshal(x, &product)
		if err != nil {
			fmt.Println(err.Error())
		}

		result = append(result, product)
	}

	return result, nil
}
