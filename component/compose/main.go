package main

import (
	"fmt"

	. "github.com/AlexanderChen1989/tea/component"
)

func main() {
	fmt.Println("Hello")
	fmt.Printf("%+v\n", Div(nil, Div(nil, Div(nil))))
}
