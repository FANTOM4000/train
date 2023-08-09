package main

import "fmt"

// "interface" หมายถึงการกำหนดสัญญาว่าตัวแปรหรือฟังก์ชันใดๆ ต้องมีการทำงานบางอย่างเป็นแบบใดกับวัตถุที่มี type ที่ระบุไว้ใน interface นั้น

// ประกาศ interface
type Shape interface {
	Area() float64
	DoubleArea() float64
	TripleArea() float64
}

// ประกาศ struct ที่เป็น instance ของ interface Shape
type Circle struct {
	Radius float64
}

type Triangel struct{
	Base float64
	Height float64
}

func (t Triangel) Area() float64 {
	return 0.5 *(t.Base * t.Height)
}

func (t Triangel) DoubleArea() float64 {
	return ((t.Base*t.Height)/2)*2
}
func (t Triangel) TripleArea() float64 {
	return ((t.Base*t.Height)/2)*3
}


// สร้าง method Area สำหรับ Circle
func (c Circle) Area() float64 {
	return 3.14 * c.Radius * c.Radius
}

func (c Circle) DoubleArea() float64 {
	return (3.14 * c.Radius * c.Radius)*2
}

func (c Circle) TripleArea() float64 {
	return (3.14 * c.Radius * c.Radius)*3
}

// ฟังก์ชันทดสอบ
func PrintArea(s Shape) {
	fmt.Println("พื้นที่ของรูป:", s.Area())
}

type Square struct {
	Height float64
	Width float64
}

func(s Square) Area()float64  {
	return s.Height*s.Width
}

func(s Square) DoubleArea()float64  {
	return (s.Height*s.Width)*2
}

func(s Square) TripleArea()float64  {
	return (s.Height*s.Width)*3
}



func main() {
	circle := Circle{Radius: 5}
	square := Square{Height: 4,Width: 5}
	triangle := Triangel{Base: 5, Height: 8}
	PrintArea(circle)
	PrintArea(square)
	PrintArea(triangle)

}
