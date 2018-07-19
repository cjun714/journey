package main

import (
	"dao"
	"model"
	"net/http"
	"util/log"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{}
var clientMgr = NewClientMgr()

func handleConnections(w http.ResponseWriter, r *http.Request) {
	ws, e := upgrader.Upgrade(w, r, nil)
	if e != nil {
		log.E(e)
	}

	var login model.Login

	e = ws.ReadJSON(&login)
	if e != nil {
		log.E(e)
		ws.Close()
	}
	log.I("received data:", login)

	id := login.WXID
	if id == "" {
		log.E("login wxid is empty")
		ws.Close()
	}

	e = ws.WriteJSON(true) // if login, return true to client
	if e != nil {
		log.E(e)
		ws.Close()
	}

	log.I("registed: ", id)

	clientMgr.Set(login.WXID, ws) // register ws

	sendOfflineMsgs(login.WXID) // send offline message

	for {
		var msg model.Msg
		e := ws.ReadJSON(&msg)
		if e != nil {
			log.E(e)
			clientMgr.Remove(login.WXID)
			break
		}

		log.I("received data:", msg)
	}
}

func handleMessages() {

}

func sendOfflineMsgs(wxid string) {
	offlineMsgs, e := dao.GetOfflineMsgs(wxid)
	if e != nil {
		log.E(e)
		return
	}

	for _, msg := range offlineMsgs {
		wx, ok := clientMgr.Get(wxid)
		if ok {
			e := wx.WriteJSON(msg)
			if e == nil {
				dao.DeleteMsg(&msg)
			}
		}
	}
}
