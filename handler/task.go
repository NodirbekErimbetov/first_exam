package handler

import (
	"fmt"
	"log"
	"playground/newProject/models"
	"sort"
)

// task1

func (c *handler) MaxTranzit() {

	resp, err := c.strg.BranchProductTransaction().GetListBranchProductTransaction()
	if err != nil {
		log.Println("err task1", err)
		return
	}
	res1, err := c.strg.Branch().GetListBranches()
	if err != nil {
		log.Println("err task1", err)
		return
	}
	fmt.Println(res1)
	var resp_data_count = make(map[string]int)

	for _, value := range resp.BranchProductTransactions {
		resp_data_count[value.BranchID]++
	}

	for k, v := range resp_data_count {
		for _, value := range res1.Branches {
			if value.Id == k {
				fmt.Print(value.Name)
			}
		}
		max := resp_data_count[k]
		if v > max {
			max = v
		}
		fmt.Println(" ", max)
	}
}

// task2

func (h *handler) TopBranchSum() {
	transactions, err := h.strg.BranchProductTransaction().GetListBranchProductTransaction()
	if err != nil {
		log.Println("err task2", err)
		return
	}
	products, err := h.strg.Product().GetListProduct()
	if err != nil {
		log.Println("err task2", err)
		return
	}
	branch, err := h.strg.Branch().GetListBranches()
	if err != nil {
		log.Println(err)
		return
	}
	result := make(map[string]int)
	sum := 0
	result2 := make(map[int]string)
	for _, v := range transactions.BranchProductTransactions {
		// fmt.Println("Value ", v.ProductID)
		for _, val := range products.Products {
			if v.ProductID == val.Id {
				sum = v.Quantity * val.Price
			}
		}
		result[v.BranchID] = sum
		sum = 0
	}
	for k, v := range result {
		for _, value := range branch.Branches {
			if k == value.Id {
				result2[v] = value.Name
			}
		}
	}

	keys := make([]int, 0, len(result2))
	for k := range result2 {
		keys = append(keys, k)

	}
	sort.Sort(sort.Reverse(sort.IntSlice(keys)))
	for _, k := range keys {
		fmt.Println(result2[k], k)
	}
}

// task3

type KeyValue struct {
	Key   string
	Value int
}

func (g *handler) TopProductTransaction() {
	transaction, err := g.strg.BranchProductTransaction().GetListBranchProductTransaction()
	if err != nil {
		log.Println("task3>>>", err.Error())
		return
	}
	var result = map[string]int{}
	for _, v := range transaction.BranchProductTransactions {
		product, err := g.strg.Product().GetProduct(models.IdRequestProduct{
			Id: v.ProductID,
		})
		if err != nil {
			log.Println("task3>>>", err.Error())
			return
		}
		result[product.Name]++
	}
	var mpslice []KeyValue
	for key, value := range result {
		mpslice = append(mpslice, KeyValue{key, value})
	}
	sort.Slice(mpslice, func(i, j int) bool {
		return mpslice[i].Value > mpslice[j].Value
	})

	for _, kv := range mpslice {
		fmt.Println(kv.Key, kv.Value)
	}

}

// task 4
func (g *handler) TopCategoryTransaction() {
	transaction, err := g.strg.BranchProductTransaction().GetListBranchProductTransaction()
	if err != nil {
		log.Println("task4>>>", err.Error())
		return
	}
	var result = map[string]int{}
	for _, v := range transaction.BranchProductTransactions {
		product, err := g.strg.Product().GetProduct(models.IdRequestProduct{
			Id: v.ProductID,
		})
		if err != nil {
			log.Println("task4>>>", err.Error())
			return
		}

		category, err := g.strg.Category().GetCategory(models.IdRequestCategory{
			Id: product.CategoryId,
		})
		if err != nil {
			log.Println("task4>>>", err.Error())
			return
		}
		result[category.Name]++
	}
	var mpslice []KeyValue
	for key, value := range result {
		mpslice = append(mpslice, KeyValue{key, value})
	}
	sort.Slice(mpslice, func(i, j int) bool {
		return mpslice[i].Value > mpslice[j].Value
	})

	for _, kv := range mpslice {
		fmt.Println(kv.Key, kv.Value)
	}

}

// task6
func (h *handler) BranchTransactionCount() {
	transaction, err := h.strg.BranchProductTransaction().GetListBranchProductTransaction()
	if err != nil {
		log.Println("task6 GetListBranchProductTransaction>>>", err.Error())
		return
	}
	sumplus := 0
	summinus := 0
	plusmap := make(map[string]int)
	sumplusmap := make(map[string]int)
	minusmap := make(map[string]int)
	summinusmap := make(map[string]int)
	for _, v := range transaction.BranchProductTransactions {
		product, err := h.strg.Product().GetProduct(models.IdRequestProduct{
			Id: v.ProductID,
		})
		if err != nil {
			log.Println("task6 GetProduct>>>", err.Error())
			return
		}

		branch, err := h.strg.Branch().GetBranch(models.IdRequest{
			Id: v.BranchID,
		})
		if err != nil {
			log.Println("task6 GetBranch>>>", err.Error())
			return
		}
		switch v.Type {
		case "plus":
			plusmap[branch.Name]++
			sumplus = v.Quantity * product.Price
			sumplusmap[branch.Name] = sumplus
		default:
			minusmap[branch.Name]++
			summinus = v.Quantity * product.Price
			summinusmap[branch.Name] = summinus
		}
		sumplus = 0
		summinus = 0

	}
	fmt.Println("          Transaction      Sum")
	fmt.Println("          Plus  Minus", "plussum      minussum")
	for k := range plusmap {

		fmt.Println(k, " ", plusmap[k], " ", minusmap[k], " ", sumplusmap[k], " ", summinusmap[k])
	}
}

