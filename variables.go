package main
import ("fmt")

// declaring variables outside of a function, with the var keyword:
var a int
var b int = 2
var c = 3

// Since := is used outside of a function, running the program results in an error.
// d := 1

func main() {
  var student1 string = "John" //type is string
  var student2 = "Jane" //type is inferred
  x := 2 //type is inferred

  fmt.Println(student1)
  fmt.Println(student2)
  fmt.Println(x)

  var a bool = true     // Boolean
  var b int = 5         // Integer
  var c float32 = 3.14  // Floating point number
  var d string = "Hi!"  // String

  fmt.Println(a)
  fmt.Println(b)
  fmt.Println(c)
  fmt.Println(d)
  
  // In Go, all variables are initialized. 
  // So, if you declare a variable without an initial value, 
  // its value will be set to the default value of its type:
  var e string
  var f int
  var g bool

  fmt.Println(e)
  fmt.Println(f)
  fmt.Println(g)

  // If you use the type keyword, it is only possible to declare one type of variable per line.
  var h, i, j, k int = 1, 3, 5, 7

  fmt.Println(h)
  fmt.Println(i)
  fmt.Println(j)
  fmt.Println(k)

  // If the type keyword is not specified, you can declare different types of variables in the same line:
  var l, m = 6, "Hello"
  n, o := 7, "World!"

  fmt.Println(l)
  fmt.Println(m)
  fmt.Println(n)
  fmt.Println(o)

  const PI = 3.14
  const Q int = 1

  var r, s string = "Hello", "World"

  fmt.Print(r, s, "\n")
  fmt.Print(r, "\n")
  fmt.Print(s, "\n")
  fmt.Println(r, s) // fmt.Print(r, " ", s)
}
