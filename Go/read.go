package main
import(
	"fmt"
	"os"
	"bufio"
	"strings"
)

type person struct{
	fname string
	lname string
}

func main(){
	fmt.Println("Enter filename in current directory:")
    var filename string
    fmt.Scan(&filename)
	var s []person
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("File does not exist")
	}else{
		
		scanner := bufio.NewScanner(file)
		for scanner.Scan(){
			name := strings.Fields(scanner.Text())
			s = append(s, person{fname: name[0], lname: name[1]})
		}
		
	}
	for _, obs := range s {
		fmt.Printf("%s %s\n", obs.fname, obs.lname)
	}
	file.Close()
}