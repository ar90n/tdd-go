package main

import (
	"fmt"
)

type Bank struct {
	rates map[string]int
}

func NewBank() *Bank {
	return &Bank{
		rates: map[string]int{},
	}
}

func (b *Bank) Reduce (source Expression, to string) *Money {
	return source.Reduce(b, to)
}

func (b *Bank) AddRate(from, to string, rate int) {
	key := fmt.Sprintf("%s_%s", from, to)
	b.rates[key] = rate
}

func (b *Bank) Rate(from, to string) int {
	if from == to {
		return 1
	}

	key := fmt.Sprintf("%s_%s", from, to)
	return b.rates[key]
}
