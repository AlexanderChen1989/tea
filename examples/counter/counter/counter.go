package counter

import (
	"fmt"

	"github.com/AlexanderChen1989/tea"
)

/*
type Msg enum{
	None,
	Incr(int, string),
	Decr(int),
}

msg := Msg.None
msg := Msg.Incr(10, 20)
msg := Msg.None

match Msg {
case None:
case Incr(v1, v2):
case Decr(v):
default:
}


type {
	MsgNone struct {}

	MsgIncr struct {
		V1 int
		V2 string
	}

	Decr struct {
		V1 int
	}
}

var msg interface{} = MsgNone{}
var msg interface{} = MsgNone{V1: 10, V2: 20}

switch msg := Msg.(type); msg {
case None:
case Incr:
	v1, v2 := msg.V1, msg.V2
case Decr:
	v := msg.V1
default:
}

*/

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
