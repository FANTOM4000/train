package main

import (
	"fmt"
	"strings"
)

func main()  {
	fmt.Println("----Day16: review-answer#1----")
	fmt.Println(Plus(2,3))

	fmt.Println("----Day16: review-answer#2----")
	a, b := verifyOddEven(17)
	fmt.Println(a,b)
	a, b = verifyOddEven(2)
	fmt.Println(a,b)

	fmt.Println("----Day16: review-answer#3----")
	fmt.Println("The average of ", []int{2,3,4,5}, " is", Averagesum([]int{2,3,4,5}))

	fmt.Println("----Day16: review-answer#4----")
	phrase1, phrase2 := defineUpperandLower("Orc is October")
	fmt.Println(phrase1, phrase2)

	fmt.Println("----Day16: review-answer#5----")
	fmt.Println(findingMax([]int{2,3,4}))

}

//b1.js โปรแกรมสำหรับคำนวณผลบวกของสองจำนวนที่ผู้ใช้งานป้อนเข้ามา 2 และ 3 ได้ 5
func Plus(x int, y int) int  {
	return x + y
}

//b2.js โปรแกรมที่รับตัวเลขจำนวนเต็มจากผู้ใช้และตรวจสอบว่าตัวเลขนั้นเป็นเลขคู่หรือเลขคี่ 9 คือเลขคี่
func verifyOddEven(n int) (string,bool)  {
	if n%2==0 {
		return "Even number", true
	}
	return "Odd number", false
}

//b3.js โปรแกรมสำหรับคำนวณค่าเฉลี่ยของลิสต์ที่มีข้อมูลตัวเลข 2,3,4,5 คือ 7
func Averagesum(a []int) float64  {
	totalsum := 0
	for i := 0; i < len(a); i++ {
        totalsum += a[i]
    }
	return float64(totalsum) / float64(len(a))
	
}

//b4.js โปรแกรมที่รับอินพุตเป็นอักขระ
// b4.js แล้วแสดงผลลัพธ์เป็นตัวพิมพ์ใหญ่หรือตัวพิมพ์เล็กขึ้นอย่างละหนึ่งตัว a แสดง a และ A
func defineUpperandLower(wording string) (string,string) {

	resultLower := strings.ToLower(wording)
	resultUpper := strings.ToUpper(wording)

	return resultLower, resultUpper

}

// b5.js โปรแกรมสำหรับหาตัวเลขที่มากที่สุดในลิสต์ที่กำหนด 2,3,4 คือ 4
func findingMax(m []int) int  {
	max := m[0]
	for i := 0; i < len(m); i++ {
        if m[i] > max {
            max = m[i]
        }
    }
	return max
}