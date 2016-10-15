package counters

import (
	"fmt"

	"github.com/AlexanderChen1989/tea"
	"github.com/AlexanderChen1989/tea/examples/counter/counter"
)

type (
	Model struct {
		CounterA counter.Model
		CounterB counter.Model
	}

	CounterAMsg tea.Msg
	CounterBMsg tea.Msg
)

func CounterA(msg tea.Msg) CounterAMsg {
	return CounterAMsg(msg)
}

func CounterB(msg tea.Msg) CounterBMsg {
	return CounterBMsg(msg)
}

func Init() tea.Init {
	return tea.Init{Model: Model{0, 10}, Cmd: tea.Cmd{}}
}

func Update(msg tea.Msg, model tea.Model) (tea.Model, tea.Cmd) {
	m := model.(Model)
	switch msg.(type) {
	case CounterAMsg:
		ma, cmd := counter.Update(msg, m.CounterA)
		m.CounterA = ma.(counter.Model)
		return m, cmd
	case CounterBMsg:
		mb, cmd := counter.Update(msg, m.CounterB)
		m.CounterB = mb.(counter.Model)
		return m, cmd
	default:
		return m, tea.Cmd{}
	}
}

func View(model tea.Model) string {
	m := model.(Model)
	return fmt.Sprintf(
		"A => %s\nB => %s\n",
		counter.View(m.CounterA),
		counter.View(m.CounterB),
	)
}
