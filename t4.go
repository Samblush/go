package main

import "fmt"

func p(x string, y string,z int) (string, string, int) {
	return x, y, z
}

func main() {
	a, b, c := p("Shubham", "Patna India", 18)
	fmt.Println("name = ",a)
	 fmt.Println("Address = ",b)
	  fmt.Println("Age = ",c)


}
