package main

import "fmt"

// insertDataBank function to insert the bank data
func insertDataBank(worldBank *WorldBank) {
	// var retry string = "y"
	var found bool

	var startInsertBank bool = true

	for startInsertBank {
		fmt.Println()
		fmt.Println("Bank", worldBank.nBank+1)

		for worldBank.Banks[worldBank.nBank].name == "" || len(worldBank.Banks[worldBank.nBank].name) < 3 {
			fmt.Print("Name : ")
			fmt.Scan(&worldBank.Banks[worldBank.nBank].name)
			if worldBank.Banks[worldBank.nBank].name == "" {
				fmt.Println("Name field may not be empty.")
			} else if len(worldBank.Banks[worldBank.nBank].name) < 3 {
				fmt.Println("Name should be contain more than 3 char.")
			}
		}

		var bankJalan bool = true
		// var retryBranch string
		for bankJalan {
			fmt.Println()
			fmt.Println("Branch", worldBank.Banks[worldBank.nBank].nBranch+1)

			for worldBank.Banks[worldBank.nBank].branch[worldBank.Banks[worldBank.nBank].nBranch].district == "" || len(worldBank.Banks[worldBank.nBank].branch[worldBank.Banks[worldBank.nBank].nBranch].district) < 3 {
				fmt.Print("District : ")
				fmt.Scan(&worldBank.Banks[worldBank.nBank].branch[worldBank.Banks[worldBank.nBank].nBranch].district)
				if worldBank.Banks[worldBank.nBank].branch[worldBank.Banks[worldBank.nBank].nBranch].district == "" {
					fmt.Println("District field may not be empty.")
				} else if len(worldBank.Banks[worldBank.nBank].branch[worldBank.Banks[worldBank.nBank].nBranch].district) < 3 {
					fmt.Println("District should be contain more than 3 char.")
				}
			}

			for worldBank.Banks[worldBank.nBank].branch[worldBank.Banks[worldBank.nBank].nBranch].city == "" || len(worldBank.Banks[worldBank.nBank].branch[worldBank.Banks[worldBank.nBank].nBranch].city) < 3 {
				fmt.Print("City : ")
				fmt.Scan(&worldBank.Banks[worldBank.nBank].branch[worldBank.Banks[worldBank.nBank].nBranch].city)
				if worldBank.Banks[worldBank.nBank].branch[worldBank.Banks[worldBank.nBank].nBranch].city == "" {
					fmt.Println("City field may not be empty.")
				} else if len(worldBank.Banks[worldBank.nBank].branch[worldBank.Banks[worldBank.nBank].nBranch].city) < 3 {
					fmt.Println("City should be contain more than 3 char.")
				}
			}

			for worldBank.Banks[worldBank.nBank].branch[worldBank.Banks[worldBank.nBank].nBranch].province == "" || len(worldBank.Banks[worldBank.nBank].branch[worldBank.Banks[worldBank.nBank].nBranch].province) < 3 {
				fmt.Print("Province : ")
				fmt.Scan(&worldBank.Banks[worldBank.nBank].branch[worldBank.Banks[worldBank.nBank].nBranch].province)
				if worldBank.Banks[worldBank.nBank].branch[worldBank.Banks[worldBank.nBank].nBranch].province == "" {
					fmt.Println("Province field may not be empty.")
				} else if len(worldBank.Banks[worldBank.nBank].branch[worldBank.Banks[worldBank.nBank].nBranch].province) < 3 {
					fmt.Println("Province should be contain more than 3 char.")
				}
			}

			worldBank.Banks[worldBank.nBank].nBranch++
			fmt.Println("Branch Data Created Successfully")

			bankJalan = false

			// fmt.Println("\nWill You Create Branch Again? (Y/N)")
			// fmt.Scan(&retryBranch)
			// for retryBranch != "Y" && retryBranch != "N" && retryBranch != "y" && retryBranch != "n" {
			// 	fmt.Scan(&retryBranch)
			// }
		}

		for worldBank.Banks[worldBank.nBank].uniqueCode == 0 || worldBank.Banks[worldBank.nBank].uniqueCode < 100 || worldBank.Banks[worldBank.nBank].uniqueCode > 999 || found {
			found = false

			fmt.Print("Unique Code : ")
			fmt.Scan(&worldBank.Banks[worldBank.nBank].uniqueCode)

			for i := 0; i < worldBank.nBank && !found; i++ {
				if worldBank.Banks[i].uniqueCode == worldBank.Banks[worldBank.nBank].uniqueCode {
					found = true
				}
			}
			if worldBank.Banks[worldBank.nBank].uniqueCode == 0 {
				fmt.Println("Unique Code field may not be empty.")
			} else if worldBank.Banks[worldBank.nBank].uniqueCode < 100 || worldBank.Banks[worldBank.nBank].uniqueCode > 999 {
				fmt.Println("Unique Code should be greater than 99 and less than 1000")
			} else if found {
				fmt.Println("Unique Code is not available.")
			}
		}

		// dummyDataAdmin(worldBank.nBank)
		worldBank.Banks[worldBank.nBank].nCustomer = 0
		worldBank.nBank++

		fmt.Println("Bank Data Created Successfully")

		startInsertBank = false

		// fmt.Println("\nWill Create Bank Data Again? (Y/N)")
		// fmt.Scan(&retry)
		// // var startRetry bool = true
		// for {
		// 	if retry == "N" || retry == "n" {
		// 		break
		// 		mainMenuSuperAdmin()
		// 	} else if retry == "Y" || retry == "y" {
		// 		startRetry = false
		// 		startInsertBank = true
		// 	} else {
		// 		fmt.Println("Please input the right option")
		// 		fmt.Scan(&retry)
		// 	}
		// }
	}
}

