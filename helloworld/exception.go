package main
import "fmt"



func Sqrt(f float32)(float32,error){
if f<0{
    return 0,errors.New("math:square root of negative number") // error<F6>
}

return f,nil
}


func testcase(){

    result,err:=Sqrt(-1)
    if err!=nil{
    fmt.Println(err)
    }
}



func main(){
// func body

//error
//testcase()

}
