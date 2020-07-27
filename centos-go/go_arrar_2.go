package main

import "fmt"

func main() {
   /* 数组长度为 5 */
   var  balance = [5]int {1000, 2, 3, 17, 50}
   var avg float32

   /* 数组作为参数传递给函数 */
   avg = getAverage( balance, 5 ) ;

   /* 输出返回的平均值 */
   fmt.Println( "平均值为: %f",avg);

   fun_1()
   fun_2()
}
func getAverage(arr [5]int, size int) float32 {
   var i,sum int
   var avg float32  

   for i = 0; i < size;i++ {
      sum += arr[i]
   }

   avg = float32(sum) / float32(size)

   return avg;
}

//精度测试1
func fun_1(){
    a := 1.69
    b := 1.7
    c := a * b      // 结果应该是2.873
    fmt.Println(c)  // 输出的是2.8729999999999998
}

func fun_2(){
    a := 1690           // 表示1.69
    b := 1700           // 表示1.70
    c := a * b          // 结果应该是2873000表示 2.873
    fmt.Println(c)      // 内部编码
    fmt.Println(float64(c) / 1000000) // 显示
}
