package main
import (
	"fmt"
	"bufio"
	"os"
	"encoding/json"
)
func main(){
	m := make(map[string]string)
	in := bufio.NewReader(os.Stdin)

	fmt.Printf("Enter name: ")
	name, _ := in.ReadString('\n')
	fmt.Printf("Enter address: ")
	address, _ := in.ReadString('\n')
	
	m["name"] = name
	m["address"] = address

	barr, _ := json.Marshal(m)
	fmt.Printf(string(barr))
}