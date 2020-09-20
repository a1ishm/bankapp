package card

import (
	"bank/pkg/bank/types"
)

// Total f
func Total(cards []types.Card) types.Money {
	var total types.Money
	for _, card := range cards {
		if !card.Active || card.Balance < 0 {
			continue
		}
		total += card.Balance
	}
	return total
}