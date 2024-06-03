package main

func searchBankByUniqueCode(search int) int {
	for i := 0; i < worldBank.nBank; i++ {
		if search == worldBank.Banks[i].uniqueCode {
			return i
		}
	}
	return -1
}