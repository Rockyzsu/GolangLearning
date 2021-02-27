package main
import "fmt"
func main(){
var a uint=1
var b uint=0
high :=999
b=a<<2
fmt.Println("b=",b)
var ptr *int
ptr = &high
fmt.Println("pointer",ptr)
fmt.Println("pointer value" ,*ptr)

}
