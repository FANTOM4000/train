package main

import (
	"encoding/json"
	"fmt"
	"math"
)

type Ladlong struct {
	Name       string `json:"name"`
	Gender     string `json:"gender"`
	IsActive   bool	`json:"isActive"`
	Age        int `json:"age"`
	Latitude   float64 `json:"latitude"`
	Longtitude float64 `json:"longitude"`
}

type FullnameLadtitude struct {
	FirstName       string	`json:"firstname"`
	Gender			string	`json:"gender"`
	Status	    	bool	`json:"status"`
	Age            int        `json:"age"`
	Latitude       float64    `json:"lat"`
	Longtitude     float64    `json:"lng"`
}

func main() {

	ladlong1 := Ladlong{
		Name:       "Nobita",
		Gender:     "Male",
		IsActive:   true,
		Age:        25,
		Latitude:   28.6139,
		Longtitude: 77.1025,
	}
	fmt.Println("\n----No#4.Instance of Ladlong------")
	fmt.Println(ladlong1)

	fmt.Println("\n----No#5.[Array] Instance of Ladlong------")
	ladlong2 := []Ladlong{}
	ladlong2 = append(ladlong2, Ladlong{
		Name:       "Shisuka",
		Gender:     "Female",
		IsActive:   true,
		Age:        20,
		Latitude:   18.6139,
		Longtitude: 17.1025,
	})
	ladlong2 = append(ladlong2, Ladlong{
		Name:       "Gaint",
		Gender:     "Male",
		IsActive:   false,
		Age:        30,
		Latitude:   108.6139,
		Longtitude: 107.1025,
	})
	for i := 0; i < len(ladlong2); i++ {
		fmt.Println(ladlong2[i])
	}
	fmt.Println(ladlong2)

	// fmt.Println("\n----Answer#No1----")
	// answer := Addition(2,3)
	// fmt.Printf("The addition result is %d", answer)

	// fmt.Println("\n----Answer#No3: Finding avarage----")
	// average_answer := Average([]int {2,3,4,5})
	// fmt.Println("The average result of ",[]int{2, 3, 4, 5}, " is ", average_answer)

	// fmt.Println("\n----Answer#No5: Maximum value----")
	// max_valueanswer := MaximumValue([]int{2,3,4})
	// fmt.Println("The maximum value of ",[]int{2, 3, 4}, " is ", max_valueanswer)

	// fmt.Println("\n----Answer#No7: Temperature C' to F'----")
	// fmt.Println(CelciusToFarenheit(37.5))

	// fmt.Println("\n----Answer#No9: Mathematics Power----")
	// powercal := mathPowerresult(2,8)
	// fmt.Println("The power b of base number is : ", powercal)

	ll := Ladlong{}

	jjson := `{
		"name":"John",
		"gender":"male",
		"isActive":true,
		"age": 12,
		"latitude": 98.332,
		"longitude": 87.214
	}`
	err := json.Unmarshal([]byte(jjson), &ll)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(ll)

	answer7 := FullnameLadtitude{}
	jjson7 :=`{
		"firstname":"John",
		"gender":"male",
		"status":true,
		"age": 12,
		"lat": 98.332,
		"lng": 87.214
	}`
	err7 := json.Unmarshal([]byte(jjson7), &answer7)
	if err7 != nil {
		fmt.Println(err7)
	}
	fmt.Println(answer7)

	answer8 := []FullnameLadtitude{}
	jjson8 :=`[
		{
			"firstname":"John",
			"gender":"male",
			"status":true,
			"age": 12,
			"lat": 98.332,
			"lng": 87.214
		},
		{
			"firstname":"Sam",
			"gender":"male",
			"status":false,
			"age": 55,
			"lat": 9.332,
			"lng": 8.214
		}
	]`
	err8 := json.Unmarshal([]byte(jjson8), &answer8)
	if err8 != nil {
		fmt.Println(err8)
	}
	fmt.Println(answer8)
}


// b1.js โปรแกรมสำหรับคำนวณผลบวกของสองจำนวนที่ผู้ใช้งานป้อนเข้ามา 2 และ 3 ได้ 5
func Addition(a int, b int) int {
	return a + b
}

// b3.js โปรแกรมสำหรับคำนวณค่าเฉลี่ยของลิสต์ที่มีข้อมูลตัวเลข 2,3,4,5 คือ 7
func Average(a []int) float64 {
	total := 0
	for i := 0; i < len(a); i++ {
		total += a[i]
	}
	return float64(total) / float64(len(a))
}

// b5.js โปรแกรมสำหรับหาตัวเลขที่มากที่สุดในลิสต์ที่กำหนด 2,3,4 คือ 4
func MaximumValue(x []int) int {
	if len(x) == 0 { //Empty array verify
		return 0
	}

	max := x[0]
	for index := 0; index < len(x); index++ {
		if x[index] > max {
			max = x[index]
		}
	}
	return max
}

// b7.js โปรแกรมสำหรับแปลงอุณหภูมิจากเซลเซียสเป็นฟาเรนไฮต์
func CelciusToFarenheit(tempurature float32) float32 {
	resultFarenheit := (tempurature * 1.8) + 32
	return resultFarenheit
}

// b9.js โปรแกรมสำหรับคำนวณค่าเลขยกกำลังของจำนวนเต็มที่ผู้ใช้ป้อนเข้ามา เช่น 3 ยกกำลัง 2 คือ 9
func mathPowerresult(baseNum int, powerNum int) int {

	return int(math.Pow(float64(baseNum), float64(powerNum)))
}
