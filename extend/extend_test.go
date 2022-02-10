package extend

import (
	"fmt"
	"testing"
)

//基类
type Msg struct {
	msgId   int
	msgType int
}

func (m *Msg) SetId(id int) {
	m.msgId = id
}

func (m *Msg) SetType(t int) {
	m.msgType = t
}

//子类
type WxMsg struct {
	Msg
	msgId int
}

func TestExtend(t *testing.T) {
	t.Run("test0", func(t *testing.T) {
		ws := &WxMsg{}

		//ws.SetId(123)
		ws.SetType(999)

		ws.msgId = 567
		ws.msgType = 666

		fmt.Printf("ws.msgId = %d ; ws.Msg.msgId = %d \n\n", ws.msgId, ws.Msg.msgId)
		fmt.Printf("ws.msgType = %d ; ws.Msg.msgType = %d \n\n", ws.msgType, ws.Msg.msgType)
	})
}
