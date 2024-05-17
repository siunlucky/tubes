package main

// Return Idx

func searchBankByUniqueCode(search int) int {
	for i := 0; i < worldBank.nBank; i++ {
		if search == worldBank.Banks[i].uniqueCode {
			return i
		}
	}
	return -1
}

func searchBankByName(search string) int {
	for i := 0; i < worldBank.nBank; i++ {
		if search == worldBank.Banks[i].name {
			return i
		}
	}
	return -1
}
