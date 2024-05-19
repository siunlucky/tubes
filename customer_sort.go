package main

func searchCustomerByNIK(bankIdx int, search int) int {
	for i := 0; i < worldBank.Banks[bankIdx].nCustomer; i++ {
		if search == worldBank.Banks[bankIdx].customers[i].NIK {
			return i
		}
	}
	return -1
}

func searchCustomerByAccountNumber(bankIdx int, search int) int {
	for i := 0; i < worldBank.Banks[bankIdx].nCustomer; i++ {
		if search == worldBank.Banks[bankIdx].customers[i].accountNumber {
			return i
		}
	}
	return -1
}

func searchCustomerByCardNumber(bankIdx int, search int) int {
	for i := 0; i < worldBank.Banks[bankIdx].nCustomer; i++ {
		if search == worldBank.Banks[bankIdx].customers[i].cardNumber {
			return i
		}
	}
	return -1
}
