package counter

import (
	"fmt"

	"github.com/AlexanderChen1989/tea"
)

type (
	Model int

	IncrMsg struct{}
	DecrMsg struct{}
)

func Incr() tea.Msg {
	return IncrMsg{}
}

func Decr() tea.Msg {
	return DecrMsg{}
}

func Init() (tea.Model, tea.Cmd) {
	return Model(10), tea.Cmd{}
}

func Update(msg tea.Msg, model tea.Model) (tea.Model, tea.Cmd) {
	m := model.(Model)
	switch msg.(type) {
	case IncrMsg:
		return m + 1, tea.Cmd{}
	case DecrMsg:
		return m - 1, tea.Cmd{}
	default:
		return m, tea.Cmd{}
	}
}

func View(model tea.Model) string {
	return fmt.Sprint("Counter: ", model)
}
