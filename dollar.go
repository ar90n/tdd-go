package main

type Dollar struct {
	money *Money
}

func NewDollar(amount int) *Dollar {
	return &Dollar{
		money: NewMoney(amount),
	}
}

func (d *Dollar) Times(n int) *Dollar {
	return &Dollar{
		money: d.money.Times(n),
	}
}

func (d *Dollar) Equals(a interface{}) bool {
	dollar, ok := a.(*Dollar)
	if !ok {
		return false
	}
	return d.money.Equals(dollar.money)
}
