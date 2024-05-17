package main

import "fmt"

// mainMenuSuperAdmin function to show the main menu
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
