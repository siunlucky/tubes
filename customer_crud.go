package main

import "fmt"


func viewDataCustomer(uniqueBankCode int, worldBank WorldBank) {
	var startViewCustomer bool = true

	var bankIdx int = searchBankByUniqueCode(uniqueBankCode)

	if bankIdx == -1 {
		fmt.Println("Bank not found.")
		return
	}

	for startViewCustomer {
		if worldBank.nBank == 0 {
			fmt.Println("There is no data to be viewed.")
			startViewCustomer = false
			adminChoice = 0
		} else {
			fmt.Println()
			fmt.Println("Customer Data :")
			fmt.Printf("%-5s %-20s %-20s %-20s %-20s %-20s %-20s %20s\n", "Account Number", "Card Number", "PIN", "NIK", "Name", "Address", "Balance", "Is Suspended")
			for i := 0; i < worldBank.Banks[bankIdx].nCustomer; i++ {
				fmt.Printf("%-5d %-20d %-20d %-20d %-20s %-20s%s %s %-20d %20b\n", worldBank.Banks[bankIdx].customers[i].accountNumber, worldBank.Banks[bankIdx].customers[i].cardNumber, worldBank.Banks[bankIdx].customers[i].PIN, worldBank.Banks[bankIdx].customers[i].NIK, worldBank.Banks[bankIdx].customers[i].name, worldBank.Banks[bankIdx].customers[i].address.district, worldBank.Banks[bankIdx].customers[i].address.city, worldBank.Banks[bankIdx].customers[i].address.province, worldBank.Banks[bankIdx].customers[i].balance, worldBank.Banks[bankIdx].customers[i].isSuspended)
			}
			fmt.Println("Total Data : ", worldBank.Banks[bankIdx].nCustomer)
			startViewCustomer = false
			adminChoice = 0
		}
	}
}

