package main
import "fmt"

func main(){

    // func body
    //var my_array [10]float32

    var my_array [10]int
    //my_array={1,2,3}
    var i int
    for i=0;i<10;i++{

        my_array[i]=i

    }

    fmt.Println(my_array)

    var balance =[5]float32{1,2,3,4,5}
    fmt.Println()
    fmt.Printf("%T",balance)

    balance2 :=[5]int{0:99,3:88}

    fmt.Println(balance2)
    fmt.Printf("%T",balance2)
}
