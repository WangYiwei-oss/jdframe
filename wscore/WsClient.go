package wscore

import (
	"github.com/gorilla/websocket"
)

type WsClientLabel interface{}
type WsSendStrategy func(WsClientLabel) bool
type ReadCallback func(int, []byte)

type WsClient struct {
	Label        WsClientLabel  //允许用户自定义的字段
	SendStrategy WsSendStrategy //允许用户自定义的发送策略，为true时才会发送
	ReadCallback ReadCallback   //允许用户自定义的读取策略回调函数
	conn         *websocket.Conn
	readChan     chan *WsMessage //读队列
	closeChan    chan struct{}   //失败队列
}

func defaultSendStrategy(l WsClientLabel) bool {
	return true
}

func NewWsClient(conn *websocket.Conn, label WsClientLabel, sendStrategy WsSendStrategy, readCallback ReadCallback) *WsClient {
	var s WsSendStrategy
	if sendStrategy == nil {
		s = defaultSendStrategy
	} else {
		s = sendStrategy
	}
	return &WsClient{
		Label:        label,
		SendStrategy: s,
		ReadCallback: readCallback,
		conn:         conn,
		readChan:     make(chan *WsMessage),
		closeChan:    make(chan struct{}),
	}
}

func (w *WsClient) ReadLoop() {
	for {
		t, data, err := w.conn.ReadMessage()
		if err != nil {
			w.conn.Close()
			WebSocketFactory.Delete(w.conn)
			w.closeChan <- struct{}{}
			break
		}
		w.ReadCallback(t, data)
		//w.readChan <- NewWsMessage(t, data)
		//
	}
}
