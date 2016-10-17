# The Elm Architecture In Go

## The Elm Version

```elm
program
    : { init : (model, Cmd msg)
      , update : msg -> model -> (model, Cmd msg)
      , subscriptions : model -> Sub msg
      , view : model -> Html msg
      } -> Program Never
```

## The Go Version without subscriptions
```go

type Msg interface{}
type Model interface{}
type Cmd []func() Msg

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
```
