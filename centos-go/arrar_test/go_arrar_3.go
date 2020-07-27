package main
import "fmt"

func main(){
	fun_1()
}

/*
- Go 语言的数组是值，其长度是其类型的一部分，作为函数参数时，是 值传递，函数中的修改对调用者不可见

- Go 语言中对数组的处理，一般采用 切片 的方式，切片包含对底层数组内容的引用，作为函数参数时，类似于 指针传递，函数中的修改对调用者可见
*/
func fun_1(){
	// 数组
	b := [...]int{2, 3, 5, 7, 11, 13}

	boo(b)
	fmt.Println(b) // [2 3 5 7 11 13]
	// 切片
	p := []int{2, 3, 5, 7, 11, 13}

	poo(p)
	fmt.Println(p)  // [13 3 5 7 11 2]
}

func boo(tt [6]int) {
    tt[0], tt[len(tt)-1] = tt[len(tt)-1], tt[0]
}

func poo(tt []int) {
    tt[0], tt[len(tt)-1] = tt[len(tt)-1], tt[0]
}