func insertDataCustomer(uniqueBankCode int, worldBank *WorldBank) {
	var retry string = "y"
	var found bool

	var startInsertCustomer bool = true

	var bankIdx int = searchBankByUniqueCode(uniqueBankCode)

	if bankIdx == -1 {
		fmt.Println("Bank not found.")
		return
	}

	for startInsertCustomer {
		fmt.Println()
		fmt.Println("Customer", worldBank.Banks[bankIdx].nCustomer+1)

		for worldBank.Banks[bankIdx].customers[worldBank.Banks[bankIdx].nCustomer].NIK < 1 || found {
			found = false

			fmt.Print("NIK : ")
			fmt.Scan(&worldBank.Banks[bankIdx].customers[worldBank.Banks[bankIdx].nCustomer].NIK)

			found = searchCustomerByNIK(bankIdx, worldBank.Banks[bankIdx].customers[worldBank.Banks[bankIdx].nCustomer].NIK) != -1

			if worldBank.Banks[bankIdx].customers[worldBank.Banks[bankIdx].nCustomer].NIK == 0 {
				fmt.Println("NIK field may not be empty.")
			} else if worldBank.Banks[bankIdx].customers[worldBank.Banks[bankIdx].nCustomer].NIK < 0 {
				fmt.Println("NIK is invalid.")
			} else if found {
				fmt.Println("NIK is already exist.")
			}
		}

		for worldBank.Banks[bankIdx].customers[worldBank.Banks[bankIdx].nCustomer].name == "" || len(worldBank.Banks[bankIdx].customers[worldBank.Banks[bankIdx].nCustomer].name) < 3 {
			fmt.Print("Name : ")
			fmt.Scan(&worldBank.Banks[bankIdx].customers[worldBank.Banks[bankIdx].nCustomer].name)
			if worldBank.Banks[bankIdx].customers[worldBank.Banks[bankIdx].nCustomer].name == "" {
				fmt.Println("Name field may not be empty.")
			} else if len(worldBank.Banks[bankIdx].customers[worldBank.Banks[bankIdx].nCustomer].name) < 3 {
				fmt.Println("Name should be contain more than 3 char.")
			}
		}

		fmt.Println("Address :")

		for worldBank.Banks[bankIdx].customers[worldBank.Banks[bankIdx].nCustomer].address.district == "" || len(worldBank.Banks[bankIdx].customers[worldBank.Banks[bankIdx].nCustomer].address.district) < 3 {
			fmt.Print("District : ")
			fmt.Scan(&worldBank.Banks[bankIdx].customers[worldBank.Banks[bankIdx].nCustomer].address.district)
			if worldBank.Banks[bankIdx].customers[worldBank.Banks[bankIdx].nCustomer].address.district == "" {
				fmt.Println("District field may not be empty.")
			} else if len(worldBank.Banks[bankIdx].customers[worldBank.Banks[bankIdx].nCustomer].address.district) < 3 {
				fmt.Println("District should be contain more than 3 char.")
			}
		}

		for worldBank.Banks[bankIdx].customers[worldBank.Banks[bankIdx].nCustomer].address.city == "" || len(worldBank.Banks[bankIdx].customers[worldBank.Banks[bankIdx].nCustomer].address.city) < 3 {
			fmt.Print("City : ")
			fmt.Scan(&worldBank.Banks[bankIdx].customers[worldBank.Banks[bankIdx].nCustomer].address.city)
			if worldBank.Banks[bankIdx].customers[worldBank.Banks[bankIdx].nCustomer].address.city == "" {
				fmt.Println("City field may not be empty.")
			} else if len(worldBank.Banks[bankIdx].customers[worldBank.Banks[bankIdx].nCustomer].address.city) < 3 {
				fmt.Println("City should be contain more than 3 char.")
			}
		}

		for worldBank.Banks[bankIdx].customers[worldBank.Banks[bankIdx].nCustomer].address.province == "" || len(worldBank.Banks[bankIdx].customers[worldBank.Banks[bankIdx].nCustomer].address.province) < 3 {
			fmt.Print("Province : ")
			fmt.Scan(&worldBank.Banks[bankIdx].customers[worldBank.Banks[bankIdx].nCustomer].address.province)
			if worldBank.Banks[bankIdx].customers[worldBank.Banks[bankIdx].nCustomer].address.province == "" {
				fmt.Println("Province field may not be empty.")
			} else if len(worldBank.Banks[bankIdx].customers[worldBank.Banks[bankIdx].nCustomer].address.province) < 3 {
				fmt.Println("Province should be contain more than 3 char.")
			}
		}

		for worldBank.Banks[bankIdx].customers[worldBank.Banks[bankIdx].nCustomer].PIN == "" || len(worldBank.Banks[bankIdx].customers[worldBank.Banks[bankIdx].nCustomer].PIN) != 6 {
			fmt.Print("PIN : ")
			fmt.Scan(&worldBank.Banks[bankIdx].customers[worldBank.Banks[bankIdx].nCustomer].PIN)
			if worldBank.Banks[bankIdx].customers[worldBank.Banks[bankIdx].nCustomer].PIN == "" {
				fmt.Println("PIN field may not be empty.")
			} else if len(worldBank.Banks[bankIdx].customers[worldBank.Banks[bankIdx].nCustomer].PIN) != 6 {
				fmt.Println("PIN should be contain 6  char.")
			}
		}

		worldBank.Banks[bankIdx].customers[worldBank.Banks[bankIdx].nCustomer].isSuspended = false
		worldBank.Banks[bankIdx].customers[worldBank.Banks[bankIdx].nCustomer].balance = 0
		worldBank.Banks[bankIdx].customers[worldBank.Banks[bankIdx].nCustomer].nTransaction = 0
		worldBank.Banks[bankIdx].nCustomer++

		fmt.Println("Customer Data Created Successfully")

		fmt.Println("\nWill Create Customer Data Again? (Y/N)")
		fmt.Scan(&retry)
		var startRetry bool = true
		for startRetry {
			if retry == "N" || retry == "n" {
				startRetry = false
				startInsertCustomer = false
				mainMenuAdmin()
			} else if retry == "Y" || retry == "y" {
				startRetry = false
				startInsertCustomer = true
			} else {
				fmt.Println("Please input the right option")
				fmt.Scan(&retry)
			}
		}
	}
}

func topUpSaldo(customer *Customer) {
	var amount int
	fmt.Print("The minimum top-up amount is Rp 50.000. Input the amount you want to top-up: Rp.")
	fmt.Scan(&amount)

	if amount < 50000 {
		fmt.Println("Amount is less than the minimum top-up amount, please input more than Rp 50.000.")
		return
	}

	customer.balance += amount
	fmt.Printf("Top-up success, now your balance is Rp.%d\n", customer.balance)
}