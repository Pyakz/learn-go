package main

import (
	"fmt"
	"math"
)

type geometry interface {
	areas() float64
	perim() float64
	show() string
}

type rect struct {
	width, height float64
}
type circle struct {
	radius float64
}

func (r rect) areas() float64 {
	return r.width * r.height
}
func (r rect) perim() float64 {
	return 2*r.width + 2*r.height
}
func (r rect) show() string {
	return fmt.Sprintf("SHOW REC: %v", r.areas())
}

func (c circle) areas() float64 {
	return math.Pi * c.radius * c.radius
}
func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.areas())
	fmt.Println(g.perim())
	fmt.Println(g.show())

}

func (c circle) show() string {
	return fmt.Sprintf("SHOW CIRCLE: %v", c.areas())
}

func main() {
	r := rect{width: 3, height: 4}
	c := circle{radius: 5}

	measure(r)
	measure(c)

}
