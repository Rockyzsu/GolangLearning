package main
import "fmt"

func main(){
    ch:=make(chan int,3)
    fmt.Println("buffer len is ",len(ch))


    ch<-1
    ch<-1
    ch<-1
    ch<-1
    //ch<-1

    fmt.Println("buffer len is ",len(ch))
    fmt.Println("End of main")

}
