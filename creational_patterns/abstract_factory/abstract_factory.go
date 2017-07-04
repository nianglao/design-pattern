package main

import (
	"fmt"
)

var gTotal int64

type Shape interface {
	Draw()
}

type Circle struct {
	id int64
}

func NewCircle() Shape {
	gTotal++
	return Circle{id: gTotal}
}

func (c Circle) Draw() {
	fmt.Printf("Circle : %d\n", c.id)
}

type Square struct {
	id int64
}

func NewSquare() Shape {
	gTotal++
	return Square{id: gTotal}
}

func (s Square) Draw() {
	fmt.Printf("Square : %d\n", s.id)
}

type Ellipse struct {
	id int64
}

func NewEllipse() Shape {
	gTotal++
	return Ellipse{id: gTotal}
}

func (e Ellipse) Draw() {
	fmt.Printf("Ellipse : %d\n", e.id)
}

type Rectangle struct {
	id int64
}

func NewRectangle() Shape {
	gTotal++
	return Rectangle{id: gTotal}
}

func (r Rectangle) Draw() {
	fmt.Printf("Rectangle : %d\n", r.id)
}

type Factory interface {
	CreateCurvedInstance() Shape
	CreateStraightInstance() Shape
}

type SimpleShapeFactory struct{}

func (s SimpleShapeFactory) CreateCurvedInstance() Shape {
	return NewCircle()
}

func (s SimpleShapeFactory) CreateStraightInstance() Shape {
	return NewSquare()
}

type RobustShapeFactory struct{}

func (r RobustShapeFactory) CreateCurvedInstance() Shape {
	return NewEllipse()
}

func (r RobustShapeFactory) CreateStraightInstance() Shape {
	return NewRectangle()
}

func main() {
	var factory Factory
	factory = SimpleShapeFactory{}
	// factory = RobustShapeFactory{}

	var shapes [3]Shape
	shapes[0] = factory.CreateCurvedInstance()
	shapes[1] = factory.CreateStraightInstance()
	shapes[2] = factory.CreateCurvedInstance()

	for i := 0; i < 3; i++ {
		shapes[i].Draw()
	}

}
