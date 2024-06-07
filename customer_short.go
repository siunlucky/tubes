package main

import "fmt"

// insertion sort A-Z
func sortCustomerByName() {
	bankIdx := searchBankByUniqueCode(uniqueBankCode)
	if bankIdx == -1 {
		fmt.Println("Bank not found")
		return
	}

	for pass := 1; pass < worldBank.Banks[bankIdx].nCustomer; pass++ {
		temp := worldBank.Banks[bankIdx].customers[pass]
		i := pass
		for i > 0 && worldBank.Banks[bankIdx].customers[i-1].name > temp.name {
			worldBank.Banks[bankIdx].customers[i] = worldBank.Banks[bankIdx].customers[i-1]
			i--
		}
		worldBank.Banks[bankIdx].customers[i] = temp
	}
}

// insertion sort ascending
func sortCustomerByAccountNumber() {
	bankIdx := searchBankByUniqueCode(uniqueBankCode)
	if bankIdx == -1 {
		fmt.Println("Bank not found")
		return
	}

	for pass := 1; pass < worldBank.Banks[bankIdx].nCustomer; pass++ {
		temp := worldBank.Banks[bankIdx].customers[pass]
		i := pass
		for i > 0 && worldBank.Banks[bankIdx].customers[i-1].accountNumber > temp.accountNumber {
			worldBank.Banks[bankIdx].customers[i] = worldBank.Banks[bankIdx].customers[i-1]
			i--
		}
		worldBank.Banks[bankIdx].customers[i] = temp
	}
}

// selection sort descending
func sortCustomerByCardNumber() {
	bankIdx := searchBankByUniqueCode(uniqueBankCode)
	if bankIdx == -1 {
		fmt.Println("Bank not found")
		return
	}

	for pass := 0; pass < worldBank.Banks[bankIdx].nCustomer-1; pass++ {
		idx := pass
		for i := pass + 1; i < worldBank.Banks[bankIdx].nCustomer; i++ {
			if worldBank.Banks[bankIdx].customers[i].cardNumber > worldBank.Banks[bankIdx].customers[idx].cardNumber {
				idx = i
			}
		}
		temp := worldBank.Banks[bankIdx].customers[pass]
		worldBank.Banks[bankIdx].customers[pass] = worldBank.Banks[bankIdx].customers[idx]
		worldBank.Banks[bankIdx].customers[idx] = temp
	}
}












// func sortCustomerByName() {
// 	var bankIdx, pass, i int
// 	var temp Customer
// 	bankIdx = searchBankByUniqueCode(uniqueBankCode)
// 	pass = 1

// 	if bankIdx == -1 {
// 		return
// 	}

// 	for pass <= customerBank.nCustomer -1 {
// 		i = pass
// 		temp = customerBank.customers[pass]
// 		for i > 0 && customerBank.customers[i-1].name > temp.name {
// 			customerBank.customers[i] = customerBank.customers[i-1]
// 			i = i - 1
// 		}
// 		customerBank.customers[i] = temp
// 		pass = pass + 1
// 	}	
// }

// func sortCustomerByAccountNumber() {
// 	var bankIdx, pass, i int
// 	var temp Customer
// 	bankIdx = searchBankByUniqueCode(uniqueBankCode)
// 	pass = 1

// 	if bankIdx == -1 {
// 		return
// 	}

// 	for pass <= customerBank.nCustomer -1 {
// 		i = pass
// 		temp = customerBank.customers[pass]
// 		for i > 0 && customerBank.customers[i-1].accountNumber > temp.accountNumber {
// 			customerBank.customers[i] = customerBank.customers[i-1]
// 			i = i - 1
// 		}
// 		customerBank.customers[i] = temp
// 		pass = pass + 1
// 	}
// }