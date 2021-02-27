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

func swap_ref(x *int,y *int){
    temp:=*y
    *y=*x
    *x=temp
}

func getAverage(arr [10]int,size int) float32{
    i:=0
    sum:=0
    for ;i<size;i++{
    sum+=arr[i]
    }
    return float32(sum)/float32(size)
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

    x:=10
    y:=99
    fmt.Printf("origin x %d, origin y %d ",x,y)
    swap_ref(&x,&y)
    fmt.Printf("after  x %d, after y %d \n",x,y)
    score :=[10]int{7,7,7,7,7,7,7,7,5,4}
    result := getAverage(score,10)
    fmt.Println("avg is ",result)
}
