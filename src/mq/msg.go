package mq

// Msg is interface of message
type Msg interface {
	GetType() int
	GetAttachments() []interface{}
}

// MsgImpl implements Msg
type MsgImpl struct {
	msgType     int
	attachments []interface{}
}

// NewMsg is to create new Msg instance
func NewMsg(msgType int, att []interface{}) Msg {
	msg := MsgImpl{msgType, att}
	return &msg
}

// GetType returns Msg type
func (msg *MsgImpl) GetType() int {
	return msg.msgType
}

// GetAttachments returns attachments binded in Msg
func (msg *MsgImpl) GetAttachments() []interface{} {
	return msg.attachments
}
