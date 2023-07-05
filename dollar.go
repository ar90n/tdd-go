package main

type Dollar struct {
	amount int
}

func NewDollar(amount int) *Dollar {
	return &Dollar{
		amount: amount,
	}
}

func (d *Dollar) Times(multiplier int) *Dollar {
	return NewDollar(d.amount * multiplier)
}

func (d *Dollar) Equals(object interface{}) bool {
	dollar := object.(*Dollar)
	return d.amount == dollar.amount
}
