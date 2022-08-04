package main
import (
	"fmt"
)
func main(){
// declaring a slice of slices of
    // type integer with a length of 3
   
    preslice :=make([]int,0)
    for i:=1;i<=5;i++{
        slice :=make([]int,i)
        for j:=0;j<i;j++{
            if j==0 || j==i-1{
                slice[j]=1
            }else
            {
                slice[j]=preslice[j]+preslice[j-1]
            }

        }
        preslice=slice
        fmt.Println(slice)
    }

}