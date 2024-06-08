package main

import "fmt"

func loginSuperAdmin() {
	var credential Credential
	var maxTrial int = 1
	fmt.Println()
	fmt.Println("============================================================")
	fmt.Println("=======    Welcome to Dashboard World Bank System    =======")
	fmt.Println("=======                  Login Menu                  =======")
	fmt.Println("============================================================")
	fmt.Print("Please Input Username : ")
	fmt.Scan(&credential.username)
	fmt.Print("Please Input Password : ")
	fmt.Scan(&credential.password)

	for !isSuperAdmin(credential) && maxTrial < 4 {
		maxTrial++
		if maxTrial == 4 {
			fmt.Println("===========================================================")
			fmt.Println("=== SORRY, YOU HAVE REACH THE MAXIMUM TRIAL FOR 3 TIMES ===")
			fmt.Println("===                   Back To Menu                      ===")
			fmt.Println("===========================================================")
			fmt.Print("\n")
			maxTrial = 4
			menu()
			return
		}

		fmt.Print("Login failed please try again\n")
		fmt.Print("Please Input Username : ")
		fmt.Scan(&credential.username)
		fmt.Print("Please Input Password : ")
		fmt.Scan(&credential.password)
	}

	fmt.Println("Login Success!")
	maxTrial = 4
	mainMenuSuperAdmin()
}

var uniqueBankCode int

func loginAdmin() {
	var credential Credential
	var maxTrial int

	maxTrial = 1
	fmt.Println()
	fmt.Println("============================================================")
	fmt.Println("=======    Welcome to Dashboard Admin Bank System    =======")
	fmt.Println("=======                  Login Menu                  =======")
	fmt.Println("============================================================")
	fmt.Print("Please Input Unique Bank Code : ")
	fmt.Scan(&uniqueBankCode)
	fmt.Print("Please Input Username : ")
	fmt.Scan(&credential.username)
	fmt.Print("Please Input Password : ")
	fmt.Scan(&credential.password)

	for !isAdmin(uniqueBankCode, credential) && maxTrial < 4 {
		maxTrial++
		if maxTrial == 4 {
			fmt.Println()
			fmt.Println("===========================================================")
			fmt.Println("=== SORRY, YOU HAVE REACH THE MAXIMUM TRIAL FOR 3 TIMES ===")
			fmt.Println("===                   Back To Menu                      ===")
			fmt.Println("===========================================================")
			fmt.Println()
			maxTrial = 4
			menu()
			return
		}

		fmt.Println("Login failed please try again")
		fmt.Println()
		fmt.Print("Please Input Unique Bank Code : ")
		fmt.Scan(&uniqueBankCode)
		fmt.Print("Please Input Username : ")
		fmt.Scan(&credential.username)
		fmt.Print("Please Input Password : ")
		fmt.Scan(&credential.password)
	}

	fmt.Println("Login Success!")
	maxTrial = 4
	mainMenuAdmin()
}

func loginCustomer() {
	var accountNumber int
	var PIN string
	var maxTrial int
	var customerIdx, bankIdx int

	maxTrial = 1
	fmt.Println()
	fmt.Println("============================================================")
	fmt.Println("=======   Welcome to Dashboard Customer Bank System  =======")
	fmt.Println("=======                  Login Menu                  =======")
	fmt.Println("============================================================")
	fmt.Print("Please Input Unique Bank Code : ")
	fmt.Scan(&uniqueBankCode)
	fmt.Print("Please Input Account Number : ")
	fmt.Scan(&accountNumber)
	fmt.Print("Please Input PIN : ")
	fmt.Scan(&PIN)

	for !isCustomer(uniqueBankCode, accountNumber, PIN) && maxTrial < 4 {
		maxTrial++
		if maxTrial == 4 {
			fmt.Println()
			fmt.Println("===========================================================")
			fmt.Println("=== SORRY, YOU HAVE REACH THE MAXIMUM TRIAL FOR 3 TIMES ===")
			fmt.Println("===                   Back To Menu                      ===")
			fmt.Println("===========================================================")
			fmt.Println()
			maxTrial = 4
			menu()
			return
		}

		fmt.Println("Login failed please try again")
		fmt.Println()
		fmt.Print("Please Input Unique Customer Bank Code : ")
		fmt.Scan(&uniqueBankCode)
		fmt.Print("Please Input Account Number : ")
		fmt.Scan(&accountNumber)
		fmt.Print("Please Input PIN : ")
		fmt.Scan(&PIN)
	}

	bankIdx = searchBankByUniqueCode(uniqueBankCode)
	customerIdx = searchCustomerByCredentials(bankIdx, accountNumber, PIN)

	fmt.Println("Login Success!")
	maxTrial = 4

	customerMenu(&worldBank.Banks[bankIdx].customers[customerIdx])
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
	return false
}

func isCustomer(uniqueBankCode int, accountNumber int, PIN string) bool {
	var idx int = searchBankByUniqueCode(uniqueBankCode)
	if idx == -1 {
		return false
	}

	for i := 0; i < worldBank.Banks[idx].nCustomer; i++ {
		if accountNumber == worldBank.Banks[idx].customers[i].accountNumber && PIN == worldBank.Banks[idx].customers[i].PIN {
			return true
		}
	}
	return false
}
