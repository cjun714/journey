package model

const (
	MsgCommon int = iota // 0
	MsgInvite
	MsgAchievement
	MsgGift
)

type Msg struct {
	Type int         `json:"type"` // 0:common;1:invite;2:achievement;3:gift
	Data interface{} `json:"data,string"`
}
