package main
import(
	"fmt"
	"strings"
	"bufio"
	"os"
	"strconv"
	"sync"
	"sort"
)

func sorting(a []int, wg *sync.WaitGroup, c chan []int){
	sort.Ints(a)
	c <- a
	wg.Done()
}

func main(){
	in := bufio.NewReader(os.Stdin)

	fmt.Printf("Enter number seqence seperated by spaces (Atleast 4): ")
	nums, _ := in.ReadString('\n')
	nums = nums[:len(nums)-2]

	var numarr []int
	for _, val := range strings.Split(nums, " ") {
		temp, _ := strconv.Atoi(val)
		numarr = append(numarr, temp)
	}
	numerator := len(numarr) / 4
	denominator := len(numarr) % 4
	c := make(chan []int)
	var wg sync.WaitGroup

	wg.Add(4)
	i := 0
	for i=0; i<3*numerator; i+=numerator {
		fmt.Println(numarr[i:i+numerator])
		go sorting(numarr[i:i+numerator], &wg, c)
	}
	fmt.Println(numarr[i:i+numerator+denominator])
	go sorting(numarr[i:i+numerator+denominator], &wg, c)
	


	arr := <- c	
	for i:=0; i<3; i++ {
		temp := <- c
		var temparr []int
		
		x := 0
		y := 0
		for x < len(arr) && y < len(temp) {
			if(arr[x] <= temp[y]){
				temparr = append(temparr, arr[x])
				x++
			}else{
				temparr = append(temparr, temp[y])
				y++
			}
		}
		for x < len(arr){
				temparr = append(temparr, arr[x])
				x++
			}
		for y < len(temp){
				temparr = append(temparr, temp[y])
				y++
			}
			
		
		arr = temparr

	}
	fmt.Printf("Sorted array: ")
	fmt.Println(arr)
	wg.Wait()
}