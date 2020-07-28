package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

//为结构体类型定义方法
func (v Vertex) Abs_1() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

type MyFloat float64
//为非结构体类型声明方法
func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

//正常的函数
func Abs_2(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := Vertex{3, 4}
	fmt.Println(v.Abs_1())
      fmt.Println(Abs_2(v))

	f := MyFloat(-math.Sqrt2)
	fmt.Println(f.Abs())
}

