package main

import "fmt"

func viewSaldo(customer *Customer) {
    fmt.Printf("Balance in your account is: %d\n", customer.balance)
}

func transfer(customer *Customer) {
    var recipientAccountNumber, amount int

    fmt.Print("Input the recipient account number: ")
    fmt.Scan(&recipientAccountNumber)

    fmt.Print("Input the amount you want to transfer: ")
    fmt.Scan(&amount)

    var recipient Customer
    found := false
    for i := 0; i < customerBank.nCustomer; i++ {
        if customerBank.customers[i].accountNumber == recipientAccountNumber {
            recipient = customerBank.customers[i]
            found = true
        }
    }

    if !found {
        fmt.Println("Recipient account number is not found")
        return
    }

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
            i = 1000000
        }
    }

    fmt.Println("Transfer success")
}