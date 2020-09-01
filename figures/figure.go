package figures

import (
	"fmt"
	"math"
)

// Circle struct
type Circle struct {
	Radius float64
}

// Square struct
type Square struct {
	Base float64
}

// Figure interface
type Figure interface {
	Area() float64
}

// receiver function
func (c *Circle) Area() float64 {
	return math.Pi * math.Pow(c.Radius, 2)
}

// receiver function
func (s *Square) Area() float64 {
	return s.Base * s.Base
}

// common logic using interfaces
func GetFigureArea(f Figure) {
	fmt.Println(f.Area())
}


/*func main() {
	// how to declare and assign a variable in go
	// var x int
	// x = 5

	// how to do the same above in one step
	// x := 5

	// Print the variable using fmt
	fmt.Println(x)

	// slices
	//y := []string{}
	 z := []string{ "esto es go", "no sé"}
	 z = append(z, "esto es un string")

	//maps
	 //list := map[string]int{}
	 //test := make([]int, 4)

	//control structures
	// if in go
	 if x > 5 {
	 	fmt.Println("Es falso")
	 } else if x < 5 {
		 fmt.Println("Ok no")
	 } else {
	 	fmt.Println("")
	 }

	// switch in go
	 expression := "+"
	switch expression {
	case "+":
		fmt.Println("Ok es suma")
	case "-":
		fmt.Println("Ok es resta")
	default:
		fmt.Println("Ni idea")
	}

	//loops in go
	for x:= 0; x < 6; x++ {
		fmt.Printf("El número es %d", x)
	}

	y := 0
	for {
		fmt.Println("Numero", y)
		if y == 6 {
			break
		}
		y++
	}

	for key, value := range z {
		fmt.Println("key", key, "value", value)
	}
}*/