package main

import "fmt"

// register function to register as customer
func register() {
	if customerBank.nCustomer >= NMAX {
		fmt.Println("Customer bank is full, unable to register new customer.")
		return
	}

	var customer Customer
	fmt.Print("Masukkan nomor rekening: ")
	fmt.Scan(&customer.accountNumber)
	fmt.Print("Masukkan saldo awal: ")
	fmt.Scan(&customer.balance)
	fmt.Print("Masukkan nomor kartu: ")
	fmt.Scan(&customer.cardNumber)
	fmt.Print("Masukkan PIN: ")
	fmt.Scan(&customer.PIN)
	fmt.Print("Masukkan NIK: ")
	fmt.Scan(&customer.NIK)
	fmt.Print("Masukkan nama: ")
	fmt.Scan(&customer.name)
	fmt.Print("Masukkan alamat (district, city, province): ")
	fmt.Scan(&customer.address.district, &customer.address.city, &customer.address.province)

	customerBank.customers[customerBank.nCustomer] = customer
	customerBank.nCustomer++

	fmt.Println("Registrasi berhasil.")
}
