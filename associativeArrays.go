package main

import "fmt"

func main(){
	var myKeys = make( map[string]int)
        v1 := 100;

	myKeys["k1"] = 1
	myKeys["k2"] = 2
        fmt.Println(myKeys, v1)
        _,k1 := myKeys["k1"]

	fmt.Println( k1)

	//delete(myKeys, "k1")


}