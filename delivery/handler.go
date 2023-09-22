package delivery

import (
	"encoding/json"
	"fmt"
	"net/http"
	"search/models"
	"search/usecase"
	"strconv"
)

type Handler struct {
	ProductUsecase usecase.ProductUsecase
}

func NewHandler(productUsecase usecase.ProductUsecase) Handler {
	return Handler{productUsecase}
}

func (h *Handler) Search(w http.ResponseWriter, r *http.Request) {
	page, err := strconv.Atoi(r.FormValue("page"))
	if err != nil {
		w.Write([]byte(err.Error()))
	}

	limit, err := strconv.Atoi(r.FormValue("limit"))
	if err != nil {
		w.Write([]byte(err.Error()))
	}

	params := models.SearchRequest{
		Keyword:      r.FormValue("keyword"),
		Page:         page,
		Limit:        limit,
		SortingField: r.FormValue("sorting_field"),
		SortingOrder: r.FormValue("sorting_order"),
	}

	products, err := h.ProductUsecase.Search(r.Context(), params)
	if err != nil {
		w.Write([]byte(err.Error()))
	}

	b, err := json.Marshal(products)
	if err != nil {
		w.Write([]byte(err.Error()))
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(fmt.Sprintf("%+v\n", string(b))))
}

func (h *Handler) AutoComplete(w http.ResponseWriter, r *http.Request) {
	keyword := r.FormValue("keyword")
	products, err := h.ProductUsecase.AutoComplete(r.Context(), keyword)
	if err != nil {
		w.Write([]byte(err.Error()))
	}

	b, err := json.Marshal(products)
	if err != nil {
		w.Write([]byte(err.Error()))
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(fmt.Sprintf("%+v\n", string(b))))
}
