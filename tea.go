package tea

import (
	"context"
	"errors"
	"fmt"
	"reflect"
)

type Msg interface{}
type Model interface{}

type Cmd []func() Msg

func Tag(cmd Cmd, tag func(Msg) Msg) Cmd {
	for i, fn := range cmd {
		if fn == nil {
			continue
		}
		cmd[i] = func() Msg {
			return tag(fn())
		}
	}
	return cmd
}

type Init func() (Model, Cmd)
type Update func(Msg, Model) (Model, Cmd)
type View func(Model) string

type Program struct {
	init   Init
	update Update
	view   View

	context context.Context
	cancel  func()
}

func NewProgram(init Init, update Update, view View) (*Program, error) {
	if init == nil || update == nil || view == nil {
		return nil, errors.New("init or update or view cannot be nill")
	}
	ctx, cancel := context.WithCancel(context.Background())
	program := &Program{
		init:    init,
		update:  update,
		view:    view,
		context: ctx,
		cancel:  cancel,
	}
	return program, nil
}

func execCmd(ch chan Msg, cmd Cmd) {
	for _, fn := range cmd {
		if fn != nil {
			ch <- fn()
		}
	}
}

func (program *Program) Start() chan Msg {
	ch := make(chan Msg)
	go func() {
		model, cmd := program.init()
		fmt.Println(program.view(model))
		go execCmd(ch, cmd)

		for {
			select {
			case msg := <-ch:
				newModel, cmd := program.update(msg, model)
				if !reflect.DeepEqual(newModel, model) {
					model = newModel
					fmt.Println(program.view(model))
				}
				go execCmd(ch, cmd)
			case <-program.context.Done():
				return
			}
		}
	}()
	return ch
}

func (program *Program) Stop() {
	program.cancel()
}
