package main
import "fmt"


func sum(num []int,c chan int){

    sum:=0
    for _,v:=range num{
    sum+=v
    }
    fmt.Println(sum)
    c<-sum
}



func fibonacci(n int,c chan int){
    x,y:=0,1
    for i:=0;i<n;i++{
    c<-x
    x,y=y,x+y

    }
    close(c)

}


func main_fib(){
    c:=make(chan int,10)
    go fibonacci(cap(c),c)
    for i:= range c{
    fmt.Println(i)
    }
}


func main(){
// func body


s:=[]int{1,2,3,4,5,6,7,8,9,10,11,12,13,14,15,16}
c:=make(chan int)
go sum(s[:len(s)/2],c)
go sum(s[len(s)/2:],c)
x,y:=<-c,<-c
fmt.Println(x,y,x+y)

ch :=make(chan int,1)

ch<-1
fmt.Println(<-ch)
//fmt.Println(<-ch) //it will lead deadlock

ch<-2
fmt.Println(<-ch)

main_fib()
}
