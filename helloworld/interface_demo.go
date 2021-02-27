package main
import "fmt"


type Phone interface{
call()
}

type NokiaPhone struct{
//nothing
}

type IPhone struct{
//nothing
}


func (nokia NokiaPhone) call(){
fmt.Println("I am nokia, i am calling you !")
}


func (iphone IPhone) call(){
fmt.Println("I am iphone, i am calling you !")
}



func main(){
// func body

var phone Phone
phone = new(NokiaPhone)
phone.call()

phone = new(IPhone)
phone.call()

}
