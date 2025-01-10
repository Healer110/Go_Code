package model

import "fmt"

// Person 结构体设置为私有化，通过工厂模式提供对外的api
type person struct {
	Name   string
	age    int
	salary float64
}

func NewPerson(name string) *person {
	return &person{
		Name: name,
	}
}

func (p *person) SetAge(age int) {
	if age > 0 && age < 150 {
		p.age = age
	} else {
		fmt.Println("年龄输入有误...")
	}
}

func (p *person) GetAge() int {
	return p.age
}

func (p *person) SetSalary(sal float64) {
	if sal > 0 {
		p.salary = sal
	} else {
		fmt.Println("薪资入有误...")
	}
}

func (p *person) GetSalary() float64 {
	return p.salary
}
