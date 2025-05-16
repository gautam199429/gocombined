package data

import (
	"subgraphsecond/graph/model"
)

func ptrFloat64(f float64) *float64 {
	return &f
}

func ptrString(s string) *string {
	return &s
}

var accounts = []*model.Account{
	{
		AccountReferenceID: "ACCOUNT_A1",
		Balance:            ptrFloat64(100.10),
		Type:               ptrString("Personal"),
		Status:             &model.AllAccountStatus[0],
		Cards:              cardsRef1,
	},
	{
		AccountReferenceID: "ACCOUNT_A2",
		Balance:            ptrFloat64(109990.10),
		Type:               ptrString("Business"),
		Status:             &model.AllAccountStatus[1],
		Cards:              cardsRef2,
	},
	{
		AccountReferenceID: "ACCOUNT_B1",
		Balance:            ptrFloat64(200.10),
		Type:               ptrString("Personal"),
		Status:             &model.AllAccountStatus[0],
		Cards:              cardsRef3,
	},
	{
		AccountReferenceID: "ACCOUNT_B2",
		Balance:            ptrFloat64(209990.10),
		Type:               ptrString("Business"),
		Status:             &model.AllAccountStatus[1],
		Cards:              cardsRef2,
	},
	{
		AccountReferenceID: "ACCOUNT_C1",
		Balance:            ptrFloat64(100.10),
		Type:               ptrString("Personal"),
		Status:             &model.AllAccountStatus[0],
		Cards:              cardsRef1,
	},
	{
		AccountReferenceID: "ACCOUNT_C2",
		Balance:            ptrFloat64(109990.10),
		Type:               ptrString("Business"),
		Status:             &model.AllAccountStatus[1],
		Cards:              cardsRef2,
	},
	{
		AccountReferenceID: "ACCOUNT_D1",
		Balance:            ptrFloat64(200.10),
		Type:               ptrString("Personal"),
		Status:             &model.AllAccountStatus[0],
		Cards:              cardsRef1,
	},
	{
		AccountReferenceID: "ACCOUNT_D2",
		Balance:            ptrFloat64(209990.10),
		Type:               ptrString("Business"),
		Status:             &model.AllAccountStatus[1],
		Cards:              cardsRef3,
	},
}

func Accounts() []*model.Account {
	return accounts
}

var cardsRef1 = []*model.Card{
	{
		CardReferenceID: "CARD001",
		Status:          ptrString("Active"),
	},
}

var cardsRef2 = []*model.Card{
	{
		CardReferenceID: "CARD002",
		Status:          ptrString("Inactive"),
	},
}

var cardsRef3 = []*model.Card{
	{
		CardReferenceID: "CARD003",
		Status:          ptrString("Inactive"),
	},
}
