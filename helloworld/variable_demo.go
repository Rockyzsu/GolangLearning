package main
import "fmt"

func main(){
var b bool = true
fmt.Println(b)

var a string = "hello world"
var c int = 20

fmt.Println("a=")
fmt.Println(a)
fmt.Println("c=")
fmt.Println(c)
var intVal int
intVal,intVal2 :=1,2
fmt.Println(intVal)
fmt.Println(intVal2)
addr := &intVal
fmt.Println(addr)
e :=99 //if no identify : , if will raise error
fmt.Println(e)

//const
const pi=3.14
var area float32
area=2*pi*2
fmt.Println("area ")
fmt.Println(area)
const (male=1
female=2)

fmt.Println("male")
fmt.Println(male)
fmt.Println("female")
fmt.Println(female)

const (
a1=iota
b1
c1
d1="ha"
e1

)

fmt.Println(a1,b1,c1,d1,e1)
}
