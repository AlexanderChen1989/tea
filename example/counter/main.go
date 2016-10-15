package main

import (
	"time"

	"github.com/AlexanderChen1989/tea"
	. "github.com/AlexanderChen1989/tea/example/counter/counter"
)

func main() {
	program, err := tea.NewProgram(Init(), Update, View)
	if err != nil {
		panic(err)
	}
	defer program.Stop()

	ch := program.Start()

	ch <- Incr()
	ch <- Incr()
	ch <- Incr()
	ch <- Decr()
	ch <- Decr()
	ch <- Decr()

	time.Sleep(100 * time.Second)
}
