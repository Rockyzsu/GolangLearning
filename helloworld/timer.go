package main
import "fmt"
import "time"

func main(){
// func body
i:=0
ticker :=time.NewTicker(time.Millisecond*100)
timer :=time.NewTimer(time.Second*5) //exit main after 5s
for {
select {
case <-timer.C:
    fmt.Println("Time up exit !")
    goto Stop

case <-ticker.C:
    fmt.Println("Current time ",i)
    i++
}

}
Stop:
    fmt.Println("End of main")
}
