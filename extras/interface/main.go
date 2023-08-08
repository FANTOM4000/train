package main

import "fmt"

// "interface" หมายถึงการกำหนดสัญญาว่าตัวแปรหรือฟังก์ชันใดๆ ต้องมีการทำงานบางอย่างเป็นแบบใดกับวัตถุที่มี type ที่ระบุไว้ใน interface นั้น

// ประกาศ interface
type Shape interface {
	Area() float64
}

// ประกาศ struct ที่เป็น instance ของ interface Shape
type Circle struct {
	Radius float64
}

// สร้าง method Area สำหรับ Circle
func (c Circle) Area() float64 {
	return 3.14 * c.Radius * c.Radius
}

// ฟังก์ชันทดสอบ
func PrintArea(s Shape) {
	fmt.Println("พื้นที่ของรูป:", s.Area())
}

func main() {
	circle := Circle{Radius: 5}
	PrintArea(circle)
}
