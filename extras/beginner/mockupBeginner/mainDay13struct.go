package main

import (
	"fmt"
	"strings"
)


type User1 struct {
	Name string
	Gender string
}

type DefineActve struct{
	Name string
	Gender string
	IsActive bool
}

type AgeCategory struct {
	Name string
	Gender string
	IsActive bool
	Age int
}

//No.4 latitude longitude
type LatitudeLongitude struct {
	Name string
	Gender string
	IsActive bool
	Age int
    Latitude float64
    Longitude float64
}

//No.5


func main()  {
	
	user := User1{
		Name: "Orc Orc",
        Gender: "Male",
	}

	fmt.Println(user.Name)
	fmt.Println(user.Gender)

	userActive1 := DefineActve{
		Name: "Orc Orc",
        Gender: "Male",
        IsActive: true,
	}
	fmt.Println("--------------")
	fmt.Println(userActive1.Name)
	fmt.Println(userActive1.Gender)
	fmt.Println(userActive1.IsActive)

	ageUser := AgeCategory{
		Name: "Orc Orc",
        Gender: "Male",
        IsActive: true,
        Age: 18,
	}
	fmt.Println("-------No.3-------")
	fmt.Println(ageUser.Name)
	fmt.Println(ageUser.Gender)
	fmt.Println(ageUser.IsActive)
	fmt.Println(ageUser.Age)

	latitudeUser := LatitudeLongitude{
		Name: "Orc Orc",
        Gender: "Male",
        IsActive: true,
        Age: 18,
        Latitude: 12.34,
        Longitude: 11.50,
	}
	fmt.Println("-------No.4-------")
	fmt.Println(latitudeUser.Name)
	fmt.Println(latitudeUser.Gender)
	fmt.Println(latitudeUser.IsActive)
	fmt.Println(latitudeUser.Age)
	fmt.Println(latitudeUser.Latitude)
	fmt.Println(latitudeUser.Longitude)

	fmt.Println("-------No.5 Array instance-------")
	ladlong1 := []LatitudeLongitude {latitudeUser}
	fmt.Println(len(ladlong1))
	
	for i := 0; i < len(ladlong1) ; i++ {
		fmt.Println(ladlong1[i].Latitude)
		fmt.Println(ladlong1[i].Longitude)
	}

	ladlong2 := []LatitudeLongitude {
		latitudeUser,
		{
			Name: "Anthony Street",
            Gender: "Female",
            IsActive: false,
            Age: 20,
            Latitude: 110.08,
            Longitude: 200.50,
		},
	}
	fmt.Println(len(ladlong2))
	
	for i := 0; i < len(ladlong2) ; i++ {
		fmt.Println(ladlong2[i].Name)
		fmt.Println(ladlong2[i].Latitude)
		fmt.Println(ladlong2[i].Longitude)
	}

	ladlong3 := []LatitudeLongitude {}
	ladlong3 = append(ladlong3, LatitudeLongitude{
		Name: "Gerstien McHow",
        Gender: "Female",
        IsActive: true,
        Age: 50,
        Latitude: 0.99,
        Longitude: 0.88,
	})
	ladlong3 = append(ladlong3, latitudeUser)

	for i := 0; i < len(ladlong3) ; i++ {
		fmt.Println(ladlong3[i].Name)
		fmt.Println(ladlong3[i].Age)
		fmt.Println(ladlong3[i].Latitude)
	
	}
	// st = append(st, User1{
	// 	Name: "Orc Orc",
    //     Gender: "Male",
	// })
	// st = append(st, user)



	// ii := []int{}
	// ii = append(ii, 1)
	// ss := []string{"1","2"}
	// ss = append(ss, "4")
	// ff := []float64{1.0,2.0}
	// ff = append(ff, 3.5)
//----struct array instance
	// st := []User1{user}
	// st = append(st, User1{
	// 	Name: "Orc Orc",
    //     Gender: "Male",
	// })
	// st = append(st, user)
	// fmt.Println("This is answer for b2.js t verify odd even")
	// fmt.Println(verifyOddEven(10))
	// fmt.Println(verifyOddEven(7))
	
	// fmt.Println("-----------------")
	// fmt.Println("This is answer for b4.js to verify odd even")
	// defineUppercase("Test Golang!")

	// fmt.Println("-----------------")
	// fmt.Println("This is answer for b6.js verify LeapYear")
	// x,z := isLeapYearCalculate (1990)
	// fmt.Println(x,z)

	// fmt.Println("-----------------")
	// fmt.Println("This is answer for b8.js counting alphabets")
	// fmt.Println(countingAlphabets("ascd"))


	// fmt.Println("-----------------")
	// fmt.Println("This is answer for b10.js verify adult")
	// fmt.Println(adultVerification(25))
	// fmt.Println(adultVerification(18))
}

//b2 โปรแกรมที่รับตัวเลขจำนวนเต็มจากผู้ใช้และตรวจสอบว่าตัวเลขนั้นเป็นเลขคู่หรือเลขคี่ 9 คือเลขคี่
func verifyOddEven(n int) string {
	if n%2==0 {
		return "This number is EVEN number."
	}
	return "This number is ODD number."

}

//b4.js โปรแกรมที่รับอินพุตเป็นอักขระ 
//b4.js แล้วแสดงผลลัพธ์เป็นตัวพิมพ์ใหญ่หรือตัวพิมพ์เล็กขึ้นอย่างละหนึ่งตัว a แสดง a และ A
func defineUppercase(words string)  {
	//Uppercase
	result1 := strings.ToUpper(words)
	fmt.Println("A word/phrase convert to upper case is: ", result1)
	//Lowercase
	result2 := strings.ToLower(words)
	fmt.Println("A word/phrase convert to upper case is: ", result2)
		
}

// b6.js โปรแกรมที่รับอินพุตเป็นปี แล้วตรวจสอบว่าเป็นปีอธิกสุรทินหรือไม่ 2012 คือ ใช่
func isLeapYearCalculate(year int) (int, bool)  {
	if year%400 ==0{
		fmt.Printf("\n year %d is a leap year.\n", year)
		return year, true
	}
	if year%100 ==0{
        fmt.Printf("\n year %d is NOT a leap year.\n", year)
        return year, false
    }
	if year%4 ==0{
		fmt.Printf("\n year %d is a leap year.\n", year)
		return year, true
	}
	fmt.Printf("\n year %d is NOT a leap year.\n", year)
	return year,false
}

//b8.js โปรแกรมที่รับข้อความจากผู้ใช้และนับจำนวนตัวอักษรที่มีอยู่ในข้อความ "asdf" 4 ตัว
func countingAlphabets(words string) int { 
	ToCountingProcess := words
	Length_of_Alphabets := len(words)

	fmt.Println("String Counting: ", ToCountingProcess)
	return Length_of_Alphabets
}

//b10.js โปรแกรมที่รับชื่อผู้ใช้และอายุ แล้วแสดงข้อความทักทายพร้อมกับเช็คเงื่อนไขว่า
//b10.js ผู้ใช้เป็นเด็กหรือผู้ใหญ่ ถ้าอายุมากกว่า 20 คือผู้ใหญ่น้อยกว่าคือเด็ก
func adultVerification(age int) string  {
	if age >= 20 {
        return "You are an adult."
    }
    return "You are NOT adult."
}

