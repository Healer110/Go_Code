package main

import "fmt"

type Cat struct {
	Name string
	Age  uint8
}

func main() {
	var allChan chan interface{} = make(chan interface{}, 10)

	allChan <- Cat{"cat1", 3}
	allChan <- Cat{"cat2", 6}
	allChan <- 11
	allChan <- 11.89
	allChan <- "Best"

	// 需要将类型做转换
	// c := <-allChan
	c := (<-allChan).(Cat)

	fmt.Printf("%T\n", c)
	fmt.Println(c.Name, c.Age)

}
