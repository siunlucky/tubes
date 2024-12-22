package main

import (
	"fmt"
	"time"
)

func searchCustomerByCredentials(bankIdx, accountNumber int, PIN string) int {
	for i := 0; i < worldBank.Banks[bankIdx].nCustomer; i++ {
		if accountNumber == worldBank.Banks[bankIdx].customers[i].accountNumber && PIN == worldBank.Banks[bankIdx].customers[i].PIN {
			return i
		}
	}
	return -1
}

func searchCustomerByName(bankIdx int, search string) int {
	for i := 0; i < worldBank.Banks[bankIdx].nCustomer; i++ {
		if search == worldBank.Banks[bankIdx].customers[i].name {
			return i
		}
	}
	return -1

}

func searchCustomerByIdx(bankIdx, search int) int {
	if search >= 0 && search < worldBank.Banks[bankIdx].nCustomer {
		return search
	}

	return -1
}

func searchCustomerByNIK(bankIdx, search int) int {
	for i := 0; i < worldBank.Banks[bankIdx].nCustomer; i++ {
		if search == worldBank.Banks[bankIdx].customers[i].NIK {
			return i
		}
	}
	return -1
}

func searchCustomerByAccountNumber(bankIdx, search int) int {
	start := time.Now()
	time.Sleep(2 * time.Millisecond)
	for i := 0; i < worldBank.Banks[bankIdx].nCustomer; i++ {
		if search == worldBank.Banks[bankIdx].customers[i].accountNumber {
			elapsed := time.Since(start)
			fmt.Printf("searchCustomerByAccountNumber selesai dalam %s\n", elapsed)
			fmt.Printf("Waktu eksekusi (nanodetik): %d ns\n", elapsed.Nanoseconds())
			fmt.Printf("Waktu eksekusi (detik): %.9f s\n", elapsed.Seconds())
			return i
		}
	}
	
	elapsed := time.Since(start)
	fmt.Printf("searchCustomerByAccountNumber selesai dalam 2 %s\n", elapsed)
	return -1
}

func searchCustomerByAccountNumberv2(bankIdx, search, idx int) int {
	start := time.Now()
	time.Sleep(2 * time.Millisecond)

	if idx == 0 {
		start = time.Now()
		time.Sleep(2 * time.Millisecond)
	}

	if idx >= worldBank.Banks[bankIdx].nCustomer {
		if idx == 0 {
			elapsed := time.Since(start)
			fmt.Printf("searchCustomerByAccountNumberv2 selesai dalam %s\n", elapsed)
			fmt.Printf("Waktu eksekusi (nanodetik): %d ns\n", elapsed.Nanoseconds())
			fmt.Printf("Waktu eksekusi (detik): %.9f s\n", elapsed.Seconds())
		}
		return -1
	}

	if search == worldBank.Banks[bankIdx].customers[idx].accountNumber {
		elapsed := time.Since(start)
		fmt.Printf("searchCustomerByAccountNumberv2 selesai dalam %s\n", elapsed)
		fmt.Printf("Waktu eksekusi (nanodetik): %d ns\n", elapsed.Nanoseconds())
		fmt.Printf("Waktu eksekusi (detik): %.9f s\n", elapsed.Seconds())
		return idx
	}

	return searchCustomerByAccountNumberv2(bankIdx, search, idx+1)
}

func searchCustomerByCardNumber(bankIdx, search int) int {
	for i := 0; i < worldBank.Banks[bankIdx].nCustomer; i++ {
		if search == worldBank.Banks[bankIdx].customers[i].cardNumber {
			return i
		}
	}
	return -1
}
