package main
import (
	"fmt"
	"math"
)

func main(){
	var input []int = make([]int, 7)
	input[0] =2445
	input[1] =4886
	input[2] =6134
	input[3] =7000
	input[4] =1235
	input[5] =3843
	input[6] =5133
	radixSort(input, 10, 4)
	for i:=0;i<7;i++{
		fmt.Println(input[i])
	}
}

func radixSort(input []int, radix int, width int){
	for i:=0; i < width; i++{
		radixSingleSort(input, i, radix)
	}

}

func radixSingleSort(input []int, position int, radix int){
	numOfItems:= len(input)
	var countingArray []int = make([]int, radix)
	for i:=0;i<numOfItems;i++{
		countingArray[getDigit(position, input[i], radix)]++
	}
	
	for j:=1; j< len(countingArray);j++{
		countingArray[j] +=countingArray[j -1]
	}

	var tempArray []int = make([] int, numOfItems)
	for tempIndex:=numOfItems-1; tempIndex >=0 ;tempIndex-- {
		countingArray[getDigit(position, input[tempIndex], radix)] -=1
		tempArray[countingArray[getDigit(position, input[tempIndex], radix)]] = input[tempIndex];
	}

	copy(input, tempArray)

}

func getDigit(position int, value int, radix int) int{
	return (value / int(math.Pow(float64(radix), float64(position)))) % radix
}