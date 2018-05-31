package main
import "fmt"

func Task (n int) int {
	result:= 1 
	for {
		result*=2 
		if result > n { break }
	}
	return result
}

func main() {
	var n int
	n = 100
	fmt.Println("The result is :", Task(n))
}
