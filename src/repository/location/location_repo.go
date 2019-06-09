package location

import "obas/src/repository"

type Account struct {
	Id      int
	Balance float64
}

func CreateAccount() {
	var db = repository.GetDB()
	defer db.Close()
	// Create
	db.Create(&Account{Id: 10, Balance: 1000})
}

func ReadAccount() Account {
	var db = repository.GetDB()
	defer db.Close()
	// Create
	var account Account
	db.First(&account, 1)            // find product with id 1
	db.First(&account, "Id = ?", 10) // find product with code l1212
	return account
}
