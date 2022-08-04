package main
​
import "fmt"
​
func main() {
	matrix := [][]int{{1, 4, 6, 2, 54}, {0, 2, 6, 7, 3}, {5, 0, 9, 2, 3}, {6, 5, 77, 32, 6}, {78, 4, 60, 20, 14}}
	xIndex := make([]int, len(matrix))
	yIndex := make([]int, len(matrix[0]))
	fmt.Println("Before changing:\n ")
	for i := 0; i < len(matrix); i++ {
		fmt.Println(matrix[i])
	}
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if matrix[i][j] == 0 {
				xIndex[i] = 1
				yIndex[j] = 1
				break
			} /*else {
			//xIndex[i] = 0
			//yIndex[j] = 0
			}*/
		}
	}
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if xIndex[i] == 1 {
				matrix[i][j] = 0
			} else if yIndex[j] == 1 {
				matrix[i][j] = 0
			}
			//fmt.Println((matrix))
		}
	}
	fmt.Println("\n---------------\nAfter changing:\n ")
	for i := 0; i < len(matrix); i++ {
		fmt.Println(matrix[i])
	}
	/*fmt.Println("-------------")
	fmt.Println(xIndex)
	fmt.Println(yIndex)*/
}