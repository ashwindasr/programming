package main
import(
	"fmt"
)
func GenDisplaceFunc(a,v,s float64) func(float64) float64{
	return func(t float64) float64{
		return 0.5*a*t*t + v*t + s
	}
}
func main(){
	var acc, vel, disp, time float64
	fmt.Printf("Enter intial acceleration: ")
	fmt.Scan(&acc)
	fmt.Printf("Enter intial velocity: ")
	fmt.Scan(&vel)
	fmt.Printf("Enter intial displacement: ")
	fmt.Scan(&disp)

	CalcDisp := GenDisplaceFunc(acc, vel ,disp)

	fmt.Printf("Enter time value: ")
	fmt.Scan(&time)
	
	ans := CalcDisp(time)

	fmt.Println(ans)
}