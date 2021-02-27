package main

import "fmt"

func max(number1,number2 int) int{
    if number1>number2{
    return number1
}else{
return number2
}
}

func getMA(days_list...int)(int,int){
    length := 0
    sum :=0
    for _,v :=range days_list{
    length+=1
    sum+=v
    }
    return length,sum
}

func main(){
    _max := max(10,29)
    fmt.Println("max number is ",_max)
    //var arr [10]int
    //arr={1,2,3,4,5,6,7,8,9,10}
    a :=1
    b :=4
    c :=5
    length,sum :=getMA(a,b,c)
    fmt.Println("result is ",length,sum)
}
