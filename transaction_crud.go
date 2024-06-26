package main

import "fmt"

func viewSaldo(customer *Customer) {
	fmt.Printf("Balance in your account is: Rp.%d\n", customer.balance)
}

func transfer(customer *Customer) {
	var recipientAccountNumber, amount, recipientIdx int
	var input int

	for {
		fmt.Println()
		fmt.Println("========================================")
		fmt.Println("=               Transfer               =")
		fmt.Println("1. Account Number")
		fmt.Println("2. Another Bank")
		fmt.Println("3. Back to main menu")
		fmt.Print("Choose the menu option with number: ")
		fmt.Scan(&input)

		if input == 1 {
			var bankIdx int
			fmt.Print("Input the recipient account number: ")
			fmt.Scan(&recipientAccountNumber)
			bankIdx = searchBankByUniqueCode(customer.bankCode)
			recipientIdx = searchCustomerByAccountNumber(bankIdx, recipientAccountNumber)

			for recipientIdx == -1 || recipientAccountNumber == customer.accountNumber {
				fmt.Println("Recipient account number is not found, please try again")
				fmt.Print("Input the recipient account number: ")
				fmt.Scan(&recipientAccountNumber)
				recipientIdx = searchCustomerByAccountNumber(bankIdx, recipientAccountNumber)
			}

			fmt.Print("Input the amount you want to transfer: Rp.")
			fmt.Scan(&amount)

			for amount < 50000 {
				fmt.Println("The minimum amount to transfer is Rp.50000, please input the right amount")
				fmt.Print("Input the amount you want to transfer: Rp.")
				fmt.Scan(&amount)
			}

			if customer.balance < amount {
				fmt.Println("Balance is not enough to transfer")
			} else {
				processTransfer(customer, bankIdx, recipientIdx, amount, false)
				fmt.Println("Transfer success")
			}
		} else if input == 2 {
			var bankIdx, j int
			j = 0
			fmt.Println("========================================")
			fmt.Println("=            Select Bank               =")
			fmt.Println("========================================")
			for i := 0; i < worldBank.nBank; i++ {
				if worldBank.Banks[i].uniqueCode != customer.bankCode {
					fmt.Printf("%d. %s (%d)\n", j+1, worldBank.Banks[i].name, worldBank.Banks[i].uniqueCode)
					j++
				}
			}
			fmt.Print("Choose the bank option with number : ")
			fmt.Scan(&bankIdx)
			for bankIdx < 1 || bankIdx > worldBank.nBank-1 {
				fmt.Println("Invalid bank choice, please choose the right number")
				fmt.Print("Choose the bank option with number : ")
				fmt.Scan(&bankIdx)
			}

			var found bool = true
			for i := 0; i < worldBank.nBank && found; i++ {
				if worldBank.Banks[i].uniqueCode != customer.bankCode {
					bankIdx--
					if bankIdx == 0 {
						bankIdx = i
						found = false
					}
				}
			}

			fmt.Print("Input the recipient account number: ")
			fmt.Scan(&recipientAccountNumber)
			recipientIdx = searchCustomerByAccountNumber(bankIdx, recipientAccountNumber)

			for recipientIdx == -1 {
				fmt.Println("Recipient account number is not found, please try again")
				fmt.Print("Input the recipient account number: ")
				fmt.Scan(&recipientAccountNumber)
				recipientIdx = searchCustomerByAccountNumber(bankIdx, recipientAccountNumber)
			}

			fmt.Print("Input the amount you want to transfer: Rp.")
			fmt.Scan(&amount)

			for amount < 50000 {
				fmt.Println("The minimum amount to transfer is Rp.50000, please input the right amount")
				fmt.Print("Input the amount you want to transfer: Rp.")
				fmt.Scan(&amount)
			}

			if customer.balance < amount+5000 {
				fmt.Println("Balance is not enough to transfer")
			} else {
				processTransfer(customer, bankIdx, recipientIdx, amount, true)
				fmt.Println("Transfer success")
			}

		} else if input == 3 {
			break
		} else {
			fmt.Println("Input is not valid, please input with right option")
		}
	}
}

