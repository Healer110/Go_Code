package main

import "fmt"

type Monkey struct {
	varieties string
	color     string
}

type GolgenMonket struct {
	Monkey
	address string // 生活的地域
}

type Flying interface {
	Fly()
}

type Swimming interface {
	Swim()
}

func (gm *GolgenMonket) ShowMonkeyInfo() {
	fmt.Printf("猴子的品种：%v, 猴子身上的颜色：%v, 猴子生活的地方：%v \n",
		gm.varieties, gm.color, gm.address)
}

// 猴子想学习飞翔
func (gm *GolgenMonket) Fly() {
	fmt.Printf("%v会飞了.... \n", gm.varieties)
}

func (gm *GolgenMonket) Swim() {
	fmt.Printf("%v会游泳了.... \n", gm.varieties)
}

func main() {
	var monkey GolgenMonket = GolgenMonket{
		Monkey: Monkey{
			varieties: "金丝猴",
			color:     "金色",
		},
		address: "高山丛林",
	}
	monkey.ShowMonkeyInfo()
	monkey.Fly()
	monkey.Swim()
	// var f Flying = &monkey
	// f.Fly()
	// var s = &monkey
	// s.Swim()

}
