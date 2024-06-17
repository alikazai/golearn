package main

import (
	"fmt"
	"os"
)

type bill struct {
	name  string
	items map[string]float64
	tip   float64
}

func newBill(n string) bill {
	return bill{
		name:  n,
		items: map[string]float64{},
		tip:   0,
	}
}

func (b bill) format() string {
	fs := "Bill breakdown \n"
	var total float64 = b.tip

	for k, v := range b.items {
		fs += fmt.Sprintf("%-25v ...$%v\n", k+":", v)
		total += v
	}

	fs += fmt.Sprintf("%-25v ...$%0.2f\n", "tip:", b.tip)
	fs += fmt.Sprintf("%-25v ...$%0.2f\n", "total:", total)

	return fs
}

func (b *bill) updateTip(t float64) {
	b.tip = t
}

func (b *bill) additem(n string, p float64) {
	b.items[n] = p
}

func (b *bill) save() {
	data := []byte(b.format())

	err := os.WriteFile("./bills/"+b.name+".txt", data, 0644)

	if err != nil {
		panic(err)
	}

	fmt.Println("Bill was saved to file")
}
