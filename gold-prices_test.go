package main

import (
	"goldwatcher/prices"
	"testing"
)

func TestGold_GetPrices(t *testing.T) {
	g := prices.Gold{
		Prices: nil,
		Client: client,
	}

	p, err := g.GetPrices()
	if err != nil {
		t.Error(err)
	}

	if p.Price != 1849 {
		t.Error("wrong price returned:", p.Price)
	}
}
