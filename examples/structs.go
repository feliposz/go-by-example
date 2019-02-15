package examples

import (
	"fmt"
	"math"
)

// Struct

type person struct {
	name string
	age  int
}

func structExample() {
	a := person{"Alice", 37}
	b := person{name: "Bob", age: 37}
	c := &b
	a.age++
	c.age = 40
	fmt.Println(a)
	fmt.Println(b)
}

// Methods

func methodExamples() {
	r := rect{width: 10, height: 20}
	r.grow(5)
	fmt.Println("Rect:", r, "Area:", r.area(), "Perim:", r.perim())
}

type rect struct {
	width, height float64
}

func (r rect) area() float64 {
	return r.width * r.height
}

func (r rect) perim() float64 {
	return r.width*2 + r.height*2
}

func (r *rect) grow(amount float64) {
	r.width += amount
	r.height += amount
}

// Interfaces

type geometry interface {
	area() float64
	perim() float64
}

type circle struct {
	radius float64
}

func (c circle) area() float64 {
	return math.Pow(c.radius, 2) * math.Pi
}

func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

func describeGeometry(g geometry) {
	fmt.Println("Dimesions:", g, "Area:", g.area(), "Perim:", g.perim())
}

func interfaceExamples() {
	r := rect{10, 20}
	c := circle{10}
	describeGeometry(r)
	describeGeometry(c)
}

func StructExamples() {
	fmt.Println("\nSome struct examples")
	fmt.Println("====================")

	structExample()
	methodExamples()
	interfaceExamples()
}
