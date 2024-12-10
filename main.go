package main

import (
	"fmt"
	"math"
)

// Конструктор NewPoint создаёт новые точки
type Point struct {
	x float64
	y float64
}

// Функция для создания новой точки
func createPoint(x, y float64) *Point {
	return &Point{x: x, y: y}
}

// Функция для возвращения данных о координате X
func (p *Point) GetX() float64 {
	return p.x
}

// Функция для возвращения данных о координате Y
func (p *Point) GetY() float64 {
	return p.y
}

// Функция для вычисления расстояния между двумя точками
func calculateDistance(point1, point2 *Point) float64 {
	x_diff := point1.GetX() - point2.GetX()
	y_diff := point1.GetY() - point2.GetY()
	return math.Sqrt(x_diff*x_diff + y_diff*y_diff)
}

func main() {
	var x1, y1, x2, y2 float64

	fmt.Print("Введите координаты x первой точки: ")
	fmt.Scan(&x1)
	fmt.Print("Введите координаты y первой точки: ")
	fmt.Scan(&y1)

	point1 := createPoint(x1, y1)

	fmt.Print("Введите координаты x второй точки: ")
	fmt.Scan(&x2)
	fmt.Print("Введите координаты y второй точки: ")
	fmt.Scan(&y2)
	point2 := createPoint(x2, y2)

	distance := calculateDistance(point1, point2)
	fmt.Printf("Расстояние между точками: %.2f\n", distance)
}
