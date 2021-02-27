package main
import "fmt"
var name string = "Kobe Bryant"
func scope(){
    fmt.Println(name)
    var name string = "Leborn James"
    fmt.Println(name)
}

func main(){
// func body
scope()
}
