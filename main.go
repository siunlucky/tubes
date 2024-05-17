package main

import (
	"fmt"
)

var worldBank WorldBank
var customerBank CustomerBank

var startProgram bool = true

func main() {
	loading("Starting Program")
	for startProgram {
		menu()
	}
}

// main menu function to show the menu
func menu() {
	var startMenu bool = true
	var guestChoice int
	fmt.Println("What user are u?")
	fmt.Println("1. Super Admin")
	fmt.Println("2. Bank Admin")
	fmt.Println("3. Customer")
	fmt.Println("4. Exit")

	for startMenu {
		fmt.Print("Input the option in number : ")
		fmt.Scan(&guestChoice)
		if guestChoice == 1 || guestChoice == 2 || guestChoice == 3 || guestChoice == 4 {
			startMenu = false
		} else {
			fmt.Println("Please input the right number")
			fmt.Print("\n")
			startMenu = true
		}
	}

	switch guestChoice {
	case 1:
		clearTerminal()
		loginSuperAdmin()
		mainMenuSuperAdmin()
	case 2:
		clearTerminal()
		loginAdmin()
		mainMenuAdmin()
	case 3:
		mainMenuCustomer()
	case 4:
		exit()
	}
}

// exit function to exit the program
func exit() {
	var exit string
	fmt.Println("Are you sure want to exit? (Y/N)")

	fmt.Scan(&exit)
	for exit != "Y" && exit != "y" && exit != "N" && exit != "n" {
		fmt.Println("Please input the right option")
		fmt.Scan(&exit)
	}

	if exit == "Y" || exit == "y" {
		fmt.Println("Thank you for using our service")
		startProgram = false
	} else if exit == "N" || exit == "n" {
		fmt.Println("Thank you for using our service")
		fmt.Print("\n")
		menu()
	}
}