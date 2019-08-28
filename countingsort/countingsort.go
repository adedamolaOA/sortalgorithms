package main
import ("fmt")

func main(){
	var input []int = make([]int, 7)
	input[0] = 2
	input[1] =4
	input[2] =6
	input[3] =7
	input[4] =1
	input[5] =3
	input[6] =5
	countingSort(input, 1, 7)
	for i:=0; i< 7; i++{
		fmt.Println(input[i])
	}

}

func countingSort(input []int, min int, max int){
	//numOfItems:= len(input)
	var countingArray []int  = make([]int, (max -min)+1)
	for i:=0; i< len(countingArray);i++{
		countingArray[input[i] - min]++;
	}

	var k int
	for j:=min; j <=max;j++{
		for countingArray[j - min] > 0 {
			
			input[k] = j
			countingArray[j - min]--
			k++
		}
	}

}