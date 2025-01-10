package model

// 定义一个结构体
type student struct {
	Name  string
	Score float64
}

/*
因为student结构体首字母小写，因此是只能在model包使用
使用工厂模式可以解决，首字母小写还可以在外部引用的情况
*/
func NewStudengt(name string, score float64) *student {
	return &student{
		Name:  name,
		Score: score,
	}
}
