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