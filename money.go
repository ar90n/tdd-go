package main

type Money struct {
	amount int
}

func NewMoney(amount int) *Money {
	return &Money{
		amount: amount,
	}
	
}

func (d *Money) Times(n int) *Money {
	return &Money{
		amount: d.amount * n,
	}
}

func (d *Money) Equals(a *Money) bool {
	return d.amount == a.amount
}