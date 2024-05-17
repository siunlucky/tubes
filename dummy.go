package main

import "fmt"

func dummyData() {
	dummyCustomers := []Customer{
		{accountNumber: 1001, balance: 0, PIN: "123456", NIK: 111111, name: "Dummy1", address: Address{district: "Dist1", city: "City1", province: "Prov1"}, isSuspended: false},
		{accountNumber: 1002, balance: 0, PIN: "123456", NIK: 111112, name: "Dummy2", address: Address{district: "Dist2", city: "City2", province: "Prov2"}, isSuspended: false},
		{accountNumber: 1003, balance: 0, PIN: "123456", NIK: 111113, name: "Dummy3", address: Address{district: "Dist3", city: "City3", province: "Prov3"}, isSuspended: false},
		{accountNumber: 1004, balance: 0, PIN: "123456", NIK: 111114, name: "Dummy4", address: Address{district: "Dist4", city: "City4", province: "Prov4"}, isSuspended: false},
		{accountNumber: 1005, balance: 0, PIN: "123456", NIK: 111115, name: "Dummy5", address: Address{district: "Dist5", city: "City5", province: "Prov5"}, isSuspended: false},
	}

	for i := 0; i < len(dummyCustomers); i++ {
		if customerBank.nCustomer < NMAX {
			customerBank.customers[customerBank.nCustomer] = dummyCustomers[i]
			customerBank.nCustomer++
		} else {
			fmt.Println("Customer bank is full, unable to add more customers")
			i = 2000
			
			fmt.Println("Back to Main Menu")
			loadingAuth("Loading")
			fmt.Println()
			menu()
		}
	}
}

