package main

//https://echo.labstack.com/

import (
	//"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

// No#2
type User struct {
	Firstname string `json:"firstName"` //Camelcase variable
	Lastname  string `json:"last_name"` //snake case vairable
	Age       int    `json:"age"`
}

var Users []User //No#3 Global variable

//https://stackoverflow.com/questions/37334119/how-to-delete-an-element-from-a-slice-in-golang
// Removes slice element at index(s) and returns new slice
func Remove[T any](Users []T, indexDelInt int) []T {
    return append(Users[:indexDelInt], Users[indexDelInt+1:]...)
}

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	//e.Logger.Fatal(e.Start(":1323")) No.#4

	// No#4 https://echo.labstack.com/docs/binding
	e.POST("/user", func(c echo.Context) error {
		addUser := User{}
		err := c.Bind(&addUser)
		if err != nil {
			return c.String(http.StatusBadRequest, "400: bad request")
		}

		Users = append(Users, addUser)
		return c.String(http.StatusOK, "200: Success!")
	})

	e.GET("/user", func(c echo.Context) error {
		return c.JSON(200, Users)
	})

	//ุุ6. สร้าง api 1 เส้นโดยมี routing ชื่อว่า "/user/:index" Method PUT จากนั้นนำ json ที่แนบมากับ request  binding เข้ากับ instance User ชั่วคราว
	// แล้วนำ instance นั้นไปทำการ update Users ที่ index โดย index จะได้มาจาก param variable ที่ชื่อว่า 'index'
	// response http 200 พร้อมกับ json ของ Users ที่แก้ไขไป
	e.PUT("/user/:index", func(c echo.Context) error {
		indexString := c.Param("index")
		indexInt, err := strconv.Atoi(indexString)
		if err != nil {
			return c.String(http.StatusBadRequest, "400: bad request")
		}

		updateUserByIndex := User{}
		err = c.Bind(&updateUserByIndex)
		if err != nil {
			return c.String(http.StatusBadRequest, "400: bad request")
		}

		Users[indexInt] = updateUserByIndex
		return c.JSON(200, Users[indexInt])
		//return c.String(http.StatusOK, "200: Update PUT-method succeed!")

	})

	// ึ7. สร้าง api 1 เส้นโดยมี routing ชื่อว่า "/user/:index/:field" Method PATCH จากนั้นนำ json ที่แนบมากับ request  binding เข้ากับ instance User ชั่วคราว
	// แล้วนำ instance นั้นไปทำการ update Users ที่ index โดย index จะได้มาจาก param variable ที่ชื่อว่า 'index' โดย update เฉพาะ field กำกำหนดโดย param variable ที่ชื่อว่า 'field'
	// จะเป็นตัวกำหนดว่าจะต้อง update field ไหน
	// response http 200 พร้อมกับ json ของ Users index ที่แก้ไขไป
	e.PATCH("/user/:index/:field", func(c echo.Context) error {
		indexString := c.Param("index")
		indexInt, err := strconv.Atoi(indexString)
		if err != nil {
			return c.String(http.StatusBadRequest, "400: bad request")
		}
		// Bidning data
		updateUserByIndex := User{}
		err = c.Bind(&updateUserByIndex)
		if err != nil {
			return c.String(http.StatusBadRequest, "400: bad request")
		}

		//Verification field comparison
		Field_nameString := c.Param("field")
		if Field_nameString == "firstName" {
			Users[indexInt].Firstname = updateUserByIndex.Firstname
		}
		if Field_nameString == "last_name" {
			Users[indexInt].Lastname = updateUserByIndex.Lastname
		}
		if Field_nameString == "age" {
			Users[indexInt].Age = updateUserByIndex.Age
		}
		return c.JSON(200, Users[indexInt])
		//return c.String(http.StatusOK, "200: Update PATCH-method succeed!")
	})

	// 8. สร้าง api 1 เส้นโดยมี routing ชื่อว่า "/user/:index" Method GET
	// response http 200 พร้อมกับ json ของ Users ที่ index โดย index จะได้มาจาก param variable ที่ชื่อว่า 'index'
	e.GET("user/:index", func(c echo.Context) error {
		indexString := c.Param("index")
		indexInt, err := strconv.Atoi(indexString)
		if err != nil {
			return c.String(http.StatusBadRequest, "400: bad request")
		}

		return c.JSON(200, Users[indexInt])
	})

//Homework#Day18 Tue-July25, 2023 หาวิธีลบ array index โดยระบุ id (index) Http method request: DEL user/:index

	e.DELETE("user/:index", func(c echo.Context) error {
		indexDelString := c.Param("index")
		indexDelInt, err := strconv.Atoi(indexDelString)
		if err != nil {
			return c.String(http.StatusBadRequest, "400: bad request")
		}
		print(indexDelInt)
		// delete
		//https://stackoverflow.com/questions/37334119/how-to-delete-an-element-from-a-slice-in-golang
		// Removes slice element at index(s) and returns new slice
// Removes slice element at index(s) and returns new slice
func Remove[T any](Users []T, indexDelInt int) []T {
    return append(Users[:indexDelInt], Users[indexDelInt+1:]...)
},
		// delete
		return c.String(http.StatusOK, "200: Delete success")
	})

	e.Logger.Fatal(e.Start(":1323"))
}

