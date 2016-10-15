package tea

import (
	"context"
	"errors"
	"fmt"
)

type Msg interface{}
type Model interface{}

type Cmd struct {
	Task func() Msg
}

type Init struct {
	Model Model
	Cmd   Cmd
}

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
	if update == nil || view == nil {
		return nil, errors.New(" update or view cannot be nill")
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

func (program *Program) Start() chan Msg {
	ch := make(chan Msg)
	go func() {
		model, cmd := program.init.Model, program.init.Cmd
		for {
			select {
			case msg := <-ch:
				model, cmd = program.update(msg, model)
				task := cmd.Task
				if task != nil {
					go func() { ch <- task() }()
				}
				fmt.Println(program.view(model))
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
