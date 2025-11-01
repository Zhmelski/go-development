/*
Опишите структуру 'Point' для представления точки на плоскости (две координаты X и Y — вещественные числа).
Создайте метод для структуры Point для установки новых координат точки (изменения X и Y).
Создайте структуру 'PointLabeled' — это точка на плоскости, снабжённая текстовой меткой.
Создайте функцию 'normalize':
Она должна получать на вход срез из точек, причём это могут быть как экземпляры структуры 'Point', так и экземпляры структуры 'PointLabeled' (причём даже вперемешку).
Функция нормализует координаты всех точек в срезе — делает так, чтобы координаты вписывались в единичный квадрат [0, 1]x[0, 1].
Так, минимальная координата X всех точек становится нулём, максимальная — 1, а остальные изменяются пропорционально.
В этой задаче вы можете самостоятельно создавать необходимые (или полезные) для решения задачи интерфейсы и методы.
*/

package main

import "fmt"

// 'Point' represents a point on a plane
type Point struct {
	X, Y float64
}

// 'SetCoordinates' sets new coordinates for a 'Point'
func (p *Point) SetCoordinates(x, y float64) {
	p.X = x
	p.Y = y
}

// 'GetCoordinates' returns the coordinates of a 'Point'
func (p *Point) GetCoordinates() (float64, float64) {
	return p.X, p.Y
}

// 'PointLabeled' represents a point with a text label
type PointLabeled struct {
	Point
	Label string
}

// 'SetCoordinates' sets new coordinates for a 'PointLabeled'
func (pl *PointLabeled) SetCoordinates(x, y float64) {
	pl.X = x
	pl.Y = y
}

// 'GetCoordinates' returns the coordinates of a 'PointLabeled'
func (pl *PointLabeled) GetCoordinates() (float64, float64) {
	return pl.X, pl.Y
}

// 'Normalizer' is an interface for objects that can be normalized.
type Normalizer interface {
	GetCoordinates() (float64, float64)
	SetCoordinates(float64, float64)
}

// 'normalize' normalizes the coordinates of the given slice of points.
func normalize(points []Normalizer) {
	if len(points) == 0 {
		return
	}

	// Find the min and max coordinates
	minX, minY := points[0].GetCoordinates()
	maxX, maxY := minX, minY

	for _, p := range points {
		x, y := p.GetCoordinates()

		if x < minX {
			minX = x
		}
		if x > maxX {
			maxX = x
		}
		if y < minY {
			minY = y
		}
		if y > maxY {
			maxY = y
		}
	}

	// Calculate the range
	rangeX := maxX - minX
	rangeY := maxY - minY

	// Normalize the points
	for _, p := range points {
		x, y := p.GetCoordinates()

		var newX, newY float64
		if rangeX == 0 {
			newX = 0
		} else {
			newX = (x - minX) / rangeX
		}

		if rangeY == 0 {
			newY = 0
		} else {
			newY = (y - minY) / rangeY
		}

		p.SetCoordinates(newX, newY)
	}
}

func printPoints(points []Normalizer) {
	for i, p := range points {
		x, y := p.GetCoordinates()

		if pl, ok := p.(*PointLabeled); ok {
			fmt.Printf("Point %d (label: %s): X=%.4f, Y=%.4f\n", i+1, pl.Label, x, y)
		} else {
			fmt.Printf("Point %d: X=%.4f, Y=%.4f\n", i+1, x, y)
		}
	}
}

func main() {

	p1 := Point{X: 10, Y: 20}
	p2 := Point{X: 30, Y: 40}

	pl1 := PointLabeled{Point: Point{X: 50, Y: 60}, Label: "A"}
	pl2 := PointLabeled{Point: Point{X: 70, Y: 80}, Label: "B"}

	points := []Normalizer{&p1, &p2, &pl1, &pl2}

	fmt.Println("\nBefore normalization:")
	printPoints(points)

	normalize(points)

	fmt.Println("\nAfter normalization:")
	printPoints(points)

}
