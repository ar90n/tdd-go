package main


type Sum struct{
	augend Expression
	addend Expression
}

func NewSum(augend, addend Expression) *Sum {
	return &Sum{
		augend: augend,
		addend: addend,
	}
}

func (sum *Sum) Reduce (bank *Bank, to string) *Money {
	amount := sum.augend.Reduce(bank, to).amount + sum.addend.Reduce(bank, to).amount
	return NewMoneyImpl(amount, to)
}

func (sum *Sum) Plus (addend Expression) Expression {
	return NewSum(sum, addend)
}

func (sum *Sum) Times (multiplier int) Expression {
	return NewSum(sum.augend.Times(multiplier), sum.addend.Times(multiplier))
}