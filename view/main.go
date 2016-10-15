package view

type Props map[string]interface{}

type Node struct {
	Name     string
	Props    Props
	Children []Node
}

type Tag func(props Props, children ...Node) Node

func New(name string, props Props, children ...Node) Node {
	return Node{Name: name, Props: props, Children: children}
}

func Div(props Props, children ...Node) Node {
	return New("div", props, children...)
}
