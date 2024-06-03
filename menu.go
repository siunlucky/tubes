package main

import (
	"fmt"
)

var superAdminStart bool = true
var superAdminChoice int

func mainMenuSuperAdmin() {
	for superAdminStart {
		fmt.Println("\nSuper Admin Menu")
		fmt.Println("1. Insert Bank Data")
		fmt.Println("2. View Bank Data")
		fmt.Println("3. Edit Bank Data")
		fmt.Println("4. Delete Bank")
		fmt.Println("5. Logout")
		fmt.Print("Input : ")
		fmt.Scan(&superAdminChoice)
		for superAdminChoice >= 1 && superAdminChoice <= 5 {
			switch superAdminChoice {
			case 1:
				insertDataBank(&worldBank)
			case 2:
				viewDataBank(worldBank)
			case 3:
				editDataBank(&worldBank)
			case 4:
				deleteDataBank(&worldBank)
			case 5:
				logoutSuperAdmin()
			}
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
		fmt.Println("5. Logout")
		fmt.Print("Input : ")
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
			break
		} else {
			fmt.Println("Input is not valid, please input with right option")
		}
	}
}

func mainMenuCustomer() {
	for {
		fmt.Println("========================================")
		fmt.Println("=             Menu Customer            =")
		fmt.Println("========================================")
		fmt.Println("1. Login")
		fmt.Println("2. Back to main menu")
		fmt.Print("Choose the menu option with number : ")

		var choice int
		fmt.Scan(&choice)

		if choice == 1 {
			loginCustomer()
		} else if choice == 2 {
			break
		} else {
			fmt.Println("Input is not valid, please input with right option")
		}

		if choice == 2 {
			break
		}
	}
}

func customerMenu(customer *Customer) {
	dummyBills()

	for {
		fmt.Println()
		fmt.Println("========================================")
		fmt.Println("=             Customer Menu            =")
		fmt.Println("1. View Balance")
		fmt.Println("2. Transfer")
		fmt.Println("3. Top-up Saldo")
		fmt.Println("4. Payment")
		fmt.Println("5. Logout")
		fmt.Print("Choose the menu option with number: ")

		var choice int
		fmt.Scan(&choice)

		switch choice {
		case 1:
			viewSaldo(customer)
		case 2:
			transfer(customer)
		case 3:
			topUpSaldo(customer)
		case 4:
			payment(customer)
		case 5:
			fmt.Println("Logging out...")
			break
		default:
			fmt.Println("Input is not valid, please input with right option")
		}
		if choice == 5 {
			break
		}
	}
}
