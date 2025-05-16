package data

import (
	"subgraphfirst/graph/model"
)

var customers = []*model.Customer{
	{
		CustomerReferenceID: "CUSTOMER1",
		Name:                "Goutam Kumar",
		Last4ssn:            "1111",
		Age:                 40,
		Address:             "1234 XYZ Street, Test AZ 12345",
		Accounts:            sulAccounts,
		Email:               "customer1@gmail.com",
	},
	{
		CustomerReferenceID: "CUSTOMER2",
		Name:                "Ravi Kumar",
		Last4ssn:            "2222",
		Age:                 20,
		Address:             "4567 XYZ Street, Test AZ 23456",
		Accounts:            bisAccounts,
		Email:               "customer2@gmail.com",
	},
	{
		CustomerReferenceID: "CUSTOMER3",
		Name:                "Saurav Kumar",
		Last4ssn:            "3333",
		Age:                 10,
		Address:             "8901 XYZ Street, Test AZ 34567",
		Accounts:            farAccounts,
		Email:               "customer3@gmail.com",
	},
	{
		CustomerReferenceID: "dddddddddd",
		Name:                "Pradip Kumar",
		Last4ssn:            "4444",
		Age:                 5,
		Address:             "2584 XYZ Street, Test AZ 45678",
		Accounts:            fayAccounts,
		Email:               "customer4@gmail.com",
	},
}
var sulAccounts = []*model.Account{
	{
		AccountReferenceID: "ACCOUNT_A1",
	},
	{
		AccountReferenceID: "ACCOUNT_A2",
	},
}
var bisAccounts = []*model.Account{
	{
		AccountReferenceID: "ACCOUNT_B1",
	},
	{
		AccountReferenceID: "ACCOUNT_B2",
	},
}
var farAccounts = []*model.Account{
	{
		AccountReferenceID: "ACCOUNT_C1",
	},
	{
		AccountReferenceID: "ACCOUNT_C2",
	},
}
var fayAccounts = []*model.Account{
	{
		AccountReferenceID: "ACCOUNT_D1",
	},
	{
		AccountReferenceID: "ACCOUNT_D2",
	},
}

func Customers() []*model.Customer {
	return customers
}

func SulAccounts() []*model.Account {
	return sulAccounts
}

func BisAccounts() []*model.Account {
	return bisAccounts
}
