package models

type Product struct {
	ID          int64  `json:"id"`
	Description string `json:"description"`
	Name        string `json:"name"`
	Price       int64  `json:"price"`
	Rating      int    `json:"rating"`
	ImageURL    string `json:"image_url"`
	Stock       int    `json:"stock"`
}

type AutoComplete struct {
	Name string `json:"name"`
}

type SearchRequest struct {
	Keyword      string
	Page         int
	Limit        int
	SortingField string
	SortingOrder string
}
