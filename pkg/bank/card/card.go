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

// PaymentSources f
func PaymentSources(cards []types.Card) []types.PaymentSource {
	var sources []types.PaymentSource
	for _, card := range cards {
		if !card.Active || card.Balance < 0 {
			continue
		}

		var source types.PaymentSource = types.PaymentSource{
			Type:    card.Name,
			Number:  card.PAN,
			Balance: card.Balance,
		}

		sources = append(sources, source)
	}

	return sources
}
