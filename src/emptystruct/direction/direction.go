package main

import "fmt"

type Direction interface {
	Anticlockwise() Direction
	Clockwise() Direction
	DisplacementFrom(Point) Point
}

type Left struct{}

func (_ Left) Anticlockwise() Direction {
	return Down{}
}

func (_ Left) Clockwise() Direction {
	return Up{}
}

func (_ Left) DisplacementFrom(p Point) Point {
	return Point{p.X - 1, p.Y}
}

type Down struct{}

func (_ Down) Anticlockwise() Direction {
	return Right{}
}

func (_ Down) Clockwise() Direction {
	return Left{}
}

func (_ Down) DisplacementFrom(p Point) Point {
	return Point{p.X, p.Y + 1}
}

type Up struct{}

func (_ Up) Anticlockwise() Direction {
	return Left{}
}

func (_ Up) Clockwise() Direction {
	return Right{}
}

func (_ Up) DisplacementFrom(p Point) Point {
	return Point{p.X, p.Y - 1}
}

type Right struct{}

func (_ Right) Anticlockwise() Direction {
	return Up{}
}

func (_ Right) Clockwise() Direction {
	return Down{}
}

func (_ Right) DisplacementFrom(p Point) Point {
	return Point{p.X + 1, p.Y}
}

// the top left point is (0,0)
type Point struct {
	X int
	Y int
}

func main() {
	var d Direction
	d = Left{}
	fmt.Printf("direction: %T\n", d)
	d = d.Clockwise()
	fmt.Printf("turn clockwise: %T\n", d)
	p := Point{3, 4}
	fmt.Printf("current point: %v\n", p)
	next := d.DisplacementFrom(p)
	fmt.Printf("next point in direction %T: %v\n", d, next)
}
