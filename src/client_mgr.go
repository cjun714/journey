package main

import (
	"sync"

	"github.com/gorilla/websocket"
)

type ClientMgr struct {
	sockets map[string]*websocket.Conn
	lock    sync.RWMutex
}

func NewClientMgr() *ClientMgr {
	mgr := new(ClientMgr)
	mgr.sockets = make(map[string]*websocket.Conn)

	return mgr
}

func (mgr *ClientMgr) Set(wxid string, ws *websocket.Conn) {
	mgr.lock.Lock()
	defer mgr.lock.Unlock()
	mgr.sockets[wxid] = ws
}

func (mgr *ClientMgr) Get(wxid string) (*websocket.Conn, bool) {
	mgr.lock.RLock()
	defer mgr.lock.Unlock()
	val, ok := mgr.sockets[wxid]
	return val, ok
}

func (mgr *ClientMgr) Remove(wxid string) {
	mgr.lock.RLock()
	defer mgr.lock.Unlock()
	delete(mgr.sockets, wxid)
}
