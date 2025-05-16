package data

import "subgraphthird/graph/model"

// ptrFloat64 is a helper function to return a pointer to a float64 value.
func ptrFloat64(f float64) *float64 {
	return &f
}

// Cards returns a list of sample cards for testing.
func Cards() []*model.Card {
	ptr := func(s string) *string { return &s }
	return []*model.Card{
		{
			CardReferenceID: "CARD001",
			Status:          ptr("Active"),
			Type:            ptr("Credit"),
			ExpiryDate:      ptr("2025-12-31"),
			CardNumber:      ptr("1234-5678-9012-3456"),
			Transactions: []*model.Transaction{
				{
					TransactionID:   "TXN001",
					Amount:          ptrFloat64(100.50),
					Currency:        ptr("USD"),
					Status:          ptr("Completed"),
					TransactionDate: ptr("2025-05-01"),
					AvailableCreditAmount: &model.AvailableCreditAmount{
						SpendingCreditAmount: ptrFloat64(5000.00),
						CashCreditAmount:     ptrFloat64(1000.00),
					},
				},

				{
					TransactionID:   "TXN002",
					Amount:          ptrFloat64(200.75),
					Currency:        ptr("USD"),
					Status:          ptr("Pending"),
					TransactionDate: ptr("2025-05-02"),
					AvailableCreditAmount: &model.AvailableCreditAmount{
						SpendingCreditAmount: ptrFloat64(4800.00),
						CashCreditAmount:     ptrFloat64(950.00),
					},
				},
			},
		},
		{
			CardReferenceID: "CARD002",
			Status:          ptr("Inactive"),
			Type:            ptr("Debit"),
			ExpiryDate:      ptr("2026-06-30"),
			CardNumber:      ptr("9876-5432-1098-7654"),
			Transactions: []*model.Transaction{
				{
					TransactionID:   "TXN003",
					Amount:          ptrFloat64(50.00),
					Currency:        ptr("EUR"),
					Status:          ptr("Completed"),
					TransactionDate: ptr("2025-04-15"),
					AvailableCreditAmount: &model.AvailableCreditAmount{
						SpendingCreditAmount: ptrFloat64(3000.00),
						CashCreditAmount:     ptrFloat64(500.00),
					},
				},
			},
		},
		{
			CardReferenceID: "CARD003",
			Status:          ptr("Inactive"),
			Type:            ptr("Debit"),
			ExpiryDate:      ptr("2026-06-30"),
			CardNumber:      ptr("9876-5432-1098-7654"),
			Transactions: []*model.Transaction{
				{
					TransactionID:   "TXN003",
					Amount:          ptrFloat64(50.00),
					Currency:        ptr("EUR"),
					Status:          ptr("Completed"),
					TransactionDate: ptr("2025-04-15"),
					AvailableCreditAmount: &model.AvailableCreditAmount{
						SpendingCreditAmount: ptrFloat64(3000.00),
						CashCreditAmount:     ptrFloat64(500.00),
					},
				},
			},
		},
	}
}
