package main

import (
	"fmt"
	"math"
)

type geometry interface {
	area() float64
	perim() float64
}

type geometryDisplayType interface {
	geometry
	display() string
}

type rect struct {
	width, heigth float64
}

type displayForRect struct {
	rect
	displaytype string
}

type circle struct {
	radius float64
}

type displayForCircle struct {
	circle
	displaytype string
}

func (r rect) area() float64 {
	return r.width * r.heigth
}

func (r rect) perim() float64 {
	return 2*r.width + 2*r.heigth
}

func (r displayForRect) display() string {
	return r.displaytype
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

func (c displayForCircle) display() string {
	return c.displaytype
}

func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}

func measureWithDisplay(g geometryDisplayType) {
	fmt.Println(g.display())
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}

func main() {
	r := rect{width: 3, heigth: 4}
	c := circle{radius: 5}

	rdisplay := displayForRect{rect: r, displaytype: "Rectangle:  "}
	cdisplay := displayForCircle{circle: c, displaytype: "Circle:  "}

	measureWithDisplay(rdisplay)
	measureWithDisplay(cdisplay)
}
