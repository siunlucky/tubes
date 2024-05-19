package main

import "fmt"

func dummyData() {
    worldBank.Banks[0] = Bank{
        name:       "Bank A",
        uniqueCode: 123,
        nBranch:    2,
        branch: [NMAX_Bank]Address{
            {district: "Dist1", city: "City1", province: "Prov1"},
            {district: "Dist2", city: "City2", province: "Prov2"},
        },
        nCustomer: 0,
    }
    worldBank.Banks[1] = Bank{
        name:       "Bank B",
        uniqueCode: 221,
        nBranch:    2,
        branch: [NMAX_Bank]Address{
            {district: "Dist3", city: "City3", province: "Prov3"},
            {district: "Dist4", city: "City4", province: "Prov4"},
        },
        nCustomer: 0,
    }
    worldBank.nBank = 2

    dummyCustomers := []Customer{
        {accountNumber: 1001, balance: 0, PIN: "123456", NIK: 1234567890123451, name: "Dummy1", address: Address{district: "Dist1", city: "City1", province: "Prov1"}, bankCode: 123, branchAddress: Address{district: "Dist1", city: "City1", province: "Prov1"}, isSuspended: false},
        {accountNumber: 1002, balance: 0, PIN: "123456", NIK: 1234567890123452, name: "Dummy2", address: Address{district: "Dist2", city: "City2", province: "Prov2"}, bankCode: 456, branchAddress: Address{district: "Dist2", city: "City2", province: "Prov2"}, isSuspended: false},
        {accountNumber: 1003, balance: 0, PIN: "123456", NIK: 1234567890123453, name: "Dummy3", address: Address{district: "Dist3", city: "City3", province: "Prov3"}, bankCode: 789, branchAddress: Address{district: "Dist3", city: "City3", province: "Prov3"}, isSuspended: false},
        {accountNumber: 1004, balance: 0, PIN: "123456", NIK: 1234567890123454, name: "Dummy4", address: Address{district: "Dist4", city: "City4", province: "Prov4"}, bankCode: 111, branchAddress: Address{district: "Dist4", city: "City4", province: "Prov4"}, isSuspended: false},
        {accountNumber: 1005, balance: 0, PIN: "123456", NIK: 1234567890123455, name: "Dummy5", address: Address{district: "Dist5", city: "City5", province: "Prov5"}, bankCode: 222, branchAddress: Address{district: "Dist5", city: "City5", province: "Prov5"}, isSuspended: false},
    }

    for i := 0; i < len(dummyCustomers); i++ {
        if customerBank.nCustomer < NMAX {
            customerBank.customers[customerBank.nCustomer] = dummyCustomers[i]
            customerBank.nCustomer++
        } else {
            fmt.Println("Customer bank is full, unable to add more customers")
            fmt.Println("Back to Main Menu")
            loadingAuth("Loading")
            fmt.Println()
            menu()
        }
    }
}