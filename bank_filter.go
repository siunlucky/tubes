package main

func searchBankByUniqueCode(search int) int {
	var left, right, mid int

	sortBankByUniqueCode(&worldBank, "selection", "asc")

	left = 0
	right = worldBank.nBank - 1
	for left <= right {
		mid = left + (right-left)/2
		if worldBank.Banks[mid].uniqueCode == search {
			return mid
		} else if worldBank.Banks[mid].uniqueCode < search {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return -1
}