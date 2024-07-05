package main

import (
	"fmt"
	"time"
	"os/user"
)

func init(){
	fmt.Println(time.Now())
}

func main() {
	fmt.Println("Hello World")
	fmt.Println(user.Current())
}
