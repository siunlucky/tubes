package main

import "fmt"

func viewSaldo(customer *Customer) {
	fmt.Printf("Saldo Anda adalah: %d\n", customer.balance)
}

func transfer(customer *Customer) {
	var recipientAccountNumber, amount int
	fmt.Print("Masukkan nomor rekening tujuan: ")
	fmt.Scan(&recipientAccountNumber)
	fmt.Print("Masukkan jumlah transfer: ")
	fmt.Scan(&amount)

	for i := 0; i < customerBank.nCustomer; i++ {
		if customerBank.customers[i].accountNumber == recipientAccountNumber {
			if customer.balance >= amount {
				customer.balance -= amount
				customerBank.customers[i].balance += amount
				fmt.Println("Transfer berhasil.")
				return
			} else {
				fmt.Println("Saldo tidak mencukupi.")
				return
			}
		}
	}

	fmt.Println("Nomor rekening tujuan tidak ditemukan.")
}
