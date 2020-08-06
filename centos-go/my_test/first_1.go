/*
 * @Date        : 2020-08-05 23:16:28
 * @LastEditors : anlzou
 * @Github      : https://github.com/anlzou
 * @LastEditTime: 2020-08-06 01:22:45
 * @FilePath    : \linux-go_test\centos-go\my_test\first_1.go
 * @Describe    : 
 */
package main

import "fmt"

func main() {
	for n := 1; n <= 20; n++ {
		result := 1
		/*
		 * ^作二元运算符就是异或，包括符号位在内，相同为0，不相同为1
		 * (2)10 = (10)2, (64)10 = (10000000)2
		 * 10^10000000 = 10000010 = (66)10
		 */
		for i := 0; i < 2^64; i++ {
			result = (result * n) % 115763
			fmt.Println("i = ",i,", result = ",result)
		}
		fmt.Printf("pow(%d, pow(2, 64)) mod 115763 = %d\n", n, result)
	}
}