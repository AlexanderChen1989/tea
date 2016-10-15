package main

import (
	"time"

	"github.com/AlexanderChen1989/tea"
	c "github.com/AlexanderChen1989/tea/examples/counter/counter"
	cs "github.com/AlexanderChen1989/tea/examples/counters/counters"
)

func main() {
	program, err := tea.NewProgram(cs.Init(), cs.Update, cs.View)
	if err != nil {
		panic(err)
	}
	defer program.Stop()

	ch := program.Start()

	ch <- cs.CounterA(c.Incr())
	ch <- cs.CounterA(c.Incr())
	ch <- cs.CounterB(c.Incr())
	ch <- cs.CounterB(c.Incr())
	ch <- cs.CounterA(c.Decr())
	ch <- cs.CounterA(c.Decr())

	time.Sleep(100 * time.Second)
}
