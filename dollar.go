package main

type Dollar struct {
	money *MoneyImpl
}

func NewDollar(amount int) Money {
	return &Dollar{
		money: NewMoneyImpl(amount, "USD"),
	}
}

func (d *Dollar) Times(n int) Money {
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

func (d *Dollar) Currency() string {
	return d.money.Currency()
}