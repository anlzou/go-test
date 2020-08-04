package main

import "fmt"

func fibonacci2(n int) (int,int) {
  if n < 2 {
    return 0,n
  }
  a,b := fibonacci2(n-1)
  return b,a+b
}

func fibonacci(n int) int {
  _,b := fibonacci2(n)
  return b
}

func main(){
	var i int
	for i = 0; i < 10; i++ {
		fmt.Printf("%d\t", fibonacci(i))
	}
	fmt.Println()
}

/*
更好的一种 fibonacci 实现，用到多返回值特性，降低复杂度
*/
