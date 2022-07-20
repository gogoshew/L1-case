package tasks

import (
	"fmt"
	"math"
)

/*
Разработать программу нахождения расстояния между двумя точками,
которые представлены в виде структуры Point с инкапсулированными параметрами x,y и конструктором.
*/

type Point struct {
	x, y float64
}

// newPoint функция создания экзмепляра точки "конструктор"
func newPoint(x, y float64) *Point {
	return &Point{x: x, y: y}
}

// getDistance функиця расчета расстояния между двумя точками
func getDistance(x1, y1, x2, y2 float64) float64 {
	a := x2 - x1
	b := y2 - y1
	return math.Sqrt(a*a + b*b)
}

func Task24() {
	point1 := newPoint(50.0, 5.0)
	point2 := newPoint(100.0, 15.0)
	fmt.Printf("Расстояние между точками = %f\n", getDistance(point1.x, point1.y, point2.x, point2.y))
}
