package main
import (
	"fmt"
)

type Animals interface{
	Eat()
	Move()
	Speak()
}

type Animal struct{
	name string
	food string
	motion string
	sound string
}
func (x Animal) Eat() {
	fmt.Println(x.food)
}
func (x Animal) Move() {
	fmt.Println(x.motion)
}
func (x Animal) Speak() {
	fmt.Println(x.sound)
}

func main(){
	animalsList := make(map[string]Animal)
	for true {
		fmt.Printf("> ")
		var command, name, param string
		fmt.Scan(&command, &name, &param)

		if(command == "newanimal"){
			if(param == "cow"){
				animalsList[name] = Animal{"cow", "grass", "walk", "moo"}
				fmt.Println("Created it!")
			}else if(param == "bird"){
				animalsList[name] = Animal{"bird", "worms", "fly", "peep"}
				fmt.Println("Created it!")
			}else if(param == "snake"){
				animalsList[name] = Animal{"snake", "mice", "slither", "hsss"}
				fmt.Println("Created it!")
			}else{
				fmt.Println("Invalid Entry. Try again.")
			}

		}else if(command == "query"){
			if(param == "eat"){
				var ob Animals
				ob = animalsList[name]
				ob.Eat()
			}else if(param == "move"){
				var ob Animals
				ob = animalsList[name]
				ob.Move()
			}else if(param == "speak"){
				var ob Animals
				ob = animalsList[name]
				ob.Speak()
			}else{
				fmt.Println("Invalid Entry. Try again")
			}
		}else{
			fmt.Println("Invalid Entry")
		}
	}
}