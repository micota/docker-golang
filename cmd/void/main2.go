package main

import (
	"fmt"
	"time"
)

const PI = 3.14

const (
	Username = "user"
	Password = "pass"
)

func init() {
	fmt.Println(time.Now())
}

func main() {
	// var (
	// 	i int = 1
	// 	f64 float64 = 1.7
	// 	t, f bool = true, false
	// )
	// fmt.Println(i, f64, t, f)

	var (
		i    int
		f64  float64
		t, f bool
	)
	fmt.Println(i, f64, t, f)

	xi := 1
	xf64 := 1.2
	fmt.Println(xi, xf64)
	fmt.Printf("%T\n", xf64)

	fmt.Println(PI, Username, Password)
}
