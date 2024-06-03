package main

func sortBankByUniqueCode(worldBank *WorldBank, method string, direction string) {
	// Selection Sort & insertion sort
	if method == "selection" {
		if direction == "asc" {
			for i := 0; i < worldBank.nBank-1; i++ {
				minIndex := i
				for j := i + 1; j < worldBank.nBank; j++ {
					if worldBank.Banks[j].uniqueCode < worldBank.Banks[minIndex].uniqueCode {
						minIndex = j
					}
				}
				worldBank.Banks[i], worldBank.Banks[minIndex] = worldBank.Banks[minIndex], worldBank.Banks[i]
			}
		} else if direction == "desc" {
			for i := 0; i < worldBank.nBank-1; i++ {
				maxIndex := i
				for j := i + 1; j < worldBank.nBank; j++ {
					if worldBank.Banks[j].uniqueCode > worldBank.Banks[maxIndex].uniqueCode {
						maxIndex = j
					}
				}
				worldBank.Banks[i], worldBank.Banks[maxIndex] = worldBank.Banks[maxIndex], worldBank.Banks[i]
			}
		}
	} else if method == "insertion" {
		if direction == "asc" {
			for i := 1; i < worldBank.nBank; i++ {
				key := worldBank.Banks[i]
				j := i - 1
				for j >= 0 && worldBank.Banks[j].uniqueCode > key.uniqueCode {
					worldBank.Banks[j+1] = worldBank.Banks[j]
					j = j - 1
				}
				worldBank.Banks[j+1] = key
			}
		} else if direction == "desc" {
			for i := 1; i < worldBank.nBank; i++ {
				key := worldBank.Banks[i]
				j := i - 1
				for j >= 0 && worldBank.Banks[j].uniqueCode < key.uniqueCode {
					worldBank.Banks[j+1] = worldBank.Banks[j]
					j = j - 1
				}
				worldBank.Banks[j+1] = key
			}
		}
	} else {
		println("Invalid method")
	}
}
