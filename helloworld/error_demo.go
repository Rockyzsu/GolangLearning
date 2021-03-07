package main
import (
    "errors"
    "fmt"
)

var errorDivisionByZero = errors.New("errorDivisionByZero")
func div(dividend,divisor int)(int,error){
    if divisor==0{
    return 0,errorDivisionByZero
}
return dividend/divisor,nil

}
func main(){
fmt.Println(div(10,0))
}
