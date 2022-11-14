package main

import "fmt"

func main() {

	groupCity := map[int][]string{
		10:   []string{"Деревня", "Село"},        // города с населением 10-99 тыс. человек
		100:  []string{"Город", "Большой город"}, // города с населением 100-999 тыс. человек
		1000: []string{"Миллионик"},              // города с населением 1000 тыс. человек и более
	}
	cityPopulation := map[string]int{
		"Село":      100,
		"Миллионик": 500,
	}

	for i := range cityPopulation {
		//fmt.Println(i)

		//if cityPopulation[f]<=100{
		var city bool
		for _, c := range groupCity[100] {
			//fmt.Println(c)

			if c == i {
				city = true
				break
			}
		}
		if !city {
			// city = false {
			delete(cityPopulation, i)
		}
	}
	fmt.Println(cityPopulation)
}

//fmt.Println(cityPopulation)

//fmt.Println(cityPopulation)

//fmt.Println(a)
//}
//if cityPopulation[f]<=100{

//}

//}
//}
//}
//fmt.Println(cityPopulation)
//}
