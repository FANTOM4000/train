package userdata


// No#2
type User struct {
	Id int `json:"id"`
	Firstname string `json:"firstName"` //Camelcase variable
	Lastname  string `json:"last_name"` //snake case vairable
	Age       int    `json:"age"`
}


var Users []User //No#3 Global variable
var userIdGenerator int


func(u User) Add(){
	userIdGenerator++
	u.Id = userIdGenerator
	Users = append(Users, u)
}

func(u User) Update()User{
	for i, v := range Users {
		if v.Id == u.Id {
			Users[i] = u
			return u
		}
	}
	return User{}
}

func(u User) UpdateFirstname()User{
	for i, v := range Users {
		if v.Id == u.Id {
			Users[i].Firstname = u.Firstname
			return Users[i]
		}
	}
	return User{}
}
func(u User) UpdateLastname()User{
	for i, v := range Users {
		if v.Id == u.Id {
			Users[i].Lastname = u.Lastname
			return Users[i]
		}
	}
	return User{}
}
func(u User) UpdateAge()User{
	for i, v := range Users {
		if v.Id == u.Id {
			Users[i].Age = u.Age
			return Users[i]
		}
	}
	return User{}
}

func(u User) Getby_id()User{
	for i, v := range Users {
		if v.Id == u.Id {
			return Users[i]
		}
	}
	return User{}
}

func(u User) Delby_id(){
	newitems := []User{}
    for _, v := range Users {
		if v.Id != u.Id {
			newitems = append(newitems, v)
		}
	}
	Users = newitems
}


//----------------------------------------------------------------
//https://stackoverflow.com/questions/37334119/how-to-delete-an-element-from-a-slice-in-golang
func removeDel(items []User, index int) []User {
    newitems := []User{}
	
    for i, v := range items {	//v คือ value-index array 
        if i != index {
            newitems = append(newitems, v)
        }
    }

    return newitems
}
//เปลี่ยนจาก string เป็นชื่อ struct ครับ
//----------------------------------------------------------------

