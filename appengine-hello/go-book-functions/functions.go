package main
import (
	"fmt"
)
func main(){
	makeRecursiveFunction();
}

func makeRecursiveFunction(){

	x := func(x,y int) int{
		return x + y
	}

	fmt.Println(x(4,5))
}