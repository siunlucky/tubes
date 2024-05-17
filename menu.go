package main

import "fmt"

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
				logout()
			}
		}
	}
}

var adminStart bool = true
var adminChoice int

func mainMenuAdmin() {
	for adminStart {
		fmt.Println("\nAdmin Menu")
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
				insertDataBank(&worldBank)
			case 2:
				viewDataCustomer(uniqueBankCode, worldBank)
			case 3:
				editDataBank(&worldBank)
			case 4:
				deleteDataBank(&worldBank)
			case 5:
				logout()
			}
		}
	}
}

func mainMenuCustomer() {
	dummy()
	for {
		fmt.Println("Menu Customer:")
		fmt.Println("1. Login")
		fmt.Println("2. Register")
		fmt.Println("3. Logout")
		fmt.Print("Pilih menu: ")

		var choice int
		fmt.Scan(&choice)

		switch choice {
		case 1:
			login()
		case 2:
			register()
		case 3:
			fmt.Println("Logout berhasil.")
			return
		default:
			fmt.Println("Pilihan tidak valid.")
		}
	}
}

func customerMenu(customer *Customer) {
	for {
		fmt.Println("Menu Customer:")
		fmt.Println("1. Cek Saldo")
		fmt.Println("2. Transfer")
		fmt.Println("3. Kembali")
		fmt.Print("Pilih menu: ")

		var choice int
		fmt.Scan(&choice)

		switch choice {
		case 1:
			viewSaldo(customer)
		case 2:
			transfer(customer)
		case 3:
			return
		default:
			fmt.Println("Pilihan tidak valid.")
		}
	}
}
