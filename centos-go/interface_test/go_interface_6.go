package main
import "fmt"

type Phone interface{
	call()
}

type MiPhone struct{
}

func (miPhone MiPhone) call(){
	fmt.Println("小米手机使用的是MIUI")
}

type VivoPhone struct{
}

func (vivoPhone VivoPhone) call(){
	fmt.Println("vivo 手机，HIFI音乐手机")
}

func main(){
	var phone Phone
	phone = new(MiPhone)
	phone.call()

	phone = new(VivoPhone)
	phone.call()
}
