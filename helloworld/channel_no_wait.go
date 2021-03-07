package main
import "fmt"
import "time"
func sender(ch chan int){
    for i:=0;i<100;i++{
        ch<-i
        time.Sleep(time.Millisecond)
        //time.Sleep(time.Second)
        //ch<-i+1
    }

    result:=<-ch
    if result==0{
        fmt.Println("stop")
    }
}

func main() {

    ch:=make(chan int)
    go sender(ch)

    for {
        data,ok:=<-ch
        if ok {
            fmt.Println("data=",data)

            if data==99 {
                break
            }

       } else {
        fmt.Println("No data")
        }
    }

    fmt.Println("Done")
    ch<-0
}


