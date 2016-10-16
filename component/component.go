package component

type Props map[string]interface{}

type Node struct {
	Name     string
	Props    Props
	Children []Node
}

func New(name string, props Props, children ...Node) Node {
	return Node{Name: name, Props: props, Children: children}
}

func Div(props Props, children ...Node) Node {
	return New("Div", props, children...)
}

func Network(props Props, children ...Node) Node {
	return New("Network", props, children...)
}

func Container(props Props, children ...Node) Node {
	return New("Container", props, children...)
}
