package main
import(
	"fmt"
)
type Animal struct{
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
	cow := Animal{"grass", "walk", "moo"}
	bird := Animal{"worms", "fly", "peep"}
	snake := Animal{"mice", "slither", "hsss"}

	var animal, category string
	for true {
		fmt.Println("Enter the inputs. Type \"stop stop\" to end the loop.")
		fmt.Printf("> ")
		fmt.Scan(&animal, &category)
		if(animal == "stop"){break}
		if(animal == "cow"){
			if(category == "eat"){
				cow.Eat()
			}else if(category == "move"){
				cow.Move()
			}else if(category == "speak"){
				cow.Speak()
			}else{
				fmt.Println("Invalid entry. Try again")
			}
			
		}else if(animal == "bird"){
			if(category == "eat"){
				bird.Eat()
			}else if(category == "move"){
				bird.Move()
			}else if(category == "speak"){
				bird.Speak()
			}else{
				fmt.Println("Invalid entry. Try again")
			}
		}else{
			if(category == "eat"){
				snake.Eat()
			}else if(category == "move"){
				snake.Move()
			}else if(category == "speak"){
				snake.Speak()
			}else{
				fmt.Println("Invalid entry. Try again")
			}
		}
	}
}