package main
import "fmt"

type book struct{
 name string
 author string
 price float32
 //isbn [13]int
}


func printBook(textbook book){
fmt.Println(textbook.name)
fmt.Println(textbook.author)
fmt.Println(textbook.price)
}

func main(){
// func body

    //python=book("python","rocky",40.0,{1,2,3,4,5,6,7,8,9,10,11,12,13})
    python:=book{name:"python",author:"locky",price:40.0}
    java:=book{"java","joe",40.0}
    fmt.Println("struct is python  ",python)
    fmt.Println("struct is java ",java)


    var ruby book
    ruby.name="Ruby"
    ruby.author="Tom"
    ruby.price=22.22

    fmt.Println("new book ruby ",ruby)

    printBook(ruby)

}
