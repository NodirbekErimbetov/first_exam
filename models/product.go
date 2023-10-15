package models

type CreateProduct struct {
	Id      string `json:"id"`
	Name    string `json:"name"`
	Price   int `json:"price"`
	CategoryId string `json:"category_id"`
}

type Product struct {
	Id      string `json:"id"`
	Name    string `json:"name"`
	Price   int `json:"price"`
	CategoryId string `json:"category_id"`
}

type IdRequestProduct struct {
	Id string `json:"id"`
}

type GetAllProductRequest struct {
	Page  int
	Limit int
	Name  string
}
type GetAllProduct struct {
	Products []Product
	Count    int
}


