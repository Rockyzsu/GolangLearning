package main
import "fmt"


func map_usage(){
//var countryMap map[string]string

countryMap := make(map[string]string)
countryMap["Shenzhen"]="0755"
countryMap["Shanghai"]="2121"
countryMap["Berling"]="0743"
countryMap["Hongkong"]="4755"
countryMap["Guilin"]="0455"
countryMap["Guanghzou"]="4355"
countryMap["Beijing"]="0735"

fmt.Println(countryMap)

delete(countryMap,"Hongkong")
for k,v :=range countryMap{
fmt.Println(k," ",v)
}
}


func main(){
// func body
map_usage()
}
