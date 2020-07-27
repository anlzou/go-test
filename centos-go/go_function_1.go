package main

import "fmt"

/* 函数返回两个数的最大值 */
func max(num1, num2 int) int {
   /* 声明局部变量 */
   var result int

   if (num1 > num2) {
      result = num1
   } else {
      result = num2
   }
   return result
}

func main(){
	var_1 := max(2,3)
	fmt.Println(var_1)
}
