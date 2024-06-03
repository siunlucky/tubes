package main

func sortBankByUniqueCode(worldBank *WorldBank, direction string) {
	var i, j int
	for i = 0; i < worldBank.nBank; i++ {
		for j = i + 1; j < worldBank.nBank; j++ {
			if direction == "asc" {
				if worldBank.Banks[i].uniqueCode > worldBank.Banks[j].uniqueCode {
					temp := worldBank.Banks[i]
					worldBank.Banks[i] = worldBank.Banks[j]
					worldBank.Banks[j] = temp
				}
			} else {
				if worldBank.Banks[i].uniqueCode < worldBank.Banks[j].uniqueCode {
					temp := worldBank.Banks[i]
					worldBank.Banks[i] = worldBank.Banks[j]
					worldBank.Banks[j] = temp
				}
			}
		}
	}
}
