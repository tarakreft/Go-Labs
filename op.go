package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("go" + "lang")
	fmt.Println("1+1 =", 1+1)
	fmt.Println("7/3 =", 7/3)
	fmt.Println("7%3 =", 7%3)
	fmt.Println("7.0/3.0 =", 7.0/3.0)
	fmt.Println(true && false)
	fmt.Println(true || false)
	fmt.Println(!true)
	fmt.Println("--------------")
	const s string = "constant"
	fmt.Println(s)
	const n = 500000000
	const d = 3e20 / n
	fmt.Println(3e20)
	fmt.Println(d)
	fmt.Println(int64(d))
	fmt.Println(math.Sin(n))
}
