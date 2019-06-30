package main

import "fmt"

func main(){
   evens := []int{2,4,6,8,10}
   alphabetFruitMap := map[string]string{ "a": "Apple", "b":"Banana", "c":"Citrus" }

   for _, even := range evens{
	fmt.Println( even)
   }

   for k,v := range alphabetFruitMap {
	fmt.Printf("%s --> %s\n", k, v)
   }

   var stringIteration = "AazZ"

   for c,u := range stringIteration{
	fmt.Println(c ,u)

   }
   
}
