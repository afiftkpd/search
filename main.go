package main

import (
	"fmt"
	"net/http"
	"search/delivery"
	"search/repository/es"
	"search/usecase"

	"github.com/elastic/go-elasticsearch/v8"
)

func main() {
	fmt.Println("running")
	elasticConn, err := elasticsearch.NewTypedClient(elasticsearch.Config{})
	if err != nil {
		panic(err)
	}

	elastic := es.NewProductRepository(elasticConn)
	uc := usecase.NewProductUsecase(elastic)
	h := delivery.NewHandler(uc)
	http.HandleFunc("/search", h.Search)
	http.HandleFunc("/autocomplete", h.AutoComplete)

	err = http.ListenAndServe(":8081", nil)
	panic(err)
}
