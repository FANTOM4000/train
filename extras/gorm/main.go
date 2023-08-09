package main

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type User struct {
	Id       int    `gorm:"id"`
	Username string `gorm:"username"`
	Password string `gorm:"password"`
}

func main() {
	client := initDB()

	//insert user
	newUser1 := User{Username: "Thunder", Password: "1234"}
	client.Create(&newUser1)
	newUser2 := User{Username: "Thor", Password: "0213"}
	client.Create(&newUser2)

	//get all users
	allUsers := []User{}
	client.Find(&allUsers)
	fmt.Println(allUsers)

	//get user by id
	user1 := User{Id: newUser1.Id}
	client.First(&user1)
	fmt.Println(user1)

	//get user by username
	user2 := User{}
	client.First(&user2,"username = ?",newUser2.Username)
	fmt.Println(user2)

	//update user
	newUser1.Password = "4321"
	client.Updates(&newUser1)

	//get user by id
	client.First(&user1)
	fmt.Println(user1)

	//delete user
	client.Delete(newUser1)
	client.Delete(newUser2)
}

func initDB() *gorm.DB {
	dsn := "host=ep-purple-wave-845613.ap-southeast-1.aws.neon.tech user=ratthapong400 password=TybZng0s8VtR dbname=neondb port=5432 "
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return db
}