func processTransfer(customer *Customer, bankIdx int, recipientIdx int, amount int, isDifferentBank bool) {
	recipientBankCode := customer.bankCode
	if isDifferentBank {
		recipientBankCode = worldBank.Banks[bankIdx].uniqueCode
		customer.balance -= 5000
	}
	customer.balance -= amount

	customer.transactions[customer.nTransaction] = Transaction{
		transactionId:          codeGenerator(1000, 100000),
		senderBankCode:         customer.bankCode,
		senderAccountNumber:    customer.accountNumber,
		recipientBankCode:      recipientBankCode,
		recipientAccountNumber: worldBank.Banks[bankIdx].customers[recipientIdx].accountNumber,
		amount:                 amount,
	}
	customer.nTransaction++

	worldBank.Banks[bankIdx].customers[recipientIdx].balance += amount
	worldBank.Banks[bankIdx].customers[recipientIdx].transactions[worldBank.Banks[bankIdx].customers[recipientIdx].nTransaction] = Transaction{
		transactionId:          codeGenerator(1000, 100000),
		senderBankCode:         customer.bankCode,
		senderAccountNumber:    customer.accountNumber,
		recipientBankCode:      recipientBankCode,
		recipientAccountNumber: worldBank.Banks[bankIdx].customers[recipientIdx].accountNumber,
		amount:                 amount,
	}
	worldBank.Banks[bankIdx].customers[recipientIdx].nTransaction++
}

func payment(customer *Customer) {
	fmt.Println("========================================")
	fmt.Println("=               Payments               =")
	fmt.Println("========================================")
	fmt.Println("1. Buy Electricity Token = Rp.100000")
	fmt.Println("2. Buy Water = Rp.50000")
	fmt.Println("3. Buy Internet = Rp.200000")

	var startPayment bool = true
	var input, amount int
	fmt.Print("Choose the payment option with number : ")
	fmt.Scan(&input)

	for startPayment {
		if input == 1 {
			fmt.Println("Input the amount you want to pay : Rp.100000")
			fmt.Scan(&amount)
			payBill(customer, amount, 1)
			startPayment = false
		} else if input == 2 {
			fmt.Println("Input the amount you want to pay : Rp.50000")
			fmt.Scan(&amount)
			payBill(customer, amount, 2)
			startPayment = false
		} else if input == 3 {
			fmt.Println("Input the amount you want to pay : Rp.200000")
			fmt.Scan(&amount)
			payBill(customer, amount, 4)
			startPayment = false
		} else {
			fmt.Println()
			fmt.Println("Input is not valid, please input with right option")
			fmt.Println()
			fmt.Println("1. Buy Electricity Token = Rp.100000")
			fmt.Println("2. Buy Water = Rp.50000")
			fmt.Println("3. Buy Internet = Rp.200000")
			fmt.Scan(&input)
		}
	}
}

func payBill(customer *Customer, payment int, billIndex int) {
	var bill int
	if billIndex == 1 {
		bill = 100000
	} else if billIndex == 2 {
		bill = 50000
	} else {
		bill = 200000
	}

	if payment != bill {
		fmt.Println("The amount to pay must be same as the bill amount, please input the right amount")
		fmt.Println()
	} else if customer.balance < payment {
		fmt.Println("Balance is not enough to pay the bill, please top up your balance first")
		fmt.Println()
	} else {
		customer.balance -= payment
		fmt.Println("The bill has been paid successfully")
		fmt.Printf("Remaining balance: Rp.%d\n", customer.balance)
		fmt.Println()
	}
}

func transactionHistory(customer *Customer) {
	if customer.nTransaction == 0 {
		fmt.Println("No transaction history available")
	} else {
		fmt.Println("========================================")
		fmt.Println("=           Transaction History        =")
		fmt.Println("========================================")
		for i := 0; i < customer.nTransaction; i++ {
			fmt.Printf("Transaction ID: %d\n", customer.transactions[i].transactionId)
			fmt.Printf("Sender Bank Code: %d\n", customer.transactions[i].senderBankCode)
			fmt.Printf("Sender Account Number: %d\n", customer.transactions[i].senderAccountNumber)
			fmt.Printf("Recipient Bank Code: %d\n", customer.transactions[i].recipientBankCode)
			fmt.Printf("Recipient Account Number: %d\n", customer.transactions[i].recipientAccountNumber)
			fmt.Printf("Amount: Rp.%d\n", customer.transactions[i].amount)
			fmt.Println()
		}
	}
}
