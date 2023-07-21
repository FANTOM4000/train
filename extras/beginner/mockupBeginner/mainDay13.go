package main

import (
	"fmt"
	"math"
	"strings"
)

func main() {

	fmt.Println(Addition(2, 3))

	//b2.js โปรแกรมที่รับตัวเลขจำนวนเต็มจากผู้ใช้และตรวจสอบว่าตัวเลขนั้นเป็นเลขคู่หรือเลขคี่ 9 คือเลขคี่
	x, c := verifyEvenOdd(13)
	fmt.Println(x, c)
	x, c = verifyEvenOdd(10)
	fmt.Println(x, c)
	x, _ = verifyEvenOdd(8)
	fmt.Println(x, c)
	_, c = verifyEvenOdd(7)
	fmt.Println(x, c)

	//b3.js โปรแกรมที่รับตัวเลขจำนวนเต็มจากผู้ใช้และตรวจสอบว่าตัวเลขนั้นเป็นเลขคู่หรือเลขคี่ 9 คือเลขคี่

	fmt.Println("The average of ", []int{2, 3, 4, 5}, " is ", findAverage([]int{2, 3, 4, 5}))

	defineLowerUpper("Orc is handsome!")

	//b5 Maximum-value
	fmt.Println(findMaximum([]int{2, 3, 4}))

	//b6 calculate Leap's Year

	y, z := isLeapYear(1990)
	fmt.Println(y, z)

	//b7 โปรแกรมสำหรับแปลงอุณหภูมิจากเซลเซียสเป็นฟาเรนไฮต์
	fmt.Println(calCelciustoFarenheit(37.5))

	//b8 โปรแกรมที่รับข้อความจากผู้ใช้และนับจำนวนตัวอักษรที่มีอยู่ในข้อความ "asdf" 4 ตัว
	fmt.Println(countingAlphabets("asdf"))

//b9 โปรแกรมสำหรับคำนวณค่าเลขยกกำลังของจำนวนเต็มที่ผู้ใช้ป้อนเข้ามา เช่น 3 ยกกำลัง 2 คือ 9
	fmt.Println(mathPower(2,3))

//b10 โปรแกรมที่รับชื่อผู้ใช้และอายุ แล้วแสดงข้อความทักทายพร้อมกับเช็คเงื่อนไขว่า
//b10-Condition: ผู้ใช้เป็นเด็กหรือผู้ใหญ่ ถ้าอายุมากกว่า 20 คือผู้ใหญ่น้อยกว่าคือเด็ก
	fmt.Println(adultVerfy(25))
	fmt.Println(adultVerfy(19))

}





//b1.js โปรแกรมสำหรับคำนวณผลบวกของสองจำนวนที่ผู้ใช้งานป้อนเข้ามา 2 และ 3 ได้ 5
func Addition(a int, b int) int {

	return a + b
}

//b2.js โปรแกรมที่รับตัวเลขจำนวนเต็มจากผู้ใช้และตรวจสอบว่าตัวเลขนั้นเป็นเลขคู่หรือเลขคี่ 9 คือเลขคี่
func verifyEvenOdd(n int) (string, bool) {
	if n%2 == 0 {
		return fmt.Sprintf("%d is EVEN number", n), true
	}
	return fmt.Sprintf("%d is ODD number", n), false
	// fmt.Printf("%d is ODD number", n)
}

//b3.js โปรแกรมสำหรับคำนวณค่าเฉลี่ยของลิสต์ที่มีข้อมูลตัวเลข 2,3,4,5 คือ 7
func findAverage(a []int) float64 {
	total := 0
	for i := 0; i < len(a); i++ {
		total += a[i]
	}
	return float64(total) / float64(len(a))

}

//b4.js โปรแกรมที่รับอินพุตเป็นอักขระ
// b4.js แล้วแสดงผลลัพธ์เป็นตัวพิมพ์ใหญ่หรือตัวพิมพ์เล็กขึ้นอย่างละหนึ่งตัว a แสดง a และ A
func defineLowerUpper(phrases string) {

	// Displaying strings
	fmt.Println("B4.js convert to lower-case & upper-case")
	fmt.Printf("A phrase before convert: %v", phrases)
	// Converting all the string into lowercase
	// Using ToLower() function
	res1 := strings.ToLower(phrases)
	// Displaying the results
	fmt.Printf("\nA phrase after convert: %v ", res1)
	// Using ToLower() function
	res2 := strings.ToUpper(phrases)
	// Displaying the results
	fmt.Printf("\nA phrase after convert: %v ", res2)

}

// b5.js โปรแกรมสำหรับหาตัวเลขที่มากที่สุดในลิสต์ที่กำหนด 2,3,4 คือ 4
func findMaximum(x []int) int {
	m := -1
	for i, e := range x {
		if i == 0 || e > m {
			m = e
		}
	}
	fmt.Printf("\nMaximum is %v\n", m)
	return m

}

// โปรแกรมที่รับอินพุตเป็นปี แล้วตรวจสอบว่าเป็นปีอธิกสุรทินหรือไม่ 2012 คือ ใช่
func isLeapYear(year int) (int, bool) {
	if year%400 == 0 {
		fmt.Printf("\nyear %d is leap year.\n", year)
		return year, true
	}
	if year%100 == 0 {
		fmt.Printf("\nyear %d is NOT a leap year.\n", year)
		return year, false
	}
	if year%4 == 0 {
		fmt.Printf("\nyear %d is leap year.\n", year)
		return year, true
	}
	fmt.Printf("\nyear %d is NOT a leap year.\n", year)
	return year, false
}

// b7 โปรแกรมสำหรับแปลงอุณหภูมิจากเซลเซียสเป็นฟาเรนไฮต์
func calCelciustoFarenheit(tempurature float32) float32 {

	Farenheit := (tempurature * 1.8) + 32
	return Farenheit

}

// b8: โปรแกรมที่รับข้อความจากผู้ใช้และนับจำนวนตัวอักษรที่มีอยู่ในข้อความ "asdf" 4 ตัว
func countingAlphabets(words string) int {

	StringtoCount := words
	// Finding the length of the string
	// Using len() function
	WordLength := len(words)
//------------------------------
	// Using RuneCountInString() function
	//length2 := utf8.RuneCountInString(mystr)
//-------------------------------
	// Displaying the length of the string
	fmt.Println("string:", StringtoCount)
	return WordLength

}

//b9 โปรแกรมสำหรับคำนวณค่าเลขยกกำลังของจำนวนเต็มที่ผู้ใช้ป้อนเข้ามา เช่น 3 ยกกำลัง 2 คือ 9
func mathPower(a int, b int) int  {
	//fmt.Printf("The power b of base number is : ")
	return int(math.Pow(float64(a), float64(b)))
}


//b10 โปรแกรมที่รับชื่อผู้ใช้และอายุ แล้วแสดงข้อความทักทายพร้อมกับเช็คเงื่อนไขว่า
//b10-Condition: ผู้ใช้เป็นเด็กหรือผู้ใหญ่ ถ้าอายุมากกว่า 20 คือผู้ใหญ่น้อยกว่าคือเด็ก
func adultVerfy(age int) string  {
	if age >= 20 {
        return "You are adult"
    }
    return "You are NOT adult"
	
}

