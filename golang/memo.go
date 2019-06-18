package main

import "fmt"

type Animal struct {
	Name string
	Age  int
}

type Cat struct {
	Animal
	ServantName string
}

func main() {
	// 1
	kuro := Cat{}
	kuro.Name = "kuro"
	fmt.Println(kuro)

	// 2
	shiro := Cat{Animal{"shiro", 5}, "kjfsm"}
	fmt.Println(shiro)
}
