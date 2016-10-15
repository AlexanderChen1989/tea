package main

import (
	"fmt"

	. "github.com/AlexanderChen1989/tea/view"
)

const component = `
  <App />
`

/*

func User(name string, age int, num int = 10, children ...Node) Node {
  var pets []Node

  for i := 0; i < num; i++ {
      pets = append(pets, <Div> Pet {i} </Div>)
  }

  return (
      <Div>
        <Div>
          {children}
        </Div>
        <Div>
          Pets: {pets}
        </Div>
      </Div>
  )
}

func App() Node {
  return (
    <User name="Alex" age=25 >
      <Div>Hello</Div>
    </User>
  )
}
*/

func main() {
	fmt.Println("Hello")
	fmt.Printf("%+v\n", Div(nil, Div(nil, Div(nil))))
}
