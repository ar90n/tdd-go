package main

type Money interface {
	Times(n int) Money
	Equals(a interface{}) bool
	Currency() string
}

type MoneyImpl struct {
	amount int
	currency string
}

func NewMoneyImpl(amount int, currency string) *MoneyImpl {
	return &MoneyImpl{
		amount: amount,
		currency: currency,
	}	
}

func (d *MoneyImpl) Times(n int) *MoneyImpl {
	return &MoneyImpl{
		amount: d.amount * n,
		currency: d.currency,
	}
}

func (d *MoneyImpl) Equals(a *MoneyImpl) bool {
	return d.amount == a.amount
}

func (d *MoneyImpl) Currency() string {
	return d.currency
}