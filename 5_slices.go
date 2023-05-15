package main
import ("fmt")


func main() {
	/*
	Slices are similar to arrays, but are more powerful and flexible.
	Like arrays, slices are also used to store multiple values of the same type in a single variable.
	However, unlike arrays, the length of a slice can grow and shrink as you see fit.
	*/

	myslice1 := []int{}
	fmt.Println(len(myslice1))
	fmt.Println(cap(myslice1))
	fmt.Println(myslice1)

	myslice1 = append(myslice1, 20, 21)
	fmt.Printf("myslice1 = %v\n", myslice1)
	fmt.Printf("length = %d\n", len(myslice1))
	fmt.Printf("capacity = %d\n", cap(myslice1))
  
	myslice2 := []string{"Go", "Slices", "Are", "Powerful"}
	fmt.Println(len(myslice2))
	fmt.Println(cap(myslice2))
	fmt.Println(myslice2)

	// Create a Slice From an Array
	arr1 := [6]int{10, 11, 12, 13, 14,15} // An array
	myslice := arr1[2:4] // A slice made from the array
  
	fmt.Printf("myslice = %v\n", myslice)
	fmt.Printf("length = %d\n", len(myslice))
	fmt.Printf("capacity = %d\n", cap(myslice))

	// Create a Slice With The make() Function

	myslice3 := make([]int, 5, 10)
	fmt.Printf("myslice3 = %v\n", myslice3)
	fmt.Printf("length = %d\n", len(myslice3))
	fmt.Printf("capacity = %d\n", cap(myslice3))

	// with omitted capacity
	myslice4 := make([]int, 5)
	fmt.Printf("myslice4 = %v\n", myslice4)
	fmt.Printf("length = %d\n", len(myslice4))
	fmt.Printf("capacity = %d\n", cap(myslice4))

	prices := []int{10,20,30}
	prices[2] = 50
	fmt.Println(prices[0])
	fmt.Println(prices[2])

	// Append One Slice To Another Slice
	myslice5 := []int{1,2,3}
	myslice6 := []int{4,5,6}
	myslice7 := append(myslice5, myslice6...)
  
	fmt.Printf("myslice7=%v\n", myslice7)
	fmt.Printf("length=%d\n", len(myslice7))
	fmt.Printf("capacity=%d\n", cap(myslice7))
}
