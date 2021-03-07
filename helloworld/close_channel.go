package main
import "fmt"

func main(){
// func body

ch :=make(chan int,2)
ch<-1
ch<-3
close(ch)

for i:=0;i<cap(ch)+1;i++{
    v,ok:=<-ch
    fmt.Println(v,ok)
}

fmt.Println("end of fun")
}
