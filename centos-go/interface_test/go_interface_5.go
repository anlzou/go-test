package main

import(
	"fmt"
	//"strconv"
)

type IPAddr [4]byte

// TODO: 给 IPAddr 添加一个 "String() string" 方法 
func (ipAddr IPAddr) String() string{
	var ans string
	ans = ""
	len := len(ipAddr)
	for key,value := range ipAddr{
		//fmt.Println(value,len)
		ans += string(value)
		if(key < len-1){
			ans += "."
		}
	}
  	return ans
}

func main() {
	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}
}

/*
todo:存在type转string的乱码

题目描述：
通过让 IPAddr 类型实现 fmt.Stringer 来打印点号分隔的地址。
例如，IPAddr{1, 2, 3, 4} 应当打印为 "1.2.3.4"。
*/
