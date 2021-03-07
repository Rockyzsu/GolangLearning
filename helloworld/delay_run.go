package main
import "fmt"
func main(){

defer fmt.Println("After main end 0")
defer fmt.Println("After main end 1")
defer fmt.Println("After main end 2")
defer fmt.Println("After main end 3 ")
fmt.Println("main start")
fmt.Println("main end")

}
