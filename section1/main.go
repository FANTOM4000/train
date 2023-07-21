package main

import "fmt"

func main() {

	fmt.Println("\nLolo")

	//Section 1-1
	fmt.Println(sect1_1(3))
	fmt.Println(sect1_1(0))

	//Section 1-2
	fmt.Println("this is result Golang for Sect-1.2:")
	fmt.Println(sect1_2(5))
	fmt.Println(sect1_2(0))

}

func sect1_1(n int) int {
	if n == 0 {
		return 1
	}
	return n * sect1_1(n-1)
}

func sect1_2(m int) int {
	if m == 0 {
		return 0
	}
	return m + sect1_2(m-1)
}
