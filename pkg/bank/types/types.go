package types

// Money t
type Money int64

// Currency t
type Currency string

// PAN t
type PAN string

const (
	// TJS v
	TJS Currency = "TJS"

	// RUB v
	RUB Currency = "RUB"

	// USD v
	USD Currency = "USD"
)

// Card t
type Card struct {
	ID         int
	PAN        PAN
	Balance    Money
	MinBalance Money
	Currency   Currency
	Color      string
	Name       string
	Active     bool
}

// Payment t
type Payment struct {
	ID     int
	Amount Money
}

// PaymentSource t
type PaymentSource struct {
	Type    string
	Number  PAN
	Balance Money
}
