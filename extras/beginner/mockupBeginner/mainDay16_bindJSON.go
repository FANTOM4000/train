package main

import (
	"encoding/json"
	"fmt"
)

type LadLongtitude struct {
	Name      string  `json:"name"`
	Gender    string  `json:"gender"`
	IsActive  bool    `json:"isActive"`
	Age       int     `json:"age"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

// 7. สร้าง struct ที่สามารถรองรับ json ดังต่อไปนี้และสร้าง instance และ binding json กับ instance และ print ออกมา
type FullnameWithLadLong struct {
	FirstName string   `json:"firstname"`
	Gender    string   `json:"gender"`
	Status    bool     `json:"status"`
	Age       int      `json:"age"`
	Latitude  float64  `json:"lat"`
	Longitude float64  `json:"lng"`
	Tags      []string `json:"tags"`
}

// 10. สร้าง struct ที่สามารถรองรับ json ดังต่อไปนี้และสร้าง instance และ binding json กับ instance และ print ออกมา
type FirstnameWithInfo struct {
	FirstName string   `json:"firstname"`
	Gender    string   `json:"gender"`
	Status    bool     `json:"status"`
	Age       int      `json:"age"`
	Latitude  float64  `json:"lat"`
	Longitude float64  `json:"lng"`
	Info      Info     `json:"info"`
	Friend    []Info   `json:"friends"`
	Tags      []string `json:"tags"`
}

type Info struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func main() {

	ladlongNo4 := LadLongtitude{
		Name:     "Orc October",
		Gender:   "Male",
		IsActive: true,
		Age:      25,
	}
	fmt.Println("\n---No#4---answer---")
	fmt.Println(ladlongNo4)

	fmt.Println("\n---No#5---answer---")
	ladlongNo5 := []LadLongtitude{}
	ladlongNo5 = append(ladlongNo5, LadLongtitude{
		Name:      "Priew Suaymak",
		Gender:    "Female",
		IsActive:  false,
		Age:       25,
		Latitude:  113.11,
		Longitude: 400.05,
	},
		LadLongtitude{
			Name:      "Nobita",
			Gender:    "Male",
			IsActive:  true,
			Age:       28,
			Latitude:  11.00,
			Longitude: 40.05,
		},
		LadLongtitude{
			Name:      "Giant",
			Gender:    "Male",
			IsActive:  true,
			Age:       30,
			Latitude:  200.05,
			Longitude: 100.88,
		})
	fmt.Println(ladlongNo5)

	//6. นำ struct ที่ได้จากข้อ 5 มาสร้าง instance และ binding json กับ instance และ print ออกมา
	fmt.Println("\n---No#6---answer---")
	ladlongNo6 := LadLongtitude{}
	jjson6 := `{
		"name":"John",
		"gender":"male",
		"isActive":true,
		"age": 12,
		"latitude": 98.332,
		"longitude": 87.214
	}`
	errjson6 := json.Unmarshal([]byte(jjson6), &ladlongNo6)
	if errjson6 != nil {
		fmt.Println(errjson6)
	}
	fmt.Println(ladlongNo6)

	//7. สร้าง struct ที่สามารถรองรับ json ดังต่อไปนี้และสร้าง instance และ binding json กับ instance และ print ออกมา
	fmt.Println("\n---No#7---answer---")
	ladlongNo7 := FullnameWithLadLong{}
	jjson7 := `{
    	"firstname":"John",
    	"gender":"male",
    	"status":true,
    	"age": 12,
    	"lat": 98.332,
    	"lng": 87.214
	}`
	errjson7 := json.Unmarshal([]byte(jjson7), &ladlongNo7)
	if errjson7 != nil {
		fmt.Println(errjson7)
	}
	fmt.Println(ladlongNo7)

	// 8. สร้าง array struct ที่สามารถรองรับ json ดังต่อไปนี้และสร้าง instance และ binding json กับ instance และ print ออกมา
	fmt.Println("\n---No#8---answer---")
	ladlongNo8 := []FullnameWithLadLong{}
	jjson8 := `[
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
	errjson8 := json.Unmarshal([]byte(jjson8), &ladlongNo8)
	if errjson8 != nil {
		fmt.Println(errjson8)
	}
	fmt.Println(ladlongNo8)

	// 9. สร้าง struct ที่สามารถรองรับ json ดังต่อไปนี้และสร้าง instance และ binding json กับ instance และ print ออกมา
	fmt.Println("\n---No#9---answer JSON object in JSON---")
	ladlongNo9 := FullnameWithLadLong{}
	jjson9 := `{
    	"firstname":"John",
    	"gender":"male",
    	"status":true,
    	"age": 12,
    	"lat": 98.332,
    	"lng": 87.214,
    	"tags": [
        	"qui",
        	"consectetur",
        	"consectetur",
        	"esse",
        	"velit",
        	"do",
        	"duis"
    	]}`
	errjson9 := json.Unmarshal([]byte(jjson9), &ladlongNo9)
	if errjson9 != nil {
		fmt.Println(errjson9)
	}
	fmt.Println(ladlongNo9)

	// 10. สร้าง struct ที่สามารถรองรับ json ดังต่อไปนี้และสร้าง instance และ binding json กับ instance และ print ออกมา
	fmt.Println("\n---No#10---answer JSON object in JSON---")
	ladlongNo10 := FirstnameWithInfo{}
	jjson10 := `{
		"firstname":"John",
		"gender":"male",
		"status":true,
		"age": 12,
		"lat": 98.332,
		"lng": 87.214,
		"info":{
			"id": 0,
			"name": "Murphy Ward"
		}
	}`
	errjson10 := json.Unmarshal([]byte(jjson10), &ladlongNo10)
	if errjson10 != nil {
		fmt.Println(errjson10)
	}
	fmt.Println(ladlongNo10)

	//11. สร้าง struct ที่สามารถรองรับ json ดังต่อไปนี้และสร้าง instance และ binding json กับ instance และ print ออกมา
	fmt.Println("\n---No#11---answer JSON object combination advance---")
	ladlongNo11 := FirstnameWithInfo{}
	jjson11 := `{
		"firstname":"John",
		"gender":"male",
		"status":true,
		"age": 12,
		"lat": 98.332,
		"lng": 87.214,
		"info":{
			"id": 0,
			"name": "Murphy Ward"
		},
		"friends": [
			{
				"id": 0,
				"name": "Murphy Ward"
			},
			{
				"id": 1,
				"name": "Maxine Harris"
			},
			{
				"id": 2,
				"name": "Charles Abbott"
			}
		],
		"tags": [
			"qui",
			"consectetur",
			"consectetur",
			"esse",
			"velit",
			"do",
			"duis"
		]
	}`
	errjson11 := json.Unmarshal([]byte(jjson11), &ladlongNo11)
	if errjson11 != nil {
		fmt.Println(errjson11)
	}
	fmt.Println(ladlongNo11)

	//12. สร้าง struct ที่สามารถรองรับ json ดังต่อไปนี้และสร้าง instance และ binding json กับ instance และ print ออกมา
	fmt.Println("\n---No#12---answer JSON object combination advance---")
	ladlongNo12 := FirstnameWithInfo{}
	jjson12 := `{
		"firstname":"John",
		"gender":"male",
		"status":true,
		"age": 12,
		"lat": 98.332,
		"lng": 87.214,
		"friends": [
			{
				"id": 0,
				"name": "Murphy Ward"
			},
			{
				"id": 1,
				"name": "Maxine Harris"
			},
			{
				"id": 2,
				"name": "Charles Abbott"
			}
		]
	}`
	errjson12 := json.Unmarshal([]byte(jjson12), &ladlongNo12)
	if errjson12 != nil {
		fmt.Println(errjson12)
	}
	fmt.Println(ladlongNo12)

}
