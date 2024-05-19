package main

import "fmt"

func viewSaldo(customer *Customer) {
    fmt.Printf("Balance in your account is : Rp.%d\n", customer.balance)
}

func transfer(customer *Customer) {
	var recipientAccountNumber, amount int
	var recipient Customer
	found := false

	for !found {
		fmt.Print("Input the recipient account number : ")
		fmt.Scan(&recipientAccountNumber)

		if recipientAccountNumber == customer.accountNumber {
			fmt.Println("You cannot transfer to your own account, Please input a different account number")
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

	fmt.Print("Input the amount you want to transfer : Rp.")
	fmt.Scan(&amount)

	if customer.balance < amount {
		fmt.Println("Balance is not enough to transfer")
		return
	}

	fee := 0
	if customer.bankCode != recipient.bankCode {
		fee = 5000
	}

	if customer.balance < amount+fee {
		fmt.Println("Balance is not enough to cover the transfer amount and fee")
		return
	}

	customer.balance -= (amount + fee)

	for i := 0; i < customerBank.nCustomer; i++ {
		if customerBank.customers[i].accountNumber == recipientAccountNumber {
			customerBank.customers[i].balance += amount
			i = customerBank.nCustomer
		}
	}

	fmt.Println("Transfer success")
}