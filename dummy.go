package main

// func dummyData() {
// 	worldBank.Banks[0].name = "BankA"
// 	worldBank.Banks[0].uniqueCode = 123
// 	worldBank.Banks[0].nBranch = 1
// 	worldBank.Banks[0].branch[0] = Address{district: "Dist1", city: "City1", province: "Prov1"}
// 	worldBank.Banks[0].branch[1] = Address{district: "Dist2", city: "City2", province: "Prov2"}
// 	dummyDataAdmin(0)
	
// 	worldBank.Banks[1].name = "BankB"
// 	worldBank.Banks[1].uniqueCode = 221
// 	worldBank.Banks[1].nBranch = 2
// 	worldBank.Banks[1].branch[0] = Address{district: "Dist3", city: "City3", province: "Prov3"}
// 	worldBank.Banks[1].branch[1] = Address{district: "Dist4", city: "City4", province: "Prov4"}
// 	dummyDataAdmin(1)
	
// 	worldBank.nBank = 2

// 	var dummyCustomers [5]Customer

// 	dummyCustomers[0] = Customer{
// 		accountNumber: 1001, balance: 0, PIN: "123456", NIK: 1234567890123451, name: "Dummy1",
// 		address: Address{district: "Dist1", city: "City1", province: "Prov1"}, bankCode: 123,
// 		branchAddress: Address{district: "Dist1", city: "City1", province: "Prov1"},
// 	}
// 	dummyCustomers[1] = Customer{
// 		accountNumber: 1002, balance: 0, PIN: "123456", NIK: 1234567890123452, name: "Dummy2",
// 		address: Address{district: "Dist2", city: "City2", province: "Prov2"}, bankCode: 456,
// 		branchAddress: Address{district: "Dist2", city: "City2", province: "Prov2"},
// 	}
// 	dummyCustomers[2] = Customer{
// 		accountNumber: 1003, balance: 0, PIN: "123456", NIK: 1234567890123453, name: "Dummy3",
// 		address: Address{district: "Dist3", city: "City3", province: "Prov3"}, bankCode: 221,
// 		branchAddress: Address{district: "Dist3", city: "City3", province: "Prov3"},
// 	}
// 	dummyCustomers[3] = Customer{
// 		accountNumber: 1004, balance: 0, PIN: "123456", NIK: 1234567890123454, name: "Dummy4",
// 		address: Address{district: "Dist4", city: "City4", province: "Prov4"}, bankCode: 111,
// 		branchAddress: Address{district: "Dist4", city: "City4", province: "Prov4"},
// 	}
// 	dummyCustomers[4] = Customer{
// 		accountNumber: 1005, balance: 0, PIN: "123456", NIK: 1234567890123455, name: "Dummy5",
// 		address: Address{district: "Dist5", city: "City5", province: "Prov5"}, bankCode: 222,
// 		branchAddress: Address{district: "Dist5", city: "City5", province: "Prov5"},
// 	}

// 	for i := 0; i < len(dummyCustomers); i++ {
// 		addCustomerToBank(dummyCustomers[i])
// 	}
// }

// func dummyDataAdmin(idx int) {
// 	dummyAdmins := []Admin{
// 		{
// 			name: "admin",
// 			Credential: Credential{
// 				username: "1",
// 				password: "1",
// 			},
// 		},
// 	}

// 	for i := 0; i < len(dummyAdmins); i++ {
// 		if worldBank.Banks[idx].nAdmin < NMAX_Admin {
// 			worldBank.Banks[idx].admins[worldBank.Banks[idx].nAdmin] = dummyAdmins[i]
// 			worldBank.Banks[idx].nAdmin++
// 		} else {
// 			fmt.Println("Admin bank is full, unable to add more admins")
// 			i = 2000

// 			fmt.Println("Back to Main Menu")
// 			loadingAuth("Loading")
// 			fmt.Println()
// 			menu()
// 		}
// 	}
// }

// func addCustomerToBank(customer Customer) {
// 	for i := 0; i < worldBank.nBank; i++ {
// 		if worldBank.Banks[i].uniqueCode == customer.bankCode {
// 			if worldBank.Banks[i].nCustomer < NMAX {
// 				worldBank.Banks[i].customers[worldBank.Banks[i].nCustomer] = customer
// 				worldBank.Banks[i].nCustomer++
// 			} else {
// 				fmt.Println("Customer bank is full, unable to add more customers")
// 				fmt.Println("Back to Main Menu")
// 				loadingAuth("Loading")
// 				fmt.Println()
// 				menu()
// 			}
// 			return
// 		}
// 	}
// }

func dummyBills() {
    var bill1 Bill
    bill1.description = "Electricity Bill"
    bill1.amount = 120000
    bill1.isPaid = false
    bills[0] = bill1

    var bill2 Bill
    bill2.description = "Hutang Bill"
    bill2.amount = 500000
    bill2.isPaid = false
    bills[1] = bill2

    nBills = 2
}