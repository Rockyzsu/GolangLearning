package main
import "fmt"

func main(){
// func body
//var my_array [10]float32
var my_array [10]int
//my_array={1,2,3}
var i int
for i=0;i<10;i++{
my_array[i]=i
}
fmt.Println(my_array)
}
