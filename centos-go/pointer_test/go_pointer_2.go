package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

//结构体方法
func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)//90+160
}

/*
若使用值接收者，那么 Scale 方法会对原始 Vertex 值的副本进行操作。
*/
func (v* Vertex) Scale_1(f float64) {
	v.X = v.X * f
	fmt.Println("v.X",v.X)//30
	
	v.Y = v.Y * f//40
}

//普通法法
func Abs_2(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

/*
含有结构体指针的参数
*/
func Scale_2(v *Vertex, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	v := Vertex{3, 4}//初始化的时候x,y已经赋值
	v.Scale_1(10)//x、y乘以10
	fmt.Println(v.Abs())

	v2 := Vertex{3, 3}//初始化的时候x,y已经赋值
	Scale_2(&v2, 20)
	fmt.Println(Abs_2(v2))
}

