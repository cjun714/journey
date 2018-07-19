package mq

type MsgQueue interface {
	Put(Msg)
	Pop() Msg
}

type msgQueueImpl struct {
	queue chan Msg
}

func (mq *msgQueueImpl) Put(msg Msg) {
	// TODO should support timeout
	(mq.queue) <- msg
}

func (mq *msgQueueImpl) Pop() Msg {
	if len(mq.queue) > 0 {
		return <-mq.queue
	}
	return nil
}

func newMsgQueue() MsgQueue {
	mq := new(msgQueueImpl)
	mq.queue = make(chan Msg, 1000)

	return mq
}
