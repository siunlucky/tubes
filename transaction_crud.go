package main

import "fmt"

func viewSaldo(customer *Customer) {
    fmt.Printf("Balance in your account is: Rp.%d\n", customer.balance)
}

func transfer(customer *Customer) {
	var recipientAccountNumber, amount int
	var recipient Customer
	found := false

	for !found {
		fmt.Print("Input the recipient account number: ")
		fmt.Scan(&recipientAccountNumber)

		if recipientAccountNumber == customer.accountNumber {
			fmt.Println("You cannot transfer to your own account, please input a different account number")
		} else {
			for i := 0; i < customerBank.nCustomer; i++ {
				if customerBank.customers[i].accountNumber == recipientAccountNumber {
					recipient = customerBank.customers[i]
					found = true
					i = customerBank.nCustomer
				}
			}

			if !found {
				fmt.Println("Recipient account number is not found, please try again")
			}
		}
	}

	fmt.Print("Input the amount you want to transfer: Rp.")
	fmt.Scan(&amount)

	if customer.balance < amount {
        fmt.Println("Balance is not enough to transfer")
    } else {
        fee := 0
        if customer.bankCode != recipient.bankCode {
            fee = 5000
        }

        if customer.balance < amount+fee {
            fmt.Println("Balance is not enough to cover the transfer amount and fee")
        } else {
            customer.balance -= (amount + fee)

            for i := 0; i < customerBank.nCustomer; i++ {
                if customerBank.customers[i].accountNumber == recipientAccountNumber {
                    customerBank.customers[i].balance += amount
                    i = customerBank.nCustomer // pengganti break
                }
            }

            fmt.Println("Transfer success")
        }
    }

	// if customer.balance < amount {
	// 	fmt.Println("Balance is not enough to transfer")
	// 	return
	// }

	// fee := 0
	// if customer.bankCode != recipient.bankCode {
	// 	fee = 5000
	// }

	// if customer.balance < amount+fee {
	// 	fmt.Println("Balance is not enough to cover the transfer amount and fee")
	// 	return
	// }

	// customer.balance -= (amount + fee)

	// for i := 0; i < customerBank.nCustomer; i++ {
	// 	if customerBank.customers[i].accountNumber == recipientAccountNumber {
	// 		customerBank.customers[i].balance += amount
	// 		i = customerBank.nCustomer
	// 	}
	// }

	// fmt.Println("Transfer success")
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