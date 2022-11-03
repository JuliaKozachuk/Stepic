package main

import "fmt"

type User struct {
	ID   int
	Name string
	Sex  string
	Age  int
	Date int
}
type Car struct {
	ID     int
	Brand  string
	UserId int
}

func main() {
	var cars []Car
	brands := [...]string{"toyota", "lexus", "kia", "nissan", "subary", "mazda", "reno"}

	//for i := 0; i <= 6; i++ {
	//fmt.Println(i)
	//cars = append(cars, Car{ID: i, Brand: brands[i]})
	//fmt.Println(i, brands[i])

	//}

	//fmt.Printf("")

	//fmt.Printf("hello")

	for i, brand := range brands { //для перебора
		cars = append(cars, Car{ID: i, Brand: brand})

	}
	for i, brand := range brands { //для перебора
		cars = append(cars, Car{ID: i, Brand: brand})

	}
	fmt.Println(cars)

}

//var mashine = Mashine{
//ID:    1,
//Brand: "toyota",
//}

// 	var user = User{
// 		ID:   1,
// 		Name: "Petr",
// 		Sex:  "Men",
// 		Age:  36,
// 		Date: 2022,
// 	}

// 	fmt.Println(user.ID)

// 	for i, car := range cars {

// 		car.UserId = user.ID
// 		//fmt.Println(car)
// 		cars[i] = car

// 	}
// 	fmt.Println(cars)
// }

// func (m *Mashine) purchase(user string) error {
// 	if m.Brand == "toyota" {
// 		return nil

// 	}
// 	func Pay(m.User)  {

// 	}

// }
