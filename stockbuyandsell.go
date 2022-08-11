package main
import (
	"fmt"
)
func profit(price []int)int{
		max:=0
		var buyPrice,sellPrice int
		for i:=0;i<len(price);i++{
			for j:=i;j<len(price);j++{
				buyPrice=price[i]
				sellPrice=price[j]
				if sellPrice-buyPrice>max{
					max=sellPrice-buyPrice
				}
			}
		}

return max
}
func main(){
	var price=[]int{1,2,3,4}
	prof:=profit(price)
	fmt.Println("The max profit is ",prof)
}