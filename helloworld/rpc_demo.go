package main
import "fmt"
import "time"
import "errors"
func RPCClient(ch chan string,req string)(string,error){
ch<-req
select {
case ack:=<-ch:
return ack,nil
case <-time.After(time.Second):
    return "",errors.New("Time out in customer define")
}
}

func RPCServer(ch chan string){

for{

    data:=<-ch
    fmt.Println("server received ",data)
    time.Sleep(2*time.Second)
    ch<-"roger"
    fmt.Println("next step")
}
}


func main(){
    ch:=make(chan string)
    go RPCServer(ch)
    recv,err:=RPCClient(ch,"hello")
    if err!=nil{
    fmt.Println("error")
    fmt.Println(err)
    } else{
        fmt.Println("client received : ",recv)
    }
    fmt.Println("End of main")
}