// task 7
func (h *handler) Task7() {
	transaction, err := h.strg.BranchProductTransaction().GetListBranchProductTransaction()
	if err != nil {
		log.Println("task7>>>", err.Error())
		return
	}

	mymap := make(map[string]int)
	for _, v := range transaction.BranchProductTransactions {
		if v.Type == "plus" {
			mymap[v.CreatedAt] += v.Quantity
		}

	}

	var mpslice []KeyValue
	for key, value := range mymap {
		mpslice = append(mpslice, KeyValue{key, value})
	}
	sort.Slice(mpslice, func(i, j int) bool {
		return mpslice[i].Value > mpslice[j].Value
	})

	for _, kv := range mpslice {
		fmt.Println(kv.Key, kv.Value)
	}
}

// task 8
func (h *handler) Task8() {
	plus := 0
	minus := 0
	product, err := h.strg.Product().GetListProduct()
	if err != nil {
		log.Println("task get list8>>>", err.Error())
		return
	}
	transaction, err := h.strg.BranchProductTransaction().GetListBranchProductTransaction()
	if err != nil {
		log.Println("task8 get branch>>>", err.Error())
		return
	}

	plusproducts := make(map[string]int)
	minusproducts := make(map[string]int)
	for _, value := range product.Products {
		for _, v := range transaction.BranchProductTransactions {

			if value.Id == v.ProductID {
				switch v.Type {
				case "plus":
					plus += v.Quantity
					plusproducts[value.Name] = plus
				case "minus":
					minus += v.Quantity
					minusproducts[value.Name] = minus

				}
			}
			plus = 0
			minus = 0
		}
	}
	fmt.Println("Nomi Krim     Chqim")
	for _,k := range product.Products {
		fmt.Println(k.Name, " ", plusproducts[k.Name], "  ", minusproducts[k.Name])
	}
}

// Task 9

func (h *handler) Task9() {
	bProduct, err := h.strg.BranchProduct().GetListBranchProduct()
	if err != nil {
		log.Println("task9 GetListBranchProduct>>>", err.Error())
		return
	}
	mymap := make(map[string]int)
	for _, v := range bProduct.BranchProducts {
		product, err := h.strg.Product().GetProduct(models.IdRequestProduct{
			Id: v.ProductID,
		})
		if err != nil {
			log.Println("task9 GetProduct>>>", err.Error())
			return
		}
		branch, err := h.strg.Branch().GetBranch(models.IdRequest{
			Id: v.BranchID,
		})
		if err != nil {
			log.Println("task9 GetBranc>>>", err.Error())
			return
		}

		mymap[branch.Name] = v.Quantity * product.Price

	}
	i := 1
	for k, v := range mymap {
		fmt.Println(i, k, v)
		i++
	}

}

// Task 10
func (h *handler) Task10() {
	transaction, err := h.strg.BranchProductTransaction().GetListBranchProductTransaction()
	if err != nil {
		log.Println("task10 getlistbranchtaransaction>>>", err.Error())
		return
	}
	var result = make(map[string]int)
	for _, v := range transaction.BranchProductTransactions {
		product, err := h.strg.Product().GetProduct(models.IdRequestProduct{
			Id: v.ProductID,
		})
		if err != nil {
			log.Println("task10 getproduct>>>", err.Error())
			return
		}
		user, err := h.strg.User().GetUser(models.IdRequest{
			Id: v.UserID,
		})
		if err != nil {
			log.Println("task10 get user>>>", err.Error())
			return
		}
		result[user.Name] = v.Quantity * product.Price
	}
	for k, v := range result {
		fmt.Println(k, v)
	}

}

// Task 11
func (h *handler) Task11 () {

	user,err := h.strg.User().GetListUser()
	if err != nil {
		log.Println("task11 getlistuser>>>", err.Error())
		return
	}
	transaction,err := h.strg.BranchProductTransaction().GetListBranchProductTransaction()
	if err != nil {
		log.Println("task11 GetListBranchProductTransaction>>>", err.Error())
		return
	}
	var result =make(map[string]string)
	var result1 = make(map[int]int)

	
	


}

// Task 12
func (h *handler) Task12() {

	plus := 0
	minus := 0
	user, err := h.strg.User().GetListUser()
	if err != nil {
		log.Println("task 12>>>", err.Error())
		return
	}
	resp, err := h.strg.BranchProductTransaction().GetListBranchProductTransaction()
	if err != nil {
		log.Println("task 12>>>", err.Error())
		return
	}
	plusproducts := make(map[string]int)
	minusproducts := make(map[string]int)
	for _, value := range user.Users {
		for _, v := range resp.BranchProductTransactions {
			if value.Id == v.UserID {
				switch v.Type {
				case "plus":
					plus += v.Quantity
					plusproducts[value.Name] = plus
				case "minus":
					minus += v.Quantity
					fmt.Println(minus)
					minusproducts[value.Name] = minus
				}
			}
		}
		minus = 0
		plus = 0

	}
	fmt.Println("Ismi     Krim     Chqim")
	for _, v := range user.Users {
		fmt.Println(v.Name, "  ", plusproducts[v.Name], "  ", minusproducts[v.Name])
	}
}
