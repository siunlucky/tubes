package main

import "fmt"

func dummyData(n int) {
	for i := 0; i < n; i++ {
		worldBank.Banks[i].name = "Bank" + fmt.Sprint(i+1)
		dummyBankBranch(i, n)
		worldBank.Banks[i].uniqueCode = 100 * (i + 1)
		worldBank.nBank++
		dummyDataAdmin(i)
		dummyDataCustomer(i, n)
	}
}

func dummyBankBranch(bankIdx, n int) {
	for i := 0; i < n; i++ {
		worldBank.Banks[bankIdx].branch[i] = Address{
			district: "Dist" + fmt.Sprint(i+1),
			city:     "City" + fmt.Sprint(i+1),
			province: "Prov" + fmt.Sprint(i+1),
		}
		worldBank.Banks[bankIdx].nBranch++
	}
}

func dummyDataCustomer(bankIdx, n int) {
	for i := 0; i < n; i++ {
		worldBank.Banks[bankIdx].customers[i] = Customer{
			accountNumber: codeGenerator(10000000, 99999999),
			balance:       50000 + (i * 100000),
			cardNumber:    codeGenerator(1000000000000000, 9999999999999999),
			PIN:           "123456",
			NIK:           codeGenerator(1000000000000000, 9999999999999999),
			name:          "Customer" + fmt.Sprint(i+1),
			address: Address{
				district: "Dist" + fmt.Sprint(i+1),
				city:     "City" + fmt.Sprint(i+1),
				province: "Prov" + fmt.Sprint(i+1),
			},
			bankCode: worldBank.Banks[bankIdx].uniqueCode,
		}
		worldBank.Banks[bankIdx].nCustomer++
	}
}

func dummyDataAdmin(bankIdx int) {
	worldBank.Banks[bankIdx].admins[0] = Admin{
		name: "admin",
		Credential: Credential{
			username: "1",
			password: "1",
		},
	}
	worldBank.Banks[bankIdx].nAdmin++
}

func dummyBills() {
	var bill1 Bill
	bill1.description = "Electricity Bill"
	bill1.amount = 120000
	bill1.isPaid = false
	bills[0] = bill1

	var bill2 Bill
	bill2.description = "Water Bill"
	bill2.amount = 50000
	bill2.isPaid = false
	bills[1] = bill2

	nBills = 2
}

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