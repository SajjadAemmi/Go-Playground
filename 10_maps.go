package main
import ("fmt")

func main() {
	var a = map[string]string{"brand": "Ford", "model": "Mustang", "year": "1964"}
	b := map[string]int{"Oslo": 1, "Bergen": 2, "Trondheim": 3, "Stavanger": 4}

	fmt.Printf("a\t%v\n", a)
	fmt.Printf("b\t%v\n", b)

	var c = make(map[string]string) // The map is empty now
	c["brand"] = "Ford"
	c["model"] = "Mustang"
	c["year"] = "1964"

	fmt.Printf(c["brand"])
				
	c["year"] = "1970" // Updating an element
	c["color"] = "red" // Adding an element

	fmt.Println(c)

	delete(c,"year")

	fmt.Println(c)

	for k, v := range c {
		fmt.Printf("%v : %v, ", k, v)
	}

	d := make(map[string]int)
	d["Oslo"] = 1
	d["Bergen"] = 2
	d["Trondheim"] = 3
	d["Stavanger"] = 4

	fmt.Printf("c\t%v\n", c)
	fmt.Printf("d\t%v\n", d)
}
