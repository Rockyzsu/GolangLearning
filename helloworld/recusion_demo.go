package main
import "fmt"


func recursion(n int) int{
if n<1{
return 1
}
return recursion(n-1)*n
}

func main(){
// func body
i:=20
result:=recursion(i)
fmt.Println(result)
}