// viewDataBank function to view the bank data
func viewDataBank(worldBank WorldBank) {
	var startViewBank bool = true

	for startViewBank {
		if worldBank.nBank == 0 {
			fmt.Println("There is no data to be viewed.")
			startViewBank = false
		} else {
			fmt.Println()
			fmt.Println("Bank Data :")
			fmt.Println()
			fmt.Printf("%-5s %-20s %-20s %-20s %-20s\n", "No", "Name", "Total Branch", "Unique Code", "Total Customer")
			for i := 0; i < worldBank.nBank; i++ {
				fmt.Printf("%-5d %-20s %-20d %-20d %-20d\n", i+1, worldBank.Banks[i].name, worldBank.Banks[i].nBranch, worldBank.Banks[i].uniqueCode, worldBank.Banks[i].nCustomer)
			}
			fmt.Println()
			fmt.Println("Total Data : ", worldBank.nBank)
			startViewBank = false
		}
	}
}

// editDataBank function to edit the bank data
func editDataBank(worldBank *WorldBank) {
	var bankIndex, choice int
	var found bool

	if worldBank.nBank == 0 {
		fmt.Println("There is no data to be edited.")
		mainMenuSuperAdmin()
	}

	for bankIndex != -1 {
		viewDataBank(*worldBank)

		fmt.Print("Input Bank Number (input -1 for cancel) : ")
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

		for choice != 4 && bankIndex != -1 {
			fmt.Println("1. Edit Bank Name")
			fmt.Println("2. Edit Branch Data")
			fmt.Println("3. Edit Unique Code")
			fmt.Println("4. Save")
			fmt.Print("Input : ")
			fmt.Scan(&choice)
			for choice < 1 || choice > 4 {
				fmt.Print("Input : ")
				fmt.Scan(&choice)
			}

			switch choice {
			case 1:
				fmt.Println("Old Name :", worldBank.Banks[bankIndex-1].name)
				fmt.Print("New Name : ")
				fmt.Scan(&worldBank.Banks[bankIndex-1].name)
				for worldBank.Banks[bankIndex-1].name == "" || len(worldBank.Banks[bankIndex-1].name) < 3 {
					if worldBank.Banks[bankIndex-1].name == "" {
						fmt.Println("Name field may not be empty.")
					} else if len(worldBank.Banks[bankIndex-1].name) < 3 {
						fmt.Println("Name should be contain more than 3 char.")
					}
					fmt.Print("New Name : ")
					fmt.Scan(&worldBank.Banks[bankIndex-1].name)
				}
				fmt.Print("Bank Name Edited Successfully\n")
				viewDataBank(*worldBank)
				fmt.Println()

			case 2:
				var branchIndex, branchChoice int
				var foundBranch bool

				if worldBank.Banks[bankIndex-1].nBranch == 0 {
					fmt.Println("There is no branch data to be edited.")
				} else {
					for branchIndex != -1 {
						fmt.Println()
						fmt.Println("Branch Data :")
						fmt.Printf("%-5s %-20s %-20s %-20s\n", "No", "District", "City", "Province")
						for i := 0; i < worldBank.Banks[bankIndex-1].nBranch; i++ {
							fmt.Printf("%-5d %-20s %-20s %-20s\n", i+1, worldBank.Banks[bankIndex-1].branch[i].district, worldBank.Banks[bankIndex-1].branch[i].city, worldBank.Banks[bankIndex-1].branch[i].province)
						}
						fmt.Println("Total Data :", worldBank.Banks[bankIndex-1].nBranch)
						fmt.Print("Input Branch Number (input -1 for cancel) : ")
						fmt.Scan(&branchIndex)
						for (branchIndex-1 < 0 || branchIndex-1 > worldBank.Banks[bankIndex-1].nBranch || !foundBranch) && branchIndex != -1 {
							for i := 0; i < worldBank.Banks[bankIndex-1].nBranch && !foundBranch; i++ {
								if i == branchIndex-1 {
									foundBranch = true
								}
							}
							if branchIndex-1 < 0 || branchIndex-1 > worldBank.Banks[bankIndex-1].nBranch {
								fmt.Println("Branch Number is not available.")
								fmt.Print("Input Branch Number (input -1 for cancel) : ")

								fmt.Scan(&branchIndex)
							}
						}

						for branchChoice != 3 && branchIndex != -1 && branchIndex != 4 {
							fmt.Println("1. Edit District")
							fmt.Println("2. Edit City")
							fmt.Println("3. Edit Province")
							fmt.Println("4. Save")
							fmt.Print("Input : ")
							fmt.Scan(&branchChoice)
							for branchChoice < 1 || branchChoice > 4 {
								fmt.Print("Input : ")
								fmt.Scan(&branchChoice)
							}

							switch branchChoice {
							case 1:
								fmt.Println("Old District :", worldBank.Banks[bankIndex-1].branch[branchIndex-1].district)
								fmt.Print("New District : ")
								fmt.Scan(&worldBank.Banks[bankIndex-1].branch[branchIndex-1].district)
								for worldBank.Banks[bankIndex-1].branch[branchIndex-1].district == "" || len(worldBank.Banks[bankIndex-1].branch[branchIndex-1].district) < 3 {
									if worldBank.Banks[bankIndex-1].branch[branchIndex-1].district == "" {
										fmt.Println("District field may not be empty.")
									} else if len(worldBank.Banks[bankIndex-1].branch[branchIndex-1].district) < 3 {
										fmt.Println("District should be contain more than 3 char.")
									}
									fmt.Print("New District : ")
									fmt.Scan(&worldBank.Banks[bankIndex-1].branch[branchIndex-1].district)
								}
								fmt.Print("District Edited Successfully\n")
								viewDataBank(*worldBank)
								fmt.Println()

							case 2:
								fmt.Println("Old City :", worldBank.Banks[bankIndex-1].branch[branchIndex-1].city)
								fmt.Print("New City : ")
								fmt.Scan(&worldBank.Banks[bankIndex-1].branch[branchIndex-1].city)
								for worldBank.Banks[bankIndex-1].branch[branchIndex-1].city == "" || len(worldBank.Banks[bankIndex-1].branch[branchIndex-1].city) < 3 {
									if worldBank.Banks[bankIndex-1].branch[branchIndex-1].city == "" {
										fmt.Println("City field may not be empty.")
									} else if len(worldBank.Banks[bankIndex-1].branch[branchIndex-1].city) < 3 {
										fmt.Println("City should be contain more than 3 char.")
									}
									fmt.Print("New City : ")
									fmt.Scan(&worldBank.Banks[bankIndex-1].branch[branchIndex-1].city)
								}
								fmt.Print("City Edited Successfully\n")
								viewDataBank(*worldBank)
								fmt.Println()

							case 3:
								fmt.Println("Old Province :", worldBank.Banks[bankIndex-1].branch[branchIndex-1].province)
								fmt.Print("New Province : ")
								fmt.Scan(&worldBank.Banks[bankIndex-1].branch[branchIndex-1].province)
								for worldBank.Banks[bankIndex-1].branch[branchIndex-1].province == "" || len(worldBank.Banks[bankIndex-1].branch[branchIndex-1].province) < 3 {
									if worldBank.Banks[bankIndex-1].branch[branchIndex-1].province == "" {
										fmt.Println("Province field may not be empty.")
									} else if len(worldBank.Banks[bankIndex-1].branch[branchIndex-1].province) < 3 {
										fmt.Println("Province should be contain more than 3 char.")
									}
									fmt.Print("New Province : ")
									fmt.Scan(&worldBank.Banks[bankIndex-1].branch[branchIndex-1].province)
								}
								fmt.Print("Province Edited Successfully\n")
								viewDataBank(*worldBank)
								fmt.Println()
							}
						}
					}
				}

			case 3:
				fmt.Println("Old Unique Code :", worldBank.Banks[bankIndex-1].uniqueCode)
				fmt.Print("New Unique Code : ")
				fmt.Scan(&worldBank.Banks[bankIndex-1].uniqueCode)
				for worldBank.Banks[bankIndex-1].uniqueCode == 0 || worldBank.Banks[bankIndex-1].uniqueCode < 100 || worldBank.Banks[bankIndex-1].uniqueCode > 999 || found {
					found = false

					for i := 0; i < worldBank.nBank && !found; i++ {
						if worldBank.Banks[i].uniqueCode == worldBank.Banks[worldBank.nBank].uniqueCode {
							found = true
						}
					}

					if worldBank.Banks[bankIndex-1].uniqueCode == 0 {
						fmt.Println("Unique Code field may not be empty.")
						fmt.Print("New Unique Code : ")
						fmt.Scan(&worldBank.Banks[bankIndex-1].uniqueCode)
					} else if worldBank.Banks[bankIndex-1].uniqueCode < 100 || worldBank.Banks[bankIndex-1].uniqueCode > 999 {
						fmt.Println("Unique Code should be greater than 99 and less than 1000")
						fmt.Print("New Unique Code : ")
						fmt.Scan(&worldBank.Banks[bankIndex-1].uniqueCode)
					} else if found {
						fmt.Println("Unique Code is not available.")
						fmt.Print("New Unique Code : ")
						fmt.Scan(&worldBank.Banks[bankIndex-1].uniqueCode)
					}
				}
				fmt.Print("Unique Code Edited Successfully\n")
				viewDataBank(*worldBank)
				fmt.Println()
			}
		}
	}
}

// deleteDataBank function to delete the bank data
func deleteDataBank(worldBank *WorldBank) {
	var bankIndex int
	var found bool

	if worldBank.nBank == 0 {
		fmt.Println("There is no data to be deleted.")
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
