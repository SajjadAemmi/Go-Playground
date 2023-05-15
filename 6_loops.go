package main
import ("fmt")

func main() {
	for i:=0; i < 5; i++ {
	if i == 3 {
		continue
	}
	if i == 4 {
		break
	}
	fmt.Println(i)
	}

	adj := [2]string{"big", "tasty"}
	fruits := [3]string{"apple", "orange", "banana"}
	for i:=0; i < len(adj); i++ {
		for j:=0; j < len(fruits); j++ {
			fmt.Println(adj[i],fruits[j])
		}
	}

	for idx, val := range fruits {
		fmt.Printf("%v\t%v\n", idx, val)
	}
}
