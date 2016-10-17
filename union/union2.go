package union

type Model int

type None struct{}
type Incr int
type Decr int

type CMsg interface {
	Process(Model) Model
}

func (msg None) Process(model Model) Model {
	return model
}

func (msg Incr) Process(model Model) Model {
	return Model(int(model) + int(msg))
}

func (msg Decr) Process(model Model) Model {
	return Model(int(model) - int(msg))
}
