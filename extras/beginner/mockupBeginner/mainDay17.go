package main

//https://echo.labstack.com/

import (
	//"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

//No#2
type User struct {
	Firstname string `json:"firstName"` //Camelcase variable
	Lastname  string `json:"last_name"` //snake case vairable
	Age       int    `json:"age"`
}


var Users []User //No#3 Global variable
	


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
	e.Logger.Fatal(e.Start(":1323"))
}
