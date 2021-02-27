package main

import "fmt"

func main(){

var start int
for start=0;start<10;start++{
    fmt.Println("start",start)
}

strings :=[]string{"Google","Baidu","tencent","threeArrow"}
for i,s:=range strings{
fmt.Println(i)
fmt.Println(s)

}
}
