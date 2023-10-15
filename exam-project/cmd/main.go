package main

import (
	"playground/newProject/config"
	"playground/newProject/handler"
	"playground/newProject/models"
	"playground/newProject/storage/memory"
)

func main() {

	cfg := config.Load()
	strg := memory.NewStorage(models.FileNames{
		BranchFile:                   "data/branches.json",
		UserFile:                     "data/users.json",
		CategoryFile:                 "data/categories.json",
		ProductFile:                  "data/products.json",
		BranchProductFile:            "data/branch_products.json",
		BranchProductTransactionFile: "data/branch_pr_transaction.json",
	})
	h := handler.NewHandler(strg, *cfg)
	//h.GetAllBranchProductTransaction(1, 10, "", "")
	

	//h.CreateProduct("38994-fjnf-2973y","adkjnvlkjds","bbbbbbb",1000)
	//h.DeleteProduct("a2999206-e1aa-4c58-acaa-2b5b7f139064")
	//h.GetProduct("f6db138a-a1a0-47b2-94ff-d9972ed8a10d")
	//h.GetAllProduct(2,6," ")
	//h.UpdateProduct("662a8646-fa67-4800-8e07-75bd8a7863a1","Qo'y go'shti","Chilonzor",100000)
	
	//h.CreateCategory("Bannan")
	//h.DeleteCategory("3c3612f3-c0c4-4739-a12d-3c74e43df106")
	//h.GetCategory("d6b9e087-6c76-4355-a461-b8e6041d6c4a")
	h.GetAllCategory(2,4," ")

	//h.MaxTranzit()
	



}
