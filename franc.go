package main

type Franc struct {
	money *MoneyImpl
}

func NewFranc(amount int) Money {
	return &Franc{
		money: NewMoneyImpl(amount, "CHF"),
	}
}

func (d *Franc) Times(n int) Money {
	return &Franc{
		money: d.money.Times(n),
	}
}

func (d *Franc) Equals(a interface{}) bool {
	franc, ok := a.(*Franc)
	if !ok {
		return false
	}
	return d.money.Equals(franc.money)
}

func (d *Franc) Currency() string {
	return d.money.Currency()
}
