package main

import (
	"fmt"

	"github.com/AlexanderChen1989/tea/union"
)

func main() {
	var msg union.Msg = union.MsgIncr(10)

	switch msg := msg.(type) {
	case union.MsgNone:
	case union.MsgIncr:
		fmt.Println("Incr", msg)
	case union.MsgDecr:
		fmt.Println("Decr", msg)
	default:
	}
}
