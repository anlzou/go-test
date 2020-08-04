package main
import "fmt"

const(
	i = 1<<iota //iota 只是在同一个 const 常量组内递增，每当有新的 const 关键字时，iota 计数会重新开始
	j = 3<<iota
	k
	l
)

func test(){
	const (
	    i = iota
	    j = iota
	    x = iota
	)

	const xx = iota
	const yy = iota
	println(i, j, x, xx, yy)
}


/*==========main===========*/
func main(){
	fmt.Println("i=",i)
	fmt.Println("j=",j)
	fmt.Println("k=",k)
	fmt.Println("l=",l)

	test()
}
