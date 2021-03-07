package main
import "fmt"

func main(){
    a:=10
    b:=20
    a=a^b
    b=a^b
    a=a^b
    fmt.Printf("a=%v,b=%v",a,b)

    c:=0
    d:=4
    d=d^c
    fmt.Printf("c=%v d=%v",c,d)
}
