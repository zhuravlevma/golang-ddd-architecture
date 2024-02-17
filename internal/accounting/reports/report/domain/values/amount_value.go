package values

type AmountValue struct {
	Amount float64
	Rate   float64
}

func NewAmount(amount float64, rate float64) *AmountValue {
	return &AmountValue{
		Amount: amount,
		Rate:   rate,
	}
}

func (amount *AmountValue) ApplyDiscount(discount float64) {
	amount.Amount *= discount
}

func (amount *AmountValue) GetAmoutWithoutTax() float64 {
	return amount.Amount * (100.0 - amount.Rate)
}

func (amount *AmountValue) DifferenceAfterTax() float64 {
	return amount.Amount - amount.GetAmoutWithoutTax()
}

func (amount *AmountValue) UpdateTaxRate(rate float64) {
	amount.Rate = rate
}
