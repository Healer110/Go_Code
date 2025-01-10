package main

import (
	"abstractDemo/practice/model"
	"fmt"
)

func main() {
	p := model.NewPerson("mary")
	fmt.Println(*p)
	p.SetAge(98)
	p.SetSalary(12000)
	fmt.Println(p.GetAge())
	fmt.Println(p.GetSalary())
	fmt.Println(*p)
	fmt.Printf("%p \n", p)

}
