package handler

import (
	"fmt"
	"playground/newProject/models"
)

func (c *handler) MaxTranzit() (models.GetAllBranchProductTransactionResponce, error) {
	body, err := c.strg.BranchProductTransaction().GetAllBranchProductTransaction(models.GetAllBranchProductTransactionRequest{})
	if err != nil {
		return models.GetAllBranchProductTransactionResponce{},err
	}
	
	fmt.Println(body)

	
	return  models.GetAllBranchProductTransactionResponce{},nil
}