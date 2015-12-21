package main

import "fmt"

type Rectangle struct {
	length, width int
}

type Cube struct {
	side int
}

func (r Rectangle) Area() int {
	return r.length * r.width
}

func (c Cube) Area() int {
	return c.side * c.side * c.side
}

func main() {
	rect := Rectangle{length: 5, width: 2}
	cube := Cube{side: 3}
	fmt.Println(rect.Area())
	fmt.Println(cube.Area())
}
