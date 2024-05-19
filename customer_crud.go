package main

import "fmt"

func viewDataCustomer(uniqueBankCode int, worldBank WorldBank) {
	var startViewCustomer bool = true

	var bankIdx int = searchBankByUniqueCode(uniqueBankCode)

	var address string

	if bankIdx == -1 {
		fmt.Println("Bank not found.")
		mainMenuAdmin()
	}

	for startViewCustomer {
		if worldBank.Banks[bankIdx].nCustomer == 0 {
			fmt.Println("There is no data to be viewed.")
			startViewCustomer = false
			adminChoice = 0
		} else {
			fmt.Println()
			fmt.Println("Customer Data :")
			fmt.Printf("%-5s %-20s %-20s %-20s %-20s %-20s %-20s %-20s %20s\n", "No", "Account Number", "Card Number", "PIN", "NIK", "Name", "Address", "Balance", "Is Suspended")
			for i := 0; i < worldBank.Banks[bankIdx].nCustomer; i++ {
				address = worldBank.Banks[bankIdx].customers[i].address.district + ", " + worldBank.Banks[bankIdx].customers[i].address.city + ", " + worldBank.Banks[bankIdx].customers[i].address.province
				fmt.Printf("%-5d %-20d %-20d %-20s %-20d %-20s %-20s %-20d %20t\n", i+1, worldBank.Banks[bankIdx].customers[i].accountNumber, worldBank.Banks[bankIdx].customers[i].cardNumber, worldBank.Banks[bankIdx].customers[i].PIN, worldBank.Banks[bankIdx].customers[i].NIK, worldBank.Banks[bankIdx].customers[i].name, address, worldBank.Banks[bankIdx].customers[i].balance, worldBank.Banks[bankIdx].customers[i].isSuspended)
			}
			fmt.Println("Total Data :", worldBank.Banks[bankIdx].nCustomer)
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

		fmt.Println("\nAddress")

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
		worldBank.Banks[bankIdx].customers[worldBank.Banks[bankIdx].nCustomer].cardNumber = codeGenerator(1000000000000000, 9999999999999999)
		worldBank.Banks[bankIdx].customers[worldBank.Banks[bankIdx].nCustomer].accountNumber = codeGenerator(10000000, 99999999)
		for searchCustomerByAccountNumber(bankIdx, worldBank.Banks[bankIdx].customers[worldBank.Banks[bankIdx].nCustomer].accountNumber) != -1 || searchCustomerByCardNumber(bankIdx, worldBank.Banks[bankIdx].customers[worldBank.Banks[bankIdx].nCustomer].cardNumber) != -1 {
			if searchCustomerByAccountNumber(bankIdx, worldBank.Banks[bankIdx].customers[worldBank.Banks[bankIdx].nCustomer].accountNumber) != -1 {
				worldBank.Banks[bankIdx].customers[worldBank.Banks[bankIdx].nCustomer].accountNumber = codeGenerator(10000000, 99999999)
			}

			if searchCustomerByCardNumber(bankIdx, worldBank.Banks[bankIdx].customers[worldBank.Banks[bankIdx].nCustomer].cardNumber) != -1 {
				worldBank.Banks[bankIdx].customers[worldBank.Banks[bankIdx].nCustomer].cardNumber = codeGenerator(1000000000000000, 9999999999999999)
			}
		}
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

func editDataCustomer(uniqueBankCode int, worldBank *WorldBank) {
	var customerIndex, bankIdx, choice int

	bankIdx = searchBankByUniqueCode(uniqueBankCode)

	if worldBank.Banks[bankIdx].nCustomer == 0 {
		fmt.Println("There is no data to be edited.")
		mainMenuAdmin()
	}

	for customerIndex != -1 {
		viewDataCustomer(uniqueBankCode, *worldBank)

		fmt.Print("Input Customer Number (input -1 for cancel) : ")
		fmt.Scan(&customerIndex)
		for searchCustomerByIdx(bankIdx, customerIndex-1) == -1 && customerIndex != -1 {
			fmt.Println("Customer Number is not available.")
			fmt.Print("Input Customer Number (input -1 for cancel) : ")
			fmt.Scan(&customerIndex)
		}

		for choice != 5 && customerIndex != -1 {
			fmt.Println("1. Edit Name")
			fmt.Println("2. Edit Address")
			fmt.Println("3. Edit Suspended Status")
			fmt.Println("4. Reset PIN")
			fmt.Println("5. Save")
			fmt.Print("Input : ")
			fmt.Scan(&choice)
			for choice < 1 || choice > 5 {
				fmt.Print("Input : ")
				fmt.Scan(&choice)
			}

			switch choice {
			case 1:
				fmt.Println("Old Name :", worldBank.Banks[bankIdx].customers[customerIndex-1].name)
				fmt.Print("New Name : ")
				fmt.Scan(&worldBank.Banks[bankIdx].customers[customerIndex-1].name)
				for worldBank.Banks[bankIdx].customers[customerIndex-1].name == "" || len(worldBank.Banks[bankIdx].customers[customerIndex-1].name) < 3 {
					if worldBank.Banks[bankIdx].customers[customerIndex-1].name == "" {
						fmt.Println("Name field may not be empty.")
					} else if len(worldBank.Banks[bankIdx].customers[customerIndex-1].name) < 3 {
						fmt.Println("Name should be contain more than 3 char.")
					}
					fmt.Print("New Name : ")
					fmt.Scan(&worldBank.Banks[bankIdx].customers[customerIndex-1].name)
				}
				fmt.Print("Customer Name Edited Successfully\n")
				viewDataCustomer(uniqueBankCode, *worldBank)
				fmt.Println()
			case 2:
				fmt.Println("Old Address :", worldBank.Banks[bankIdx].customers[customerIndex-1].address.district, worldBank.Banks[bankIdx].customers[customerIndex-1].address.city, worldBank.Banks[bankIdx].customers[customerIndex-1].address.province)
				fmt.Print("New Address")
				fmt.Print("District : ")
				fmt.Scan(&worldBank.Banks[bankIdx].customers[customerIndex-1].address.district)
				for worldBank.Banks[bankIdx].customers[customerIndex-1].address.district == "" || len(worldBank.Banks[bankIdx].customers[customerIndex-1].address.district) < 3 {
					if worldBank.Banks[bankIdx].customers[customerIndex-1].address.district == "" {
						fmt.Println("District field may not be empty.")
					} else if len(worldBank.Banks[bankIdx].customers[customerIndex-1].address.district) < 3 {
						fmt.Println("District should be contain more than 3 char.")
					}
					fmt.Print("District : ")
					fmt.Scan(&worldBank.Banks[bankIdx].customers[customerIndex-1].address.district)
				}
				fmt.Print("City : ")
				fmt.Scan(&worldBank.Banks[bankIdx].customers[customerIndex-1].address.city)
				for worldBank.Banks[bankIdx].customers[customerIndex-1].address.city == "" || len(worldBank.Banks[bankIdx].customers[customerIndex-1].address.city) < 3 {
					if worldBank.Banks[bankIdx].customers[customerIndex-1].address.city == "" {
						fmt.Println("City field may not be empty.")
					} else if len(worldBank.Banks[bankIdx].customers[customerIndex-1].address.city) < 3 {
						fmt.Println("City should be contain more than 3 char.")
					}
					fmt.Print("City : ")
					fmt.Scan(&worldBank.Banks[bankIdx].customers[customerIndex-1].address.city)
				}
				fmt.Print("Province : ")
				fmt.Scan(&worldBank.Banks[bankIdx].customers[customerIndex-1].address.province)
				for worldBank.Banks[bankIdx].customers[customerIndex-1].address.province == "" || len(worldBank.Banks[bankIdx].customers[customerIndex-1].address.province) < 3 {
					if worldBank.Banks[bankIdx].customers[customerIndex-1].address.province == "" {
						fmt.Println("Province field may not be empty.")
					} else if len(worldBank.Banks[bankIdx].customers[customerIndex-1].address.province) < 3 {
						fmt.Println("Province should be contain more than 3 char.")
					}
					fmt.Print("Province : ")
					fmt.Scan(&worldBank.Banks[bankIdx].customers[customerIndex-1].address.province)
				}
				fmt.Print("Customer Address Edited Successfully\n")
				viewDataCustomer(uniqueBankCode, *worldBank)
				fmt.Println()
			case 3:
				worldBank.Banks[bankIdx].customers[customerIndex-1].isSuspended = !worldBank.Banks[bankIdx].customers[customerIndex-1].isSuspended
				fmt.Print("Customer Suspended Status Edited Successfully\n")
				viewDataCustomer(uniqueBankCode, *worldBank)
				fmt.Println()
			case 4:
				fmt.Print("New PIN : ")
				fmt.Scan(&worldBank.Banks[bankIdx].customers[customerIndex-1].PIN)
				for worldBank.Banks[bankIdx].customers[customerIndex-1].PIN == "" || len(worldBank.Banks[bankIdx].customers[customerIndex-1].PIN) != 6 {
					if worldBank.Banks[bankIdx].customers[customerIndex-1].PIN == "" {
						fmt.Println("PIN field may not be empty.")
					} else if len(worldBank.Banks[bankIdx].customers[customerIndex-1].PIN) != 6 {
						fmt.Println("PIN should be contain 6 char.")
					}
					fmt.Print("New PIN : ")
					fmt.Scan(&worldBank.Banks[bankIdx].customers[customerIndex-1].PIN)
				}
				fmt.Print("Customer PIN Reset Successfully\n")
				viewDataCustomer(uniqueBankCode, *worldBank)
				fmt.Println()
			}
		}
	}
}

func deleteDataCustomer(uniqueBankCode int, worldBank *WorldBank) {
	var customerIndex int

	var bankIdx int = searchBankByUniqueCode(uniqueBankCode)

	if worldBank.Banks[bankIdx].nCustomer == 0 {
		fmt.Println("There is no data to be deleted.")
		mainMenuAdmin()
	}

	for customerIndex != -1 {
		viewDataCustomer(uniqueBankCode, *worldBank)

		fmt.Print("Input Customer Number (input -1 for cancel) : ")
		fmt.Scan(&customerIndex)
		for searchCustomerByIdx(bankIdx, customerIndex-1) == -1 && customerIndex != -1 {
			fmt.Println("Customer Number is not available.")
			fmt.Print("Input Customer Number (input -1 for cancel) : ")
			fmt.Scan(&customerIndex)
		}

		if customerIndex != -1 {
			for i := customerIndex - 1; i < worldBank.Banks[bankIdx].nCustomer; i++ {
				worldBank.Banks[bankIdx].customers[i] = worldBank.Banks[bankIdx].customers[i+1]
			}
			worldBank.Banks[bankIdx].nCustomer--

			fmt.Println("Customer Data Deleted Successfully")
			mainMenuAdmin()
			fmt.Println()
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
