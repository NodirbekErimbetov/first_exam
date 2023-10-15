package memory

import (
	"playground/newProject/models"
	"playground/newProject/storage"
)

type store struct {
	branches                  *branchRepo
	users                     *userRepo
	branchProducts            *branchProductRepo
	branchProductTransactions *branchProductTransactionRepo
	products 				  *productRepo
	category				  *categoryRepo
}

func NewStorage(files models.FileNames) storage.StorageI {
	return &store{
		branches:                  NewBranchRepo(files.BranchFile),
		users:                     NewUserRepo(files.UserFile),
		branchProducts:            NewBranchProductRepo(files.BranchProductFile),
		branchProductTransactions: NewBranchProductTransactionRepo(files.BranchProductTransactionFile),
		products: 				   NewProductRepo(files.ProductFile),
		category:                  NewCategoryRepo(files.CategoryFile),
	}
}

func (s *store) Branch() storage.BranchesI {
	return s.branches
}
func (s *store) BranchProduct() storage.BranchesProductsI {
	return s.branchProducts
}
func (s *store) BranchProductTransaction() storage.BranchProductTransactionI {
	return s.branchProductTransactions
}
func (s *store) User() storage.UsersI {
	return s.users
}
func (s *store) Product() storage.ProductI {
	return s.products
}
func (s *store) Category() storage.CategoryI {
	return s.category
}