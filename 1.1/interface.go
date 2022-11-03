package main

import (
	"fmt"
)

type Payer interface {
	Pay(int) error
}
type Wallet struct {
	Cash  int
	Color string
}

func (w *Wallet) Pay(ammount int) error {
	 if w.Color != "RED" {
	// 	w.Color = "RED"
	// 	return nil
	// }

	if w.Cash < ammount {
		return fmt.Errorf("Не хватает денег в кошельке")

	}
	w.Cash -= ammount // w.Cash = w.Cash - ammount
	return nil

}

func Bue(p Payer) {
	err := p.Pay(10)
	if err != nil {
		panic(err)
	}
	fmt.Println("Спасибо за покупку через \n\n", p)
}

func main() {
	myWallet := &Wallet{Cash: 100}
	Bue(myWallet)
}
