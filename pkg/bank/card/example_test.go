package card

import (
	"bank/pkg/bank/types"
	"fmt"
)

func ExamplePaymentSources() {
	cards := []types.Card{
		{
			ID:         66,
			PAN:        "5555 xxxx xxxx 5555",
			Balance:    10_000_00,
			MinBalance: 5_000_00,
			Currency:   "TJS",
			Color:      "Black",
			Name:       "Card1",
			Active:     true,
		},
		{
			ID:         67,
			PAN:        "6666 xxxx xxxx 6666",
			Balance:    -5_000_00,
			MinBalance: -5_000_00,
			Currency:   "TJS",
			Color:      "Red",
			Name:       "Card2",
			Active:     true,
		},
		{
			ID:         68,
			PAN:        "7777 xxxx xxxx 7777",
			Balance:    20_000_00,
			MinBalance: 5_000_00,
			Currency:   "USD",
			Color:      "Black",
			Name:       "Card3",
			Active:     true,
		},
	}
	result := PaymentSources(cards)
	fmt.Println(result)

	// Output: [{Card1 5555 xxxx xxxx 5555 1000000} {Card3 7777 xxxx xxxx 7777 2000000}]
}
