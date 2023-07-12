package main

type Franc struct {
	money *Money
}

func NewFranc(amount int) *Franc {
	return &Franc{
		money: NewMoney(amount),
	}
}

func (d *Franc) Times(n int) *Franc {
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
