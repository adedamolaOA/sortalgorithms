package main

import "fmt"

func main(){
	shellSort(8, -5, 9, 6, 10, -22, -15, 55, -60, 10)
}

func shellSort(input ...int){
	for gap:=len(input); gap > 0; gap --{
		for i:=gap; i < len(input); i++{
			newElement:=input[i]
			j:=i
			
			for j>=i && input[j - gap] > newElement{
				input[j] = input[j - gap]
				j -=gap
			}
			input[j] = newElement
		}
	}

	for i:=0; i < len(input); i++{
		fmt.Println(input[i])
	}

}