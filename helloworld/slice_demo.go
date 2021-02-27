package main
import "fmt"


func range_demo(arr []int){
//fmt.Println(arr)
for index,num:=range arr {
fmt.Println(index," ",num)
}
}


func main(){
// func body
arr :=[]int{1,2,3}
fmt.Println("len ",len(arr))
fmt.Println("cap ",cap(arr))

//arr[4]=0
//fmt.Println("arr ",arr)

number1:=make([]int,0,5)
fmt.Println("nubmer1 ",number1)
fmt.Println("len ",len(number1))
fmt.Println("cap ",cap(number1))
//fmt.Println("nubmer1 ",number1)

arr1:=append(arr,99)
fmt.Println("new arr",arr)
fmt.Println("new arr",arr1)


range_demo(arr1)
}
