/*
 @Author f.almeida
*/

package main

import (
	"fmt"

	//"github.com/golang/example/stringutil"
)

func main(){
	//q
	fmt.Println("Started chapter6 exercises")
	//var intersint [] := [4,5,6]
	var x [3]int
	x[0] = 2
	x[1] = 5
	x[2] = 30
	sum(3,4,5)

	//2
	//fmt.Println(half(0));
	fmt.Println(half(1));
	fmt.Println(half(2));
	fmt.Println(half(4)); fmt.Println(half(12));

	//3
	fmt.Println("BIGGER IS : ", discoverBiggerNum(10,20,30))
	
	//4
	nextOdd := makeOddGenerator();
	fmt.Println(nextOdd())
	fmt.Println(nextOdd())
	fmt.Println(nextOdd())

	//5
	fmt.Println("FIBO : ", calcFibo(6))

	//10
	a := 1
	b := 2
	
	fmt.Println("A : ", a);
	fmt.Println("B : ", b);

	swap(&a, &b)
	
	fmt.Println("A : ", a);
	fmt.Println("B : ", b);
}


func swap(a *int, b *int){
	*a = 2
	*b = 1

}

func calcFibo(x uint) uint {
		if (x == 0){
			return 0
		}
		if(x == 1){
			return 1
		}
		return calcFibo(x-1) + calcFibo(x-2)
}

func sum(nums ... int) int{
	summ := 0;
	for v := range nums {
		summ += v
	}

	fmt.Println("sum : ", sum)
	return summ
}

func discoverBiggerNum(arr ... int) int{
	var bigger  int
	for _,value := range arr{
		if(value > bigger){
			bigger = value
		}
	}
	return bigger
}

/** Function inside funcrion [closure */
func makeOddGenerator() func() uint{
	i := uint(1)
	return func () (ret uint){
		ret = i
		i += 2
		return
	}
}

func half(num int) (int, bool){
	halfValue := num/2
	fmt.Println("HALF : ", halfValue)

	if(halfValue % 2 == 0){
		return halfValue,true	
	}
	return halfValue, false

}