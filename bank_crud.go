package main

import "fmt"



func insertDataBank(worldBank *WorldBank) {
	var found bool
	var startInsertBank bool = true

	for startInsertBank {
		fmt.Println()
		fmt.Println("=============== Insert Bank Data ===============")
		fmt.Println("Bank", worldBank.nBank+1)

		for worldBank.Banks[worldBank.nBank].name == "" || len(worldBank.Banks[worldBank.nBank].name) < 3 {
			fmt.Print("Input Name : ")
			fmt.Scan(&worldBank.Banks[worldBank.nBank].name)
			if worldBank.Banks[worldBank.nBank].name == "" {
				fmt.Println("Name field cannot be empty, please input again")
			} else if len(worldBank.Banks[worldBank.nBank].name) < 3 {
				fmt.Println("Name should be contain more than 3 char, please input again")
			}
		}

		var startBranchData bool = true
		for startBranchData {
			fmt.Println()
			fmt.Println("Branch", worldBank.Banks[worldBank.nBank].nBranch+1)

			for worldBank.Banks[worldBank.nBank].branch[worldBank.Banks[worldBank.nBank].nBranch].district == "" || len(worldBank.Banks[worldBank.nBank].branch[worldBank.Banks[worldBank.nBank].nBranch].district) < 3 {
				fmt.Print("Input District : ")
				fmt.Scan(&worldBank.Banks[worldBank.nBank].branch[worldBank.Banks[worldBank.nBank].nBranch].district)
				if worldBank.Banks[worldBank.nBank].branch[worldBank.Banks[worldBank.nBank].nBranch].district == "" {
					fmt.Println("District field cannot be empty, please input again")
				} else if len(worldBank.Banks[worldBank.nBank].branch[worldBank.Banks[worldBank.nBank].nBranch].district) < 3 {
					fmt.Println("District should be contain more than 3 char, please input again")
				}
			}

			for worldBank.Banks[worldBank.nBank].branch[worldBank.Banks[worldBank.nBank].nBranch].city == "" || len(worldBank.Banks[worldBank.nBank].branch[worldBank.Banks[worldBank.nBank].nBranch].city) < 3 {
				fmt.Print("City : ")
				fmt.Scan(&worldBank.Banks[worldBank.nBank].branch[worldBank.Banks[worldBank.nBank].nBranch].city)
				if worldBank.Banks[worldBank.nBank].branch[worldBank.Banks[worldBank.nBank].nBranch].city == "" {
					fmt.Println("City field cannot be empty, please input again")
				} else if len(worldBank.Banks[worldBank.nBank].branch[worldBank.Banks[worldBank.nBank].nBranch].city) < 3 {
					fmt.Println("City should be contain more than 3 char, please input again")
				}
			}

			for worldBank.Banks[worldBank.nBank].branch[worldBank.Banks[worldBank.nBank].nBranch].province == "" || len(worldBank.Banks[worldBank.nBank].branch[worldBank.Banks[worldBank.nBank].nBranch].province) < 3 {
				fmt.Print("Province : ")
				fmt.Scan(&worldBank.Banks[worldBank.nBank].branch[worldBank.Banks[worldBank.nBank].nBranch].province)
				if worldBank.Banks[worldBank.nBank].branch[worldBank.Banks[worldBank.nBank].nBranch].province == "" {
					fmt.Println("Province field cannot be empty, please input again")
				} else if len(worldBank.Banks[worldBank.nBank].branch[worldBank.Banks[worldBank.nBank].nBranch].province) < 3 {
					fmt.Println("Province should be contain more than 3 char, please input again")
				}
			}

			worldBank.Banks[worldBank.nBank].nBranch++
			fmt.Println("Branch Data Bank Created Successfully")
			startBranchData = false
			fmt.Println()
		}

		for worldBank.Banks[worldBank.nBank].uniqueCode == 0 || worldBank.Banks[worldBank.nBank].uniqueCode < 100 || worldBank.Banks[worldBank.nBank].uniqueCode > 999 || found {
			found = false

			fmt.Print("Input Unique Code : ")
			fmt.Scan(&worldBank.Banks[worldBank.nBank].uniqueCode)

			for i := 0; i < worldBank.nBank && !found; i++ {
				if worldBank.Banks[i].uniqueCode == worldBank.Banks[worldBank.nBank].uniqueCode {
					found = true
				}
			}

			if worldBank.Banks[worldBank.nBank].uniqueCode == 0 {
				fmt.Println("Unique Code field cannot be empty, please input again")
			} else if worldBank.Banks[worldBank.nBank].uniqueCode < 100 || worldBank.Banks[worldBank.nBank].uniqueCode > 999 {
				fmt.Println("Unique Code should be greater than 99 and less than 1000, please input again")
			} else if found {
				fmt.Println("Unique Code is not available, please input another unique code")
			}
		}	

		dummyDataAdmin(worldBank.nBank)
		
		worldBank.Banks[worldBank.nBank].nCustomer = 0
		worldBank.nBank++

		fmt.Println("Bank Uniqode Data Created Successfully")
		fmt.Println()
		fmt.Println("Bank Data Created Successfully")
		startInsertBank = false
	}
}

func viewDataBank(worldBank WorldBank) {
	var startViewBank bool = true

	for startViewBank {
		if worldBank.nBank == 0 {
			fmt.Println("There is no data to be viewed")
			startViewBank = false
		} else {
			fmt.Println()
			fmt.Println("========================================== Bank Data ===========================================")
			fmt.Printf("%-5s %-20s %-20s %-20s %-20s\n", "No", "Name Bank", "Total Branch", "Unique Code", "Total Customer")
			for i := 0; i < worldBank.nBank; i++ {
				fmt.Printf("%-5d %-30s %-20d %-20d %-20d\n", i+1, worldBank.Banks[i].name, worldBank.Banks[i].nBranch, worldBank.Banks[i].uniqueCode, worldBank.Banks[i].nCustomer)
			}
			fmt.Println()
			fmt.Println("Total Data : ", worldBank.nBank)
			startViewBank = false
		}
	}
}

