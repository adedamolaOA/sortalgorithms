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
	mergeSort(input, 0, 7)
	for i:=0;i<len(input);i++{
		fmt.Println(input[i])
	}
}

func mergeSort(input []int, start int, end int){
	

	if end - start < 2 {
		return
	}

	mid:= (start+end)/2
	
	mergeSort(input, start, mid)
	mergeSort(input, mid, end)
	
	merge(input, start, mid, end)
}

func merge(input []int, start int, mid int, end int) {
	if input[mid - 1] <= input[mid]{
		return
	}
	i:=start
	j:=mid
	tempIndex:=0
	var tempLength int = start+end;
	var temp []int = make([]int, tempLength)
	for i < mid && j < end{
		if input[i] <= input[j] {
			temp[tempIndex] =input[i]
			i++
		}else{
			temp[tempIndex] = input[j]
			j++
		}
		tempIndex++;
	}
	for k:=0; k<(mid - i); k++ {
		input[start+tempIndex+k] = input[i+k]  
	}
	for k:=0; k < tempIndex; k++{
		input[start+k] = temp[k]
	}
	//copy(input[(start+tempIndex):((start+tempIndex)+(mid -i))],input[i:(i + (mid-i))])
	//copy(input[0:tempIndex], temp[start:(start+tempIndex)])

}