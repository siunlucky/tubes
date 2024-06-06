package main

import "fmt"

var startProgram bool = true

func main() {
	dummyData(5)
	loading("Starting Program")
	for startProgram {
		menu()
	}
}

func menu() {
	dummyBills()
	var startMenu bool = true
	var guestChoice int
	fmt.Println()
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
	case 2:
		clearTerminal()
		loginAdmin()
	case 3:
		clearTerminal()
		mainMenuCustomer()
	case 4:
		exit()
	}
}

func exit() {
	var exit string
	fmt.Println("Are you sure want to exit? (Y/N)")

	fmt.Scan(&exit)
	for exit != "Y" && exit != "y" && exit != "N" && exit != "n" {
		fmt.Println("Please input the right option Y / N")
		fmt.Scan(&exit)
	}

	if exit == "Y" || exit == "y" {
		fmt.Println("Thank you for coming")
		startProgram = false
	} else if exit == "N" || exit == "n" {
		fmt.Println("Bank to menu")
		fmt.Print("\n")
		menu()
	}
}
