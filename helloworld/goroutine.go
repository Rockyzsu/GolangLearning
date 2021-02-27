package main

import (
"fmt"
"time"
)


func say(s string){
    for i:=0;i<5;i++{
    time.Sleep(100*time.Millisecond)
    fmt.Println(s)
    }
}

func main(){
// func body

go say("world")
//say("hello")
time.Sleep(1000*time.Millisecond)
}
