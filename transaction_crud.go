package main

import "fmt"

func viewSaldo(customer *Customer) {
	fmt.Printf("Balance in your account is: Rp.%d\n", customer.balance)
}

var transactionCounter int

func getNewTransactionId() int {
	transactionCounter++
	return transactionCounter
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
			for i := 0; i < worldBank.nBank; i++ {
				if worldBank.Banks[i].uniqueCode != customer.bankCode {
					fmt.Printf("%d. %s (%d)\n", j+1, worldBank.Banks[i].name, worldBank.Banks[i].uniqueCode)
					j++
				}
			}
			fmt.Print("Choose the bank option with number: ")
			fmt.Scan(&bankIdx)
			for bankIdx < 1 || bankIdx > worldBank.nBank-1 {
				fmt.Println("Invalid bank choice, please choose the right number")
				fmt.Print("Choose the bank option with number : ")
				fmt.Scan(&bankIdx)
			}

			// Adjust bank index for proper selection
			for i := 0; i < worldBank.nBank; i++ {
				if worldBank.Banks[i].uniqueCode != customer.bankCode {
					bankIdx--
					if bankIdx == 0 {
						bankIdx = i
						break
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
		customer.balance -= 5000 // Transfer fee
	}
	customer.balance -= amount

	customer.transactions[customer.nTransaction] = Transaction{
		transactionId:          getNewTransactionId(),
		senderBankCode:         customer.bankCode,
		senderAccountNumber:    customer.accountNumber,
		recipientBankCode:      recipientBankCode,
		recipientAccountNumber: worldBank.Banks[bankIdx].customers[recipientIdx].accountNumber,
		amount:                 amount,
	}
	customer.nTransaction++

	worldBank.Banks[bankIdx].customers[recipientIdx].balance += amount
	worldBank.Banks[bankIdx].customers[recipientIdx].transactions[worldBank.Banks[bankIdx].customers[recipientIdx].nTransaction] = Transaction{
		transactionId:          getNewTransactionId(),
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

	hasUnpaidBills := false
	for i := 0; i < nBills; i++ {
		if !bills[i].isPaid {
			fmt.Printf("%d. %s - Rp.%d\n", i+1, bills[i].description, bills[i].amount)
			hasUnpaidBills = true
		}
	}

	if !hasUnpaidBills {
		fmt.Println("No unpaid bills available")
	} else {
		fmt.Print("Choose the bill you want to pay : ")
		var billChoice int
		fmt.Scan(&billChoice)

		if billChoice < 1 || billChoice > nBills || bills[billChoice-1].isPaid {
			fmt.Println("Invalid choice, please choose the right number")
		} else {
			payBill(customer, billChoice-1)
		}
	}
}

func payBill(customer *Customer, billIndex int) {
	bill := &bills[billIndex]
	if bill.isPaid {
		fmt.Printf("The %s has already been paid\n", bill.description)
		fmt.Println()
	} else {
		fmt.Printf("The %s is Rp.%d\n", bill.description, bill.amount)
		var amount int
		valid := false

		for !valid {
			fmt.Print("Enter the amount to pay : Rp.")
			fmt.Scan(&amount)

			if amount != bill.amount {
				fmt.Println("The amount to pay must be same as the bill amount, please input the right amount")
				fmt.Println()
			} else if customer.balance < amount {
				fmt.Println("Balance is not enough to pay the bill")
				fmt.Println()
				valid = true
			} else {
				customer.balance -= amount
				bill.isPaid = true
				fmt.Printf("The %s has been paid successfully\n", bill.description)
				fmt.Printf("Remaining balance: Rp.%d\n", customer.balance)
				valid = true
			}
		}
	}
}

func transactionHistory(customer *Customer) {

	if customer.nTransaction == 0 {
		fmt.Println("No transaction history available")
	} else {
		fmt.Println("========================================")
		fmt.Println("=           Transaction History        =")
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
