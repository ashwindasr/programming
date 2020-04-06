package main
import (
	"fmt"
	"strconv"
	"sort"
)
func main() {

	slice := make([]int, 0, 3)
	
	for true {
		var num string
		fmt.Printf("Enter the numbers below. Enter X or x to stop: ")
		fmt.Scan(&num)
		if(num == "x" || num == "X"){
			break
		}
		i, _ := strconv.Atoi(num)
		slice = append(slice, i)

	
	sort.Ints(slice)
	fmt.Printf("The slice in sorted order are:\n")
	for _, val := range slice {
		fmt.Printf("%d \n", val )
	}
}
}
