package main
import ("fmt")

func myFunction(x int, y int) int {
	return x + y
}

func factorial_recursion(x float64) (y float64) {
	if x > 0 {
		y = x * factorial_recursion(x-1)
	} else {
		y = 1
	}
	return
}

func main() {
	total := myFunction(1, 2)
	fmt.Println(total)

	fmt.Println(factorial_recursion(4))
}
