package main
import (
	"fmt"
)
func main(){
	var slice =[]int {2,0,2,1,1,0}
	for i:=0;i<len(slice)-1;i++{
		if slice[i]>slice[i+1]{
			temp:=slice[i]
			slice[i]=slice[i+1]
			slice[i+1]=temp
			i=-1
		}
	}
	fmt.Println(slice)
}