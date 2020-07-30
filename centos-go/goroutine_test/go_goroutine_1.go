package main

import (
	"fmt"
	"time"
)

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func main() {
	go say("world")
	say("hello")

/*-----------------part 2-------------------*/
	s := []int{7, 2, 8, -9, 4, 0}

	c := make(chan int)
	go sum(s[:len(s)/2], c) //-9+4+0 = -5
	go sum(s[len(s)/2:], c) //7+2+8 = 17
	x, y := <-c, <-c //从 c 中接收

	fmt.Println(x, y, x+y)
}

/*
Go 程（goroutine）是由 Go 运行时管理的轻量级线程。
Go 程在相同的地址空间中运行，因此在访问共享的内存时必须进行同步。sync 包提供了这种能力，不过在 Go 中并不经常用到，因为还有其它的办法（见下）。
*/

func sum(s []int, c chan int){//chan：通信管道
	sum := 0
	for _, v := range s{
		sum += v
	}
	c <- sum //将和送入 c
}

