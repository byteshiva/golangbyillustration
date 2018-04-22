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
	Display()
}

type stddisplay struct {
	Name string
}

func (d stddisplay) Display() {
	fmt.Println("-- Type -- ", d.Name)
}

type rect struct {
	width, heigth float64
}

type displayForRect struct {
	rect
	stddisplay
}

type circle struct {
	radius float64
}

type displayForCircle struct {
	circle
	stddisplay
}

func (r rect) area() float64 {
	return r.width * r.heigth
}

func (r rect) perim() float64 {
	return 2*r.width + 2*r.heigth
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}

func measureWithDisplay(g geometryDisplayType) {
	g.Display()
	// fmt.Println(g)
	fmt.Println("Area: ", g.area())
	fmt.Println("Perimeter : ", g.perim())
}

func main() {
	r := rect{width: 3, heigth: 4}
	rDisplay := stddisplay{Name: ": Rectangle "}
	c := circle{radius: 5}
	cDisplay := stddisplay{Name: ": Circle "}

	rdisplay := displayForRect{rect: r, stddisplay: rDisplay}
	cdisplay := displayForCircle{circle: c, stddisplay: cDisplay}

	measureWithDisplay(rdisplay)
	measureWithDisplay(cdisplay)
}
