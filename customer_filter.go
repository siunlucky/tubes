package main

func searchCustomerByCredentials(bankIdx, accountNumber int, PIN string) int {
	for i := 0; i < worldBank.Banks[bankIdx].nCustomer; i++ {
		if accountNumber == worldBank.Banks[bankIdx].customers[i].accountNumber && PIN == worldBank.Banks[bankIdx].customers[i].PIN {
			return i
		}
	}
	return -1
}

func searchCustomerByName(bankIdx int, search string) int {
	for i := 0; i < worldBank.Banks[bankIdx].nCustomer; i++ {
		if search == worldBank.Banks[bankIdx].customers[i].name {
			return i
		}
	}
	return -1

}

func searchCustomerByIdx(bankIdx, search int) int {
	if search >= 0 && search < worldBank.Banks[bankIdx].nCustomer {
		return search
	}

	return -1
}

func searchCustomerByNIK(bankIdx, search int) int {
	for i := 0; i < worldBank.Banks[bankIdx].nCustomer; i++ {
		if search == worldBank.Banks[bankIdx].customers[i].NIK {
			return i
		}
	}
	return -1
}

func searchCustomerByAccountNumber(bankIdx, search int) int {
	for i := 0; i < worldBank.Banks[bankIdx].nCustomer; i++ {
		if search == worldBank.Banks[bankIdx].customers[i].accountNumber {
			return i
		}
	}
	return -1
}

func searchCustomerByCardNumber(bankIdx, search int) int {
	for i := 0; i < worldBank.Banks[bankIdx].nCustomer; i++ {
		if search == worldBank.Banks[bankIdx].customers[i].cardNumber {
			return i
		}
	}
	return -1
}
