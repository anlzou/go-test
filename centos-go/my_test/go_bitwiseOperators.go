package main

import "fmt"

func main(){
	A := 60
	B := 13

	l := A&B
	m := A|B
	n := A^B

	fmt.Println(l,m,n)
}
