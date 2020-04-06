package main
import(
	"fmt"
)
func Swap(b []int, i int){
	temp := b[i]
	b[i] = b[i+1]
	b[i+1] = temp

}
func BubbleSort(a []int) {
	for i:=0; i<len(a)-1; i++ {
		for j:=0; j<len(a)-i-1; j++ {
			if(a[j] > a[j+1]){
				Swap(a, j)				
			}
		}
	}
}
func main() {
	fmt.Printf("Input the length of the input sequence: ")
	var n int
	//Read input length
	fmt.Scan(&n)

	fmt.Printf("Enter the sequence: ")
	//Empty slice to put the inputted sequence
	nums := make([]int, 0)
	var temp int
	
	for i:=0; i<n; i++ {
		fmt.Scan(&temp)
		nums = append(nums, temp)
	}

	BubbleSort(nums)

	for _, value := range nums {
		fmt.Printf("%d ",value)
	}
}