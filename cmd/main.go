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
	//h.DeleteProduct("9ed7bd41-5a2d-4dbf-ba3a-4bcc8604ae51")
	//h.GetProduct("f6db138a-a1a0-47b2-94ff-d9972ed8a10d")
	//h.GetAllProduct(2,6)
	//h.UpdateProduct("662a8646-fa67-4800-8e07-75bd8a7863a1","Qo'y go'shti","Chilonzor",100000)

	//h.CreateCategory("Bannan")
	//h.DeleteCategory("837eea59-aff8-4e20-a9c1-9f8a02f972f3")
	//h.GetCategory("d6b9e087-6c76-4355-a461-b8e6041d6c4a")
	//h.GetAllCategory(2,4)

	//task 1
	// h.MaxTranzit()

	// task 2
	//h.TopBranchSum()

	//task 3
	//h.TopProductTransaction()

	// task 4
	//  h.TopCategoryTransaction()

	// task 6
	// h.BranchTransactionCount()

	// task7
	// h.Task7()
	
	// task 8
	h.Task8()

	//Task 9
	//h.Task9()

	//Task 10
	//h.Task10()

	// Task 12
	//h.Task12()

	// Task 11
	//h.Task11()
}
