package main

import "fmt"



func loginSuperAdmin() {
	var credential Credential
	var maxTrial int = 1
	fmt.Println()
	fmt.Println("Welcome to Dashboard World Bank System")
	fmt.Println("--------------Login Menu--------------")
	fmt.Print("Please Input Username : ")
	fmt.Scan(&credential.username)
	fmt.Print("Please Input Password : ")
	fmt.Scan(&credential.password)
	for !isSuperAdmin(credential) && maxTrial < 4 {
		loading("Checking Credentials")
		maxTrial++
		if maxTrial == 4 {
			fmt.Println("-- SORRY, YOU HAVE REACH THE MAXIMUM TRIAL FOR 3 TIMES --")
			fmt.Println("--                   Back To Menu                      --")
			fmt.Print("\n")
			maxTrial = 4
			menu()
		}

		fmt.Print("Login failed please try again\n")
		fmt.Print("Please Input Username : ")
		fmt.Scan(&credential.username)
		fmt.Print("Please Input Password : ")
		fmt.Scan(&credential.password)
	}
	loading("Checking Credentials")
	fmt.Println("Login Success!")
	mainMenuSuperAdmin()
}

var uniqueBankCode int
func loginAdmin() {
	var credential Credential
	var maxTrial int

	maxTrial = 1
	fmt.Println()
	fmt.Println("Welcome to Dashboard World Bank System")
	fmt.Println("--------------Login Menu--------------")
	fmt.Print("Please Input Unique Bank Code : ")
	fmt.Scan(&uniqueBankCode)
	fmt.Print("Please Input Username : ")
	fmt.Scan(&credential.username)
	fmt.Print("Please Input Password : ")
	fmt.Scan(&credential.password)
	for !isAdmin(uniqueBankCode, credential) && maxTrial < 4 {
		maxTrial++
		if maxTrial == 4 {
			fmt.Println("-- SORRY, YOU HAVE REACH THE MAXIMUM TRIAL FOR 3 TIMES --")
			fmt.Println("--                   Back To Menu                      --")
			fmt.Print("\n")
			maxTrial = 4
			menu()
		}

		fmt.Print("Login failed please try again\n")
		fmt.Print("Please Input Unique Bank Code : ")
		fmt.Scan(&uniqueBankCode)
		fmt.Print("Please Input Username : ")
		fmt.Scan(&credential.username)
		fmt.Print("Please Input Password : ")
		fmt.Scan(&credential.password)
	}

	fmt.Println("Login Success!")
	mainMenuAdmin()
}

func loginCustomer() {
	var accountNumber int
	var PIN string

	for attempts := 0; attempts < 3; attempts++ {
		fmt.Println()
		fmt.Println("========================================")
		fmt.Print("Input Uniqode Bank : ")
		fmt.Scan(&uniqueBankCode)
		fmt.Print("Input Account Number : ")
		fmt.Scan(&accountNumber)
		fmt.Print("Input PIN : ")
		fmt.Scan(&PIN)

		var bankIdx int = searchBankByUniqueCode(uniqueBankCode)

		if bankIdx == -1 {
			fmt.Println("Bank not found.")
		}

		// Sequential Search
		for i := 0; i < worldBank.Banks[bankIdx].nCustomer; i++ {
			if worldBank.Banks[bankIdx].customers[i].accountNumber == accountNumber && worldBank.Banks[bankIdx].customers[i].PIN == PIN {
				// loadingAuth("Checking Credentials")
				fmt.Println("Login Success!")
				customerMenu(&worldBank.Banks[bankIdx].customers[i])
				return
			}
		}
		// loadingAuth("Checking Credentials")
		fmt.Println("Your account number or PIN is incorrect")
	}
	fmt.Println()
	fmt.Println("You have reached the maximum number for 3 attempts, try again")
}

func isSuperAdmin(credential Credential) bool {
	if credential.username == SuperAdmin.username && credential.password == SuperAdmin.password {
		return true
	}
	return false
}

func isAdmin(uniqueBankCode int, credential Credential) bool {
	var idx int = searchBankByUniqueCode(uniqueBankCode)
	if idx == -1 {
		return false
	}

	for i := 0; i < worldBank.Banks[idx].nAdmin; i++ {
		if credential.username == worldBank.Banks[idx].admins[i].username && credential.password == worldBank.Banks[idx].admins[i].password {
			return true
		}
	}

	return true
}