package main

import "fmt"

var adminStart bool = true
var adminChoice int

func mainMenuAdmin() {
	for adminStart {
		fmt.Println("\nAdmin Menu")
		fmt.Println("1. Insert Customer Data")
		fmt.Println("2. View Bank Data")
		fmt.Println("3. Edit Bank Data")
		fmt.Println("4. Delete Bank")
		fmt.Println("5. Logout")
		fmt.Print("Input : ")
		fmt.Scan(&adminChoice)
		for adminChoice >= 1 && adminChoice <= 5 {
			switch adminChoice {
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
