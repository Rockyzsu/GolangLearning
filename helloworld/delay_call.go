package main
import "fmt"
import "time"
func main(){
// func body
exit := make(chan int)
t := time.Second * 3
time.AfterFunc(t,func(){
fmt.Printf("After 3 second\n")
exit<-0
})

<-exit
fmt.Printf("end of main\n")
}
