package main
import "os"
import "fmt"

func main(){
    exitChan :=make(chan int)
    go server("127.0.0.1:7001",exitChan)

    fmt.Println("server starting....")

    code:=<-exitChan
    os.Exit(code)
}
