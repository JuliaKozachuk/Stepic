package main

import "fmt"

type Fighter struct {
	On    bool
	Ammo  int
	Power int
}

func (f *Fighter) Shoot() bool {
	if f.On && f.Ammo > 0 {
		f.Ammo = (f.Ammo - 1)
		return true
	} else {
		return false
	} //else if f.On == false {
	//return false
	//}

}
func (f *Fighter) RideBike() bool {
	if f.On && f.Power > 0 {
		f.Power = (f.Power - 1)
		return true
	} else {
		return false
	} //else if f.On == false {
	//return false
	//}

}
func main() {
	testStruct := new(Fighter)
	testStruct.On = true
	testStruct.Ammo, testStruct.Power = 10, 20
	fmt.Println(testStruct.Shoot(), testStruct.Ammo)
	fmt.Println(testStruct.RideBike(), testStruct.Power)
	fmt.Println(testStruct.Shoot(), testStruct.Ammo)
	fmt.Println(testStruct.RideBike(), testStruct.Power)
}
