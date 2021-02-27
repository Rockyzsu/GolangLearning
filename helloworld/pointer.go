package main
import "fmt"

func main(){
// func body

    a:=10
    fmt.Println("address ",&a)
    fmt.Println("value ",*(&a))

    var p *int
    p=&a
    fmt.Println("pointer value ",p)
    fmt.Println("pointer value * ",*p)

    //empty pointer

    var empty_pointer *int
    fmt.Println("ponter of empty_pointer ",empty_pointer)


    //array
    arr:=[]int{10,20,30}
    for i:=0;i<3;i++{
    fmt.Println("Arr ",arr[i])
    }

    // array pointer
    var arr_p [3]*int
    for i:=0;i<3;i++{
    arr_p[i]=&arr[i]
    }

    fmt.Println("Array pointer",arr_p)

}
