```gox
type Msg enum {
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
```

```go
type {
	MsgNone struct {}

	MsgIncr struct {
		V1 int
		V2 string
	}

	MsgDecr struct {
		V1 int
	}
}

var msg interface{} = MsgNone{}
var msg interface{} = MsgNone{V1: 10, V2: 20}

switch msg := Msg.(type); msg {
case MsgNone:
case MsgIncr:
	v1, v2 := msg.V1, msg.V2
case MsgDecr:
	v := msg.V1
default:
}

```
