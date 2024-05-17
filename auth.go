package main

import (
	"fmt"
	"time"
)

// AUTH MENU

// loginSuperAdmin function to login as super admin
func loginSuperAdmin() {
	var credential Credential
	var maxTrial int = 1
	fmt.Println("Welcome to Dashboard World Bank System")
	fmt.Println("--------------Login Menu--------------")
	fmt.Print("Please Input Username : ")
	fmt.Scan(&credential.username)
	fmt.Print("Please Input Password : ")
	fmt.Scan(&credential.password)
	// coba cek iz, ini mending or apa and, gua bikin sendiri, tapi ambigu sendiri wkwk
	// Udah bener kok jadi ketika dia gagal login sama attempt kurang dari 4 dia bakal bakalan ngulang lagi -faiz
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
}

// loginAdmin function to login as admin

var uniqueBankCode int

func loginAdmin() {
	var credential Credential
	var maxTrial int

	maxTrial = 1
	fmt.Println("Welcome to Dashboard World Bank System")
	fmt.Println("--------------Login Menu--------------")
	fmt.Print("Please Input Unique Bank Code : ")
	fmt.Scan(&uniqueBankCode)
	fmt.Print("Please Input Username : ")
	fmt.Scan(&credential.username)
	fmt.Print("Please Input Password : ")
	fmt.Scan(&credential.password)
	// coba cek iz, ini mending or apa and, gua bikin sendiri, tapi ambigu sendiri wkwk
	// Udah bener kok jadi ketika dia gagal login sama attempt kurang dari 4 dia bakal bakalan ngulang lagi -faiz
	for !isAdmin(uniqueBankCode, credential) && maxTrial < 4 {
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
}

// login function to login as customer
func login() {
	var accountNumber int
	var PIN string
	for attempts := 0; attempts < 3; attempts++ {
		fmt.Print("Masukkan nomor rekening: ")
		fmt.Scan(&accountNumber)
		fmt.Print("Masukkan PIN: ")
		fmt.Scan(&PIN)

		for i := 0; i < customerBank.nCustomer; i++ {
			if customerBank.customers[i].accountNumber == accountNumber && customerBank.customers[i].PIN == PIN {
				fmt.Println("Login berhasil.")
				customerMenu(&customerBank.customers[i])
				return
			}
		}

		fmt.Println("Nomor rekening atau PIN salah.")
	}

	fmt.Println("Akun sudah masuk ke maksimal login, silakan coba lagi.")
}

// logoutSuperAdmin function to logout the super admin
func logout() {
	var logout string
	clearTerminal()
	fmt.Println("Are you sure want to logout? (Y/N)")

	fmt.Scan(&logout)
	for logout != "Y" && logout != "y" && logout != "N" && logout != "n" {
		clearTerminal()
		fmt.Println("Please input the right option")
		fmt.Scan(&logout)
	}

	if logout == "Y" || logout == "y" {
		loading("Logging Out")
		fmt.Println("Logging Out Success!\n")
		fmt.Println("Thank you for using our service!")
		fmt.Println("You will redirect to the main menu in 3 seconds")
		time.Sleep(3 * time.Second)
		clearTerminal()

		menu()
	} else if logout == "N" || logout == "n" {
		clearTerminal()
		uniqueBankCode = 0
		menu()
	}
}

// AUTH LOGIC

// isSuperAdmin function to check the credential
func isSuperAdmin(credential Credential) bool {
	if credential.username == SuperAdmin.username && credential.password == SuperAdmin.password {
		return true
	}
	return false
}

// isAdmin function to check the credential
func isAdmin(uniqueBankCode int, credential Credential) bool {
	var idx int

	idx = searchBankByUniqueCode(uniqueBankCode)

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
