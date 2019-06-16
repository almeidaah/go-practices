/*
Copyright 2014 Google Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

     http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package main

import (
	"fmt"

	//"github.com/golang/example/stringutil"
)

func main() {
	//fmt.Println(stringutil.Reverse("!selpmaxe oG ,olleH"))
	//fmt.Println("Almeida teste");

	//x := 10;
	//fmt.Println(x);
	//max := 10
	//printFor(max);
	//max = 100;
	//printByThree(max);
	chapterFiveExerc();
}

func printFor(max int){
	i := 1;
	for i <= max{
		fmt.Println(i);

	switch i{
		case 5 : fmt.Println("CHEGOU NO CINCO")	
	}
		i++;
	}

}

func chapterFiveExerc(){
	fmt.Println("CHAP 5")
	x := [5]float64{12,12,12,99,88}
	fmt.Println(x[4]) //should print 88

	y := make([]float32, 3 , 9)
	fmt.Println(len(y))

	z := [6]string{"a", "b", "c", "d", "e", "f"}
	fmt.Println(z[2:5]) //should print c d e,  index [5] excludes "f" 

	arr :=[]int{
		48,96,86,68,
		57,82,63,70,
		37,34,83,27,
		19,97,9,17,
	}

	discoverSmallNum(arr);

}

func discoverSmallNum(arr [] int){
	min := 209 ;
	for _,v := range arr {
		if( v < min ){
			min = v
		}
	}
	fmt.Println("menor valor : " , min);
}

func printByThree(max int){
	constTree := 3;
	i := 1;
	for i <= max{
		fmt.Println("NUM ATUAL : ", i)
		if((i % constTree) == 0){
			fmt.Println(i, "É DIVISÍVEL");
			i++;
			continue;
		}
		i++;
	}
}
