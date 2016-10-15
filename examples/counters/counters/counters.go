package counters

import (
	"fmt"

	"github.com/AlexanderChen1989/tea"
	"github.com/AlexanderChen1989/tea/examples/counter/counter"
)

type (
	Model struct {
		CounterA tea.Model
		CounterB tea.Model
	}

	CounterAMsg tea.Msg
	CounterBMsg tea.Msg
)

func CounterA(msg tea.Msg) tea.Msg {
	return CounterAMsg(msg)
}

func CounterB(msg tea.Msg) tea.Msg {
	return CounterBMsg(msg)
}

func Init() (tea.Model, tea.Cmd) {
	modelA, cmdA := counter.Init()
	modelB, cmdB := counter.Init()

	model := Model{modelA, modelB}
	cmd := append(
		tea.Tag(cmdA, CounterA),
		tea.Tag(cmdB, CounterB)...,
	)
	return model, cmd
}

func Update(msg tea.Msg, model tea.Model) (tea.Model, tea.Cmd) {
	m, cmd := model.(Model), tea.Cmd{}

	switch msg.(type) {
	case CounterAMsg:
		m.CounterA, cmd = counter.Update(msg, m.CounterA)
	case CounterBMsg:
		m.CounterB, cmd = counter.Update(msg, m.CounterB)
	default:
	}

	return m, cmd
}

func View(model tea.Model) string {
	m := model.(Model)
	return fmt.Sprintf(
		"A => %s\nB => %s\n",
		counter.View(m.CounterA),
		counter.View(m.CounterB),
	)
}
