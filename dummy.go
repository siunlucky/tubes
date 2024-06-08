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
