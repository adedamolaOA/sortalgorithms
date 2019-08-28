package main

import ("fmt")

func main(){
	var input []int = make([]int, 7)
	input[0] = 10
	input[1] =-4
	input[2] =8
	input[3] =7
	input[4] =1
	input[5] =-55
	input[6] =5
	quicksort(input, 0, 7)
	for i:=0; i<7;i++{
		fmt.Println(input[i])
	}
}

func quicksort(input []int, start int, end int){
	if(end - start < 2){
		return 
	}
	pivotIndex:=partition(input, start, end)
	fmt.Println("*** ",pivotIndex)
	quicksort(input, start, pivotIndex)
	quicksort(input, pivotIndex+1, end)

}

func partition(input []int, start int, end int) int {
	pivot:=input[start]
	i:=start
	j:=end
	for i < j{
		for i < j {
			j--
			if input[j] < pivot {
				input[i] = input[j]
				break
			}
		}
		for i < j {
			i++
			if input[i] > pivot {
				input[j] = input[i]
				break
			}
		}
	}
	
	input[j] = pivot
	return j

}