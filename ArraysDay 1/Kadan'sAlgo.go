package main
import (
	"fmt"
)
func sumArray(arr []int)int{
	sumA:=0
	for i:=0;i<len(arr);i++{
		sumA+=arr[i]
	}
	return sumA
}
func kadane(slice []int)int{
	sum:=0
	subArray:=[]int{}
	for i:=0;i<len(slice);i++{
		for j:=i;j<len(slice);j++{
			if i==j{
				subArray=[]int{slice[i]}
				if sum<sumArray(subArray){
					sum=sumArray(subArray)
				}
			}else{
			subArray=slice[i:j]
			if sumArray(subArray)>sum{
				sum=sumArray(subArray)
			}
		}
		}
	}
	return sum
}
func main(){
	var slice=[]int{-2,1,-3,4,-1,2,1,-5,4}
	sum1:=kadane(slice)
	fmt.Println(sum1)

}