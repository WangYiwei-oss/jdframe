package wscore

import (
	"github.com/gorilla/websocket"
)

type WsClient struct {
	conn      *websocket.Conn
	readChan  chan *WsMessage //读队列
	closeChan chan struct{}   //失败队列
}

func NewWsClient(conn *websocket.Conn) *WsClient {
	return &WsClient{conn: conn, readChan: make(chan *WsMessage), closeChan: make(chan struct{})}
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
		w.readChan <- NewWsMessage(t, data)
	}
}
