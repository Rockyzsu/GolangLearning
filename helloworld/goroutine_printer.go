package main
import "fmt"
func printer(c chan int){
    for{
        data:=<-c
        if data==0{
        break
        }

        fmt.Println(data)
    }

    c<-0
}
func main(){
// func body
    c:=make(chan int)
    go printer(c)
    for i:=1;i<10;i++{
    c<-i
    }

    c<-0
    <-c


}
