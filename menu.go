package main

import "fmt"

func mainMenuSuperAdmin() {
	var superAdminChoice int
	for {
		fmt.Println("\nSuper Admin Menu")
		fmt.Println("1. Insert Bank Data")
		fmt.Println("2. View Bank Data")
		fmt.Println("3. Edit Bank Data")
		fmt.Println("4. Delete Bank")
		fmt.Println("5. Logout")
		fmt.Print("Input : ")
		fmt.Scan(&superAdminChoice)
		if superAdminChoice == 1 {
			insertDataBank(&worldBank)
		} else if superAdminChoice == 2 {
			viewDataBank(worldBank)
		} else if superAdminChoice == 3 {
			editDataBank(&worldBank)
		} else if superAdminChoice == 4 {
			deleteDataBank(&worldBank)
		} else if superAdminChoice == 5 {
			break
		} else {
			fmt.Println("Input is not valid, please input with right option")
		}
	}
}

var adminChoice int

func mainMenuAdmin() {
	for {
		fmt.Println("\n========================================")
		fmt.Println("=              Menu Admin              =")
		fmt.Println("========================================")
		fmt.Println("1. Insert Customer Data")
		fmt.Println("2. View Customer Data")
		fmt.Println("3. Edit Customer Data")
		fmt.Println("4. Delete Customer")
		fmt.Println("5. Search Customer Data")
		fmt.Println("6. Sort Customer Data")
		fmt.Println("7. Logout")
		fmt.Print("Input : ")
		var adminChoice int
		fmt.Scan(&adminChoice)
		if adminChoice == 1 {
			insertDataCustomer(uniqueBankCode, &worldBank)
		} else if adminChoice == 2 {
			viewDataCustomer(uniqueBankCode, worldBank)
		} else if adminChoice == 3 {
			editDataCustomer(uniqueBankCode, &worldBank)
		} else if adminChoice == 4 {
			deleteDataCustomer(uniqueBankCode, &worldBank)
		} else if adminChoice == 5 {
			searchDataCustomerMenu()
		} else if adminChoice == 6 {
			sortDataCustomerMenu()
		} else if adminChoice == 7 {
			break
		} else {
			fmt.Println("Input is not valid, please input with right option")
		}
	}
}

func searchDataCustomerMenu() {
	for {
		fmt.Println()
		fmt.Println("========================================")
		fmt.Println("=            Search Customer Data      =")
		fmt.Println("========================================")
		fmt.Println("1. Search by Name")
		fmt.Println("2. Search by Account Number")
		fmt.Println("3. Search by Card Number")
		fmt.Println("4. Search by NIK")
		fmt.Println("5. Back to Admin Menu")
		fmt.Print("Input : ")
		var searchChoice int
		var bankIdx int
		var customerIdx int
		bankIdx = searchBankByUniqueCode(uniqueBankCode)
		fmt.Scan(&searchChoice)
		if searchChoice == 1 {
			fmt.Print("Input name to search : ")
			var searchName string
			fmt.Scan(&searchName)
			customerIdx = searchCustomerByName(bankIdx, searchName)
			printCustomerData(bankIdx, customerIdx)
		} else if searchChoice == 2 {
			fmt.Print("Input account number to search : ")
			var searchAccountNumber int
			fmt.Scan(&searchAccountNumber)
			customerIdx = searchCustomerByAccountNumber(bankIdx, searchAccountNumber)
			printCustomerData(bankIdx, customerIdx)
		} else if searchChoice == 3 {
			fmt.Print("Input card number to search : ")
			var searchCardNumber int
			fmt.Scan(&searchCardNumber)
			customerIdx = searchCustomerByCardNumber(bankIdx, searchCardNumber)
			printCustomerData(bankIdx, customerIdx)
		} else if searchChoice == 4 {
			fmt.Print("Input NIK to search : ")
			var searchNIK int
			fmt.Scan(&searchNIK)
			customerIdx = searchCustomerByNIK(bankIdx, searchNIK)
			printCustomerData(bankIdx, customerIdx)
		} else if searchChoice == 5 {
			break
		} else {
			fmt.Println("Input is not valid, please input with right option")
		}
	}
}

func printCustomerData(bankIdx, customerIdx int) {
	if customerIdx == -1 {
		fmt.Println("Customer not found")
	} else {
		var address string
		address = worldBank.Banks[bankIdx].customers[customerIdx].address.district + ", " + worldBank.Banks[bankIdx].customers[customerIdx].address.city + ", " + worldBank.Banks[bankIdx].customers[customerIdx].address.province
		fmt.Println("Name : ", worldBank.Banks[bankIdx].customers[customerIdx].name)
		fmt.Println("Account Number : ", worldBank.Banks[bankIdx].customers[customerIdx].accountNumber)
		fmt.Println("Card Number : ", worldBank.Banks[bankIdx].customers[customerIdx].cardNumber)
		fmt.Println("Balance : ", worldBank.Banks[bankIdx].customers[customerIdx].balance)
		fmt.Println("Address : ", address)
	}
}

func sortDataCustomerMenu() {
	for {
		fmt.Println()
		fmt.Println("========================================")
		fmt.Println("=            Sort Customer Data        =")
		fmt.Println("========================================")
		fmt.Println("1. Sort by Name (A-Z)")
		fmt.Println("2. Sort by Account Number (Ascending)")
		fmt.Println("3. Sort by Card Number (Descending)")
		fmt.Println("4. Back to Admin Menu")
		fmt.Print("Input : ")
		var sortChoice int
		fmt.Scan(&sortChoice)
		if sortChoice == 1 {
			sortCustomerByName()
			viewDataCustomer(uniqueBankCode, worldBank)
		} else if sortChoice == 2 {
			sortCustomerByAccountNumber()
			viewDataCustomer(uniqueBankCode, worldBank)
		} else if sortChoice == 3 {
			sortCustomerByCardNumber()
			viewDataCustomer(uniqueBankCode, worldBank)
		} else if sortChoice == 4 {
			break
		} else {
			fmt.Println("Input is not valid, please input with right option")
		}
	}
}

func customerMenu(customer *Customer) {
	for {
		fmt.Println()
		fmt.Println("========================================")
		fmt.Println("=             Customer Menu            =")
		fmt.Println("1. View Balance")
		fmt.Println("2. Transfer")
		fmt.Println("3. Top-up Saldo")
		fmt.Println("4. Payment")
		fmt.Println("5. Transaction History")
		fmt.Println("6. Logout")
		fmt.Print("Choose the menu option with number: ")

		var choice int
		fmt.Scan(&choice)

		if choice == 1 {
			viewSaldo(customer)
		} else if choice == 2 {
			transfer(customer)
		} else if choice == 3 {
			topUpSaldo(customer)
		} else if choice == 4 {
			payment(customer)
		} else if choice == 5 {
			transactionHistory(customer)
		} else if choice == 6 {
			break
		} else {
			fmt.Println("Input is not valid, please input with right option")
		}
	}
}

// func mainMenuCustomer() {
// 	for {
// 		fmt.Println()
// 		fmt.Println("========================================")
// 		fmt.Println("=          Menu Login Customer         =")
// 		fmt.Println("========================================")
// 		fmt.Println("1. Login")
// 		fmt.Println("2. Back to main menu")
// 		fmt.Print("Choose the menu option with number : ")

// 		var choice int
// 		fmt.Scan(&choice)

// 		if choice == 1 {
// 			loginCustomer()
// 		} else if choice == 2 {
// 			break
// 		} else {
// 			fmt.Println("Input is not valid, please input with right option")
// 		}
// 	}
// }
