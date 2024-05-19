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

func payment(customer *Customer) {
	var choice int

	fmt.Println("========================================")
	fmt.Println("=               Payments               =")
	fmt.Println("1. Pay Electricity Bill")
	fmt.Println("2. Pay HUTANG")
	fmt.Print("Choose the payment option with number : ")
	fmt.Scan(&choice)

	switch choice {
	case 1:
		payBill(customer, 0)
	case 2:
		payBill(customer, 1)
	default:
		fmt.Println("Input is not valid, please input with right option")
	}
}

func payBill(customer *Customer, billIndex int) {
	if billIndex < 0 || billIndex >= len(bills) {
		fmt.Println("Invalid bill index.")
		return
	}

	bill := &bills[billIndex]
	if bill.isPaid {
		fmt.Printf("The %s has already been paid.\n", bill.description)
		return
	}

	fmt.Printf("The %s is Rp.%d\n", bill.description, bill.amount)
	fmt.Print("Enter the amount to pay: Rp.")
	var amount int
	fmt.Scan(&amount)

	if amount != bill.amount {
		fmt.Printf("You need to pay the exact amount of Rp.%d\n", bill.amount)
		return
	}

	if customer.balance < amount {
		fmt.Println("Balance is not enough to pay the bill")
		return
	}

	customer.balance -= amount
	bill.isPaid = true
	fmt.Printf("The %s has been paid successfully.\n", bill.description)
}