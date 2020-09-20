package payment

import (
	"bank/pkg/bank/types"
	"fmt"
)

func ExampleMax() {
	payments := []types.Payment{
		{
			ID: 66,
			Amount: 100,
		},
		{
			ID: 2,
			Amount: 98,
		},
		{
			ID: 3,
			Amount: 10,
		},
		{
			ID: 4,
			Amount: 99,
		},
	}
	
	fmt.Println(Max(payments))
	//Output: {66 100}
}