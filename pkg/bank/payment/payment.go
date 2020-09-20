package payment

import (
	"bank/pkg/bank/types"
)

// Max f
func Max(payments []types.Payment) types.Payment {
	max := payments[0]
	for _, payment := range payments {
		if max.Amount < payment.Amount {
			max = payment
		}
	}
	return max
}

// PaymentSource f
func PaymentSource(cards []types.Card) []types.PaymentSource {
	var sources []types.PaymentSource
	for _, card := range cards {
		if !card.Active || card.Balance < 0 {
			continue
		}

		var source types.PaymentSource = types.PaymentSource{
			Type: card.Name,
			Number: card.PAN,
			Balance: card.Balance,
		}

		sources = append(sources, source)
	}

	return sources
}
