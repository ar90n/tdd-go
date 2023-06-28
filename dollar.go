package main

type Dollar struct {
	amount int
}

func NewDollar(amount int) *Dollar {
	return &Dollar{
		amount: amount,
	}
}

func (d *Dollar) Times(n int) *Dollar {
	return &Dollar{
		amount: d.amount * n,
	}
}

func (d *Dollar) Equals(a *Dollar) bool {
	return d.amount == a.amount
}
