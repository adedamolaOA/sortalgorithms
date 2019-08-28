package main

import (
	"fmt"
)

func main(){

	insertionSort(40, -1, 6, -15, 30, 11, 3, 9);

}

func insertionSort(input ...int){
	for i:=1; i<len(input); i++{
		newElement := input[i]
		var j int
		for j=i;j>0 && input[j -1] > newElement; j --{
			input[j] = input[j - 1]
		}
		input[j] = newElement
	}

	for i:=0; i<len(input); i++{
		fmt.Println(input[i])
	}
}