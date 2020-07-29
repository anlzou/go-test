package main

import "fmt"

func main() {
   var a int = 10  

   fmt.Printf("main():变量的地址: %x\n", &a  )

   fmt.Println("=====fun_1()=====")
   fun_1()

   fmt.Println("++++++++fun_2():range方法+++++++++")
   fun_2()

   fmt.Println("++++++++fun_2_2():for方法+++++++++")
   fun_2_2()
}

/*指针普通用法*/
func fun_1(){
   var a int= 20   /* 声明实际变量 */
   var ip *int        /* 声明指针变量 */

   ip = &a  /* 指针变量的存储地址 */

   fmt.Printf("a 变量的地址是: %x\n", &a  )

   /* 指针变量的存储地址 */
   fmt.Printf("ip 变量储存的指针地址: %x\n", ip )

   /* 使用指针访问值 */
   fmt.Printf("*ip 变量的值: %d\n", *ip )
}

/*指针数组使用range*/
func fun_2(){
    const max = 3

    number := [max]int{5, 6, 7}
    var ptrs [max]*int //指针数组
    //将number数组的值的地址赋给ptrs
    for i:= range number {
        ptrs[i] = &number[i]
    }
    for i, x := range ptrs {
        fmt.Printf("指针数组：索引:%d 值:%d 值的内存地址:%d\n", i, *x, x)
    }
}

/*指针数组使用for*/
func fun_2_2(){
    const max = 3

    number := [max]int{5, 6, 7}
    var ptrs [max]*int //指针数组
    //将number数组的值的地址赋给ptrs
    for i := 0; i < max; i++ {
        ptrs[i] = &number[i]
    }
    for i, x := range ptrs {
        fmt.Printf("指针数组：索引:%d 值:%d 值的内存地址:%d\n", i,*x, x)
    }
}
