package union

type Msg interface {
	msgType()
}

type MsgNone struct{}

func (msg MsgNone) msgType() {}

type MsgIncr int

func (msg MsgIncr) msgType() {}

type MsgDecr int

func (msg MsgDecr) msgType() {}
