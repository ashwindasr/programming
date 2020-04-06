package main
import(
	"fmt"
	"sync"
	"math/rand"
)

type ChopS struct { sync.Mutex }
type Philo struct{
	leftCS, rightCS *ChopS
}

func (p Philo) eat(wg *sync.WaitGroup, j int){
	for i:=0; i<3; i++ {
		p.leftCS.Lock()
		p.rightCS.Lock()
		fmt.Println("Starting to eat", j)

		fmt.Println("Finished Eating", j)
		p.rightCS.Unlock()
		p.leftCS.Unlock()
	}
	wg.Done()
}
func Host(i int, c chan int){
	c <- i
}
func main(){
	var wg sync.WaitGroup

	CSticks := make([]*ChopS, 5)
	for i:=0; i<5; i++{
		CSticks[i] = new(ChopS)
	}
	philos := make([]*Philo, 5)
	for i:=0; i<5; i++ {
		choice := rand.Intn(2)
		if(choice == 0){
			philos[i] = &Philo{CSticks[i], CSticks[(i+1)%5]}
		}else{
			philos[i] = &Philo{CSticks[(i+1)%5], CSticks[i]}
		}
		
	}
	c := make(chan int, 2)
	for i:=0; i<5; i++ {
		go Host(1,c)
		select{
		case <- c: 
			wg.Add(1)
			go philos[i].eat(&wg, i+1)
		}

		

	}
	wg.Wait()
}
