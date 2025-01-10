package main

import (
	"fmt"
	"os"
)

// 人员信息结构体
type Emp struct {
	Id   int
	Name string
	Next *Emp
}

// 定义EmpLink, 代表哈希表中的不同的表
// 这里不带表头，第一个节点就会存放人员数据
type EpmLink struct {
	Head *Emp
}

// 根据ID找出人员，并返该人员信息对应的指针，没有的话返回nil
func (el *EpmLink) FindById(id int) *Emp {
	temp := el.Head
	if temp == nil {
		fmt.Println("链表为空...")
		return nil
	}

	for {
		if temp.Id == id {
			return temp
		}

		if temp.Next == nil {
			return nil
		}
		temp = temp.Next
	}
}

// 添加员工的方法, 按照从小到大的顺序添加
func (el *EpmLink) Insert(emp *Emp) {
	if el.Head == nil {
		el.Head = emp
	} else {
		temp := el.Head
		// ID 值比之前的head还小，换head，然后重新建链
		if temp.Id > emp.Id {
			el.Head = emp
			el.Head.Next = temp
			return
		}
		for {
			if temp.Id == emp.Id {
				fmt.Println("ID 重复，无法插入...")
				break
			} else if temp.Next != nil {
				if temp.Next.Id == emp.Id {
					fmt.Println("ID 重复，无法插入...")
					break
				}
			}

			if temp.Next == nil {
				temp.Next = emp
				break
			}

			if temp.Id < emp.Id && temp.Next.Id > emp.Id {
				emp.Next = temp.Next
				temp.Next = emp
				break
			}
			temp = temp.Next
		}
	}
}

// 定义哈希表，含有链表的数组
type HashTable struct {
	Link [7]EpmLink
}

func (ht *HashTable) Insert(emp *Emp) {
	// 使用散列函数，确定将人员添加到哪个链表
	linkNo := ht.HashFind(emp.Id)
	// 使用对应的EmpLink添加人员
	ht.Link[linkNo].Insert(emp)
}

// 编写方法，显示所有成员
func (ht *HashTable) ShowHashTable() {
	for i := 0; i < len(ht.Link); i++ {
		fmt.Printf("table[%d]: ", i)
		temp := ht.Link[i].Head
		for {
			if temp == nil {
				fmt.Print("链表为空...")
				break
			}
			fmt.Printf("%v-->", *temp)
			if temp.Next == nil {
				break
			}
			temp = temp.Next
		}
		fmt.Println()
	}
}

// 编写一个方法，完成查找
func (ht *HashTable) FindById(id int) *Emp {
	linkNo := ht.HashFind(id)
	return ht.Link[linkNo].FindById(id)

}

// 散列函数
func (ht *HashTable) HashFind(id int) int {
	return id % len(ht.Link)
}

func main() {
	var ht HashTable
	fmt.Println(ht.Link[0].Head == nil)

	var key int
	id := 0
	name := ""
	for {
		fmt.Println("================系统菜单================")
		fmt.Println("1 添加人员信息")
		fmt.Println("2 显示人员信息")
		fmt.Println("3 查找人员信息")
		fmt.Println("4 退出系统")
		fmt.Println("请输入选项(1-4): ")
		fmt.Scanln(&key)
		switch key {
		case 1:
			fmt.Println("请输入ID：")
			fmt.Scanln(&id)
			fmt.Println("请输入name：")
			fmt.Scanln(&name)
			emp := &Emp{
				Id:   id,
				Name: name,
			}
			ht.Insert(emp)
		case 2:
			ht.ShowHashTable()
		case 3:
			fmt.Println("请输入要查找的ID号：")
			fmt.Scanln(&id)
			emp := ht.FindById(id)
			if emp != nil {
				fmt.Println(*emp)
			} else {
				fmt.Println("未找到该人员信息")
			}
		case 4:
			os.Exit(0)
		default:
			fmt.Println("输入错误，请重新输入....")
		}
	}
}
