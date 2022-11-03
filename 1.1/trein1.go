package main

import "fmt"

type Applies interface {
	Ap(int) error
}

type Hennes struct {
	Head_office string
	Area        int
	Store       int
}

func (w *Hennes) Ap(Stores int) error {
	if w.Area == 7 {
		w.Store = 747
		return nil

	}
	if w.Store < Stores {
		return fmt.Errorf("Не относится к седьмой Эрии")
	}
	w.Store -= Stores{
		return nil
	}

}
func meaning(p Applies) {
	err := p.Ap(625)
	if err != nil {
		return nil, err
		panic(err)
	}
	fmt.Println("Магазин седьмой Эрии")

}

func main() {
	myStore := &Hennes{Store:682}
	meaning(myHennes)
}
