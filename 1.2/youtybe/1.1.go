package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

func main() {

	slise := []string{}
	files, err := ioutil.ReadDir("/home/alex") //файлы из папки на компьютере
	if err != nil {
		log.Fatal(err)

	}
	for _, file := range files {
		//fmt.Println(len(files), file.Name())
		slise = append(slise, file.Name())

	}
	fmt.Println(len(slise), slise)

	massiv := make([]string, len(files))

	for a, b := range files {
		massiv[a] = b.Name()
	}

	fmt.Println(massiv, len(massiv))

}