func editDataBank(worldBank *WorldBank) {
	var bankIndex, choice int
	var found bool

	if worldBank.nBank == 0 {
		fmt.Println("There is no data to be edited")
		mainMenuSuperAdmin()
	}

	for bankIndex != -1 {
		viewDataBank(*worldBank)
		fmt.Println("========== Edit Bank Data ==========")
		fmt.Print("Input Bank Number (input -1 for cancel) : ")
		fmt.Scan(&bankIndex)
		for (bankIndex-1 < 0 || bankIndex-1 > worldBank.nBank || !found) && bankIndex != -1 {
			for i := 0; i < worldBank.nBank && !found; i++ {
				if i == bankIndex-1 {
					found = true
				}
			}
			if bankIndex-1 < 0 || bankIndex-1 > worldBank.nBank-1 {
				fmt.Println("Bank Number is not available")
				fmt.Print("Input Bank Number (input -1 for cancel) : ")
				fmt.Scan(&bankIndex)
			}
		}

		for choice != 3 && bankIndex != -1 {
			fmt.Println("1. Edit Bank Name")
			fmt.Println("2. Edit Unique Code")
			fmt.Println("3. Save")
			fmt.Print("Input Choice : ")
			fmt.Scan(&choice)
			for choice < 1 || choice > 3 {
				fmt.Print("Input Choise : ")
				fmt.Scan(&choice)
			}

			switch choice {
			case 1:
				fmt.Println("Old Name : ", worldBank.Banks[bankIndex-1].name)
				fmt.Print("Input New Name : ")
				fmt.Scan(&worldBank.Banks[bankIndex-1].name)
				for worldBank.Banks[bankIndex-1].name == "" || len(worldBank.Banks[bankIndex-1].name) < 3 {
					if worldBank.Banks[bankIndex-1].name == "" {
						fmt.Println("Name field cannot be empty, please input again")
					} else if len(worldBank.Banks[bankIndex-1].name) < 3 {
						fmt.Println("Name should be contain more than 3 char, please input again")
					}
					fmt.Print("Input New Name : ")
					fmt.Scan(&worldBank.Banks[bankIndex-1].name)
				}

				fmt.Print("Bank Name Edited Successfully\n")
				viewDataBank(*worldBank)
				fmt.Println()

			case 2:
				fmt.Println("Old Unique Code : ", worldBank.Banks[bankIndex-1].uniqueCode)
				fmt.Print("Input New Unique Code : ")
				fmt.Scan(&worldBank.Banks[bankIndex-1].uniqueCode)
				for worldBank.Banks[bankIndex-1].uniqueCode == 0 || worldBank.Banks[bankIndex-1].uniqueCode < 100 || worldBank.Banks[bankIndex-1].uniqueCode > 999 || found {
					found = false

					for i := 0; i < worldBank.nBank && !found; i++ {
						if worldBank.Banks[i].uniqueCode == worldBank.Banks[worldBank.nBank].uniqueCode {
							found = true
						}
					}

					if worldBank.Banks[bankIndex-1].uniqueCode == 0 {
						fmt.Println("Unique Code field cannot be empty, please input again")
						fmt.Print("New Unique Code : ")
						fmt.Scan(&worldBank.Banks[bankIndex-1].uniqueCode)
					} else if worldBank.Banks[bankIndex-1].uniqueCode < 100 || worldBank.Banks[bankIndex-1].uniqueCode > 999 {
						fmt.Println("Unique Code should be greater than 99 and less than 1000, please input again")
						fmt.Print("New Unique Code : ")
						fmt.Scan(&worldBank.Banks[bankIndex-1].uniqueCode)
					} else if found {
						fmt.Println("Unique Code is not available, please input another unique code")
						fmt.Print("New Unique Code : ")
						fmt.Scan(&worldBank.Banks[bankIndex-1].uniqueCode)
					}
				}

				fmt.Print("Unique Code Edited Successfully\n")
				viewDataBank(*worldBank)
				fmt.Println()

			case 3:
				fmt.Println("Bank Data Edited Successfully")
				viewDataBank(*worldBank)
				bankIndex = -1
				fmt.Println()
			}
		}
	}
}

func deleteDataBank(worldBank *WorldBank) {
	var bankIndex int
	var found bool

	if worldBank.nBank == 0 {
		fmt.Println("There is no data to be deleted")
		mainMenuSuperAdmin()
	}

	for bankIndex != -1 {
		viewDataBank(*worldBank)

		fmt.Print("Input Bank Number (input -1 for cancel) :")
		fmt.Scan(&bankIndex)
		for (bankIndex-1 < 0 || bankIndex-1 > worldBank.nBank || !found) && bankIndex != -1 {
			for i := 0; i < worldBank.nBank && !found; i++ {
				if i == bankIndex-1 {
					found = true
				}
			}
			if bankIndex-1 < 0 || bankIndex-1 > worldBank.nBank-1 {
				fmt.Println("Bank Number is not available.")
				fmt.Print("Input Bank Number (input -1 for cancel) : ")
				fmt.Scan(&bankIndex)
			}
		}

		if bankIndex != -1 {
			for i := bankIndex - 1; i < worldBank.nBank; i++ {
				worldBank.Banks[i] = worldBank.Banks[i+1]
			}
			worldBank.nBank--

			fmt.Println("Bank Data Deleted Successfully")
			fmt.Println()
		}
	}
}