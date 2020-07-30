package main

import (
	"fmt"
)

func fibonacci(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c)
}

func main() {
	c := make(chan int, 10)
	go fibonacci(cap(c), c) //cap(c)：c的容量
	for i := range c {
		fmt.Println(i)
	}
}

/*
https://tour.go-zh.org/concurrency/4
*/
