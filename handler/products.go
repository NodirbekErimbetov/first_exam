package handler

import (
	"fmt"
	"playground/newProject/models"
)

func (h *handler) CreateProduct(id, name, category string, price int )  { 
	resp, err := h.strg.Product().CreateProduct(models.CreateProduct{
	Id      :id,
    Name    :name,
    Price   :price, 
    CategoryId :category,
	})
	if err != nil {
		fmt.Println("error from CreateProduct:", err.Error())
		return 
	}
	fmt.Println("created new branch with id:", resp)
	return 
}
func (h *handler) UpdateProduct(id , name, category string, price int) {
	resp, err := h.strg.Product().UpdateProduct(models.Product{
		Id      :id,
		Name    :name,
		Price   :price, 
		CategoryId :category,
	})
	if err != nil {
		fmt.Println("error from UpdateProduct:", err.Error())
		return
	}
	fmt.Println(resp)
}

func (h *handler) GetProduct(id string) {
	resp, err := h.strg.Product().GetProduct(models.IdRequestProduct{Id: id})
	if err != nil {
		fmt.Println("error from GetProduct:", err.Error())
		return
	}
	fmt.Println(resp)
}

func (h *handler) GetAllProduct(page, limit int) {
	if page < 1 {
		page = h.cfg.Page
	}
	if limit < 1 {
		limit = h.cfg.Limit
	}
	resp, err := h.strg.Product().GetAllProduct(models.GetAllProductRequest{
		Page:  page,
		Limit: limit,
	})
	if err != nil {
		fmt.Println("error from GetAllProduct:", err.Error())
		return
	}
	fmt.Println(resp)
}
func (h *handler) DeleteProduct(id string) {
	resp, err := h.strg.Product().DeleteProduct(models.IdRequestProduct{Id: id})
	if err != nil {
		fmt.Println("error from DeleteProduct:", err.Error())
		return
	}
	fmt.Println(resp)
}




