package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMultiplication(t *testing.T) {
	five := NewDollar(5)
	assert.Equal(t, NewDollar(10), five.Times(2))
	assert.Equal(t, NewDollar(15), five.Times(3))
}

func TestEquality(t *testing.T) {
	assert.True(t, NewDollar(5).Equals(NewDollar(5)))
	assert.False(t, NewDollar(5).Equals(NewDollar(6)))
	assert.False(t, NewFranc(5).Equals(NewDollar(6)))
}

func TestCurrency(t *testing.T) {
	assert.Equal(t, "USD", NewDollar(1).Currency())
	assert.Equal(t, "CHF", NewFranc(1).Currency())
}

func TestSimpleAddition(t *testing.T) {
	five := NewDollar(5)
	bank := NewBank()
	//sum := NewDollar(5).Plus(NewDollar(5))
	sum := five.Plus(five)
	reduced := bank.Reduce(sum, "USD")
	assert.Equal(t, NewDollar(10), reduced)
}

func TestPlusReturnsSum(t *testing.T) {
	five := NewDollar(5)
	result := five.Plus(five)
	sum, ok := result.(*Sum)
	assert.True(t, ok)
	assert.Equal(t, five, sum.augend)
	assert.Equal(t, five, sum.addend)
}

func TestReduceSum(t *testing.T) {

	sum := NewSum(NewDollar(3), NewDollar(4))
	bank := NewBank()
	result := bank.Reduce(sum, "USD")
	assert.Equal(t, NewDollar(7), result)
}

func TestReduceMoney(t *testing.T) {
	bank := NewBank()
	result := bank.Reduce(NewDollar(1), "USD")
	assert.Equal(t, NewDollar(1), result)
}

func TestReduceMoneyDifferentCurrency(t *testing.T) {
	bank := NewBank()
	bank.AddRate("CHF", "USD", 2)
	result := bank.Reduce(NewFranc(2), "USD")
	assert.Equal(t, NewDollar(1), result)
}

func TestIdentityRate(t *testing.T) {
	assert.Equal(t, 1, NewBank().Rate("USD", "USD"))
}

func TestMixedAddition(t *testing.T){
	fiveBucks := NewDollar(5)
	tenFrancs := NewFranc(10)
	bank := NewBank()
	bank.AddRate("CHF", "USD", 2)
	result := bank.Reduce(fiveBucks.Plus(tenFrancs), "USD")
	assert.Equal(t, NewDollar(10), result)
}