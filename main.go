package main

import "go-basics/figures"

// An example of using packages in golang
func main() {
	c := figures.Circle{Radius: 1.25}
	figures.GetFigureArea(&c)

	s := figures.Square{Base: 5}
	figures.GetFigureArea(&s)
}