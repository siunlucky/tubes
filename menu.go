package main

import (
	"fmt"
	// "reflect"
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

var adminStart bool = true
var adminChoice int

func mainMenuAdmin() {
	for adminStart {
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
		for adminChoice >= 1 && adminChoice <= 5 {
			switch adminChoice {
			case 1:
				insertDataCustomer(uniqueBankCode, &worldBank)
			case 2:
				viewDataCustomer(uniqueBankCode, worldBank)
			case 3:
				editDataCustomer(uniqueBankCode, &worldBank)
			case 4:
				deleteDataCustomer(uniqueBankCode, &worldBank)
			case 5:
				logoutAdmin()
			}
		}
	}
}

func mainMenuCustomer() {
	dummyData()
	for {
		fmt.Println("========================================")
		fmt.Println("=             Menu Customer            =")
		fmt.Println("========================================")
		fmt.Println("1. Login")
		fmt.Print("Choose the menu option with number : ")

		var choice int
		fmt.Scan(&choice)

		if choice == 1 {
			loginCustomer()
		} else {
			fmt.Println("Input is not valid, please with right option")
			fmt.Println()
		}
	}
}

func customerMenu(customer *Customer) {
	var startCustomerMenu bool = true
	for startCustomerMenu {
		fmt.Println()
		fmt.Println("========================================")
		fmt.Println("=             Customer Menu            =")
		fmt.Println("1. View Balance")
		fmt.Println("2. Transfer")
		fmt.Println("3. Top-up Saldo")
		fmt.Println("4. Logout")
		fmt.Print("Choose the menu option with number : ")

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
			startCustomerMenu = false
		default:
			fmt.Println("Input is not valid, please input with right option")
		}
	}
}
