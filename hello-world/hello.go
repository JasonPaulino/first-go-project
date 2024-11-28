package main

import (
	"fmt"
)

func main() {
	fmt.Println(Hello("Chris"))
}

func Hello(name string) string {
	return "Hello, " + name
}

func IsLegalAge(age int) bool {
	return age >= 18
}
