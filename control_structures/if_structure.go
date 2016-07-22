package control_structures
import "fmt"
func fizzBuzz(i int32) string {
	if i% 3 ==0 && i % 5 ==0 {
		return "FizzBuzz"
	} else if i % 5 ==0 {
		return "Buzz"
	} else if i%3==0 {
		return "Fizz"
	}else {
		return fmt.Sprintf("%d",i)
	}
}
