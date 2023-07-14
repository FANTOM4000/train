package main

func main(){
	Add(2,3)
}
//go test -v
func Add(n1 int, n2 int)int{
	if n1 == 0 {
		println("this 0")
	}
	return n1+n2
}