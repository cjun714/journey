package mq

import (
	"sync"
	"util/log"
)

type handlerFunc func(msg Msg)

var q MsgQueue

var handlers map[int]handlerFunc
var lock sync.Mutex

func init() {
	log.I("init message queue")
	q = newMsgQueue()
	handlers = make(map[int]handlerFunc, 100)
}

// SendMsg is to send a message to queue
func SendMsg(msgType int, args ...interface{}) {
	q.Put(NewMsg(msgType, args))
}

// PopMsg is to pop a mesage from queue
func PopMsg() Msg {
	return q.Pop()
}

// HandleMsgs is to handle all messages
func HandleMsgs() {
	for msg := PopMsg(); msg != nil; msg = PopMsg() {
		lock.Lock()
		h := handlers[msg.GetType()]
		lock.Unlock()

		if h == nil {
			log.E("handler is nil for msessage: ", msg.GetType())
		} else {
			h(msg)
		}
	}
}

// RegisteHandler is to registe handler for specified message type
func RegisteHandler(msgType int, fun handlerFunc) {
	lock.Lock()
	handlers[msgType] = fun
	lock.Unlock()
}
