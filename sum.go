package main


type Sum struct{
	augend *Money
	addend *Money
}

func NewSum(augend, addend *Money ) *Sum {
	return &Sum{
		augend: augend,
		addend: addend,
	}
}

func (sum *Sum) Reduce (to string) *Money {
	amount := sum.augend.amount + sum.addend.amount
	return NewMoneyImpl(amount, to)
}
