package main

import "fmt"

type Money struct {
	amount int
	currency string
}

func NewMoneyImpl(amount int, currency string) *Money {
	return &Money{
		amount: amount,
		currency: currency,
	}	
}

func (d *Money) Times(n int) *Money {
	return &Money{
		amount: d.amount * n,
		currency: d.currency,
	}
}

func (d *Money) Equals(a interface{}) bool {
	m, ok := a.(*Money)
	if !ok {
		return false
	}
	return d.amount == m.amount && d.currency == m.currency
}

func (d *Money) Currency() string {
	return d.currency
}

func (d *Money) Plus(added *Money) Expression {
	return NewSum(d, added)
}

func (d *Money) Reduce(to string) *Money {
	return d
}

func (d *Money) ToString() string {
	return fmt.Sprintf("%v %v", d.amount, d.currency)
}

func NewDollar(amount int) *Money {
	return NewMoneyImpl(amount, "USD")
}

func NewFranc(amount int) *Money {
	return NewMoneyImpl(amount,"CHF")
}