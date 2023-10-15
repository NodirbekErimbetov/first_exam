package models

type Category struct {
	Id      string `json:"id"`
	Name    string `json:"name"`

}
type CreateCategory struct {
	Id      string `json:"id"`
	Name    string `json:"name"`

}
type IdRequestCategory struct {
	Id string `json:"id"`
}
type GetAllCategoryRequest struct {
	Page  int
	Limit int
	Name  string
}
type GetAllCategory struct {
	Categorys []Category
	Count    int
}


