package main

const NMAX = 1000
const NMAX_Bank = 10
const NMAX_Admin = 10

type Bank struct {
	name       string
	branch     [NMAX_Bank]Address
	nBranch    int
	uniqueCode int
	customers  [NMAX]Customer
	nCustomer  int
	admins     [NMAX_Admin]Admin
	nAdmin     int
}

type Admin struct {
	name string
	Credential
}

type WorldBank struct {
	Banks [NMAX_Bank]Bank
	nBank int
}

type Credential struct {
	username, password string
}

var SuperAdmin = Credential{
	username: "1",
	password: "1",
}

type Customer struct {
	accountNumber int
	balance       int
	transactions  []Transaction
	nTransaction  int
	cardNumber    int
	PIN           int
	NIK           int
	name          string
	address       Address
	isSuspended   bool
}

type CustomerBank struct {
	customers [NMAX]Customer
	nCustomer int
}

type Transaction struct {
	transactionId          int
	senderBankCode         int
	senderAccountNumber    int
	recipientBankCode      int
	recipientAccountNumber int
	amount                 int
}

type Address struct {
	district string
	city     string
	province string
}