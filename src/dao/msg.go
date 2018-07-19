package dao

import "model"

func GetOfflineMsgs(id string) ([]model.Msg, error) {
	ret := make([]model.Msg, 10)
	return ret, nil
}

func DeleteMsg(msg *model.Msg) error {
	return nil
}
