/*
Объявите тип 'Color', использовав в качестве базового типа массив из трёх байтов. Считаем, что экземпляр типа 'Color' представляет цвет в формате RGB.
Опишите для типа 'Color' следующие методы:
1. 'Print' выводит информацию цвета на консоль в читаемом виде (например, R=120 G=200 B=45);
2. 'GetR', 'GetG' и 'GetB' возвращают значения компонент цвета (тип возвращаемого значения – byte);
3. 'SetR', 'SetG' и 'SetB' позволяют установить соответствующие компоненты цвета;
4. Метод 'GetBrightness' возвращает яркость цвета, вычисленную по формуле: 0.2126*R + 0.7152*G + 0.0722*B.
Создайте функцию 'maxBrightness' – она получает на вход срез значений Color, а возвращает указатель на элемент из среза, имеющий максимальную яркость.
*/

package main

import "fmt"

// 'Color' represents a color using three bytes for red, green, and blue channels (R, G, B).
type Color [3]byte

// 'Print' function prints the color in the form "R=red, G=green, B=blue".
func (c Color) Print() {
	fmt.Printf("R=%d, G=%d, B=%d\n", c[0], c[1], c[2])
}

// 'GetR', 'GetG', and 'GetB' return the red, green, and blue channels of the color.
func (c Color) GetR() byte {
	return c[0]
}

func (c Color) GetG() byte {
	return c[1]
}

func (c Color) GetB() byte {
	return c[2]
}

// 'SetR', 'SetG', and 'SetB' set the red, green, and blue channels of the color.
func (c *Color) SetR(r byte) {
	if c == nil {
		return
	}

	c[0] = r
}

func (c *Color) SetG(g byte) {
	if c == nil {
		return
	}

	c[1] = g
}

func (c *Color) SetB(b byte) {
	if c == nil {
		return
	}

	c[2] = b
}

// 'GetBrightness' return the brightness of the color.
func (c Color) GetBrightness() float64 {
	return 0.2126*float64(c[0]) + 0.7152*float64(c[1]) + 0.0722*float64(c[2])
}

// 'maxBrightness' returns ptr of the slice element with the maximum brightness.
func maxBrightness(colors []Color) *Color {

	if len(colors) == 0 {
		return nil
	}

	mixIndex := 0
	maxBright := colors[0].GetBrightness()

	for i := 1; i < len(colors); i++ {
		bright := colors[i].GetBrightness()
		if bright > maxBright {
			maxBright = bright
			mixIndex = i
		}
	}

	return &colors[mixIndex]
}

func main() {

	color1 := Color{120, 200, 45}
	color1.Print() // R=120 G=200 B=45

	color2 := Color{255, 255, 255}
	color2.Print() // R=255 G=255 B=255
	color2.SetR(100)
	color2.Print() // R=100 G=255 B=255

	fmt.Printf("Brightness color1: %.2f\n", color1.GetBrightness())
	fmt.Printf("Brightness color2: %.2f\n", color2.GetBrightness())

	// Test func 'maxBrightness'
	colors := []Color{color1, color2}
	brightest := maxBrightness(colors)
	if brightest != nil {
		fmt.Print("Brightest color: ")
		brightest.Print()
		fmt.Printf("Its brightness: %.2f\n", brightest.GetBrightness())
	}
}
