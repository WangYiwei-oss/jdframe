package wscore

import (
	"github.com/WangYiwei-oss/jdframe/src/configs"
	"github.com/gorilla/websocket"
	"log"
	"sync"
	"time"
)

var WebSocketFactory *ClientMap

func init() {
	WebSocketFactory = newClientMap()
	configs.GetLogger("Global").Info("WebSocket Core: start heartbeat")
	go func() {
		for {
			WebSocketFactory.HeartBeat()
			time.Sleep(2 * time.Second)
		}
	}()
}

type ClientMap struct {
	data map[string]map[string]*WsClient //key是类型 v是map（key是ip，v是Conn）
	lock *sync.RWMutex
}

func newClientMap() *ClientMap {
	return &ClientMap{
		data: make(map[string]map[string]*WsClient),
		lock: new(sync.RWMutex),
	}
}

func (c *ClientMap) Store(class string, conn *websocket.Conn, label WsClientLabel, sendStrategy WsSendStrategy, readCallback ReadCallback) {
	c.lock.Lock()
	defer c.lock.Unlock()
	wsClient := NewWsClient(conn, label, sendStrategy, readCallback)
	if conns, ok := c.data[class]; ok {
		conns[conn.RemoteAddr().String()] = wsClient
	} else {
		newConns := make(map[string]*WsClient)
		newConns[conn.RemoteAddr().String()] = wsClient
		c.data[class] = newConns
	}
	go wsClient.ReadLoop() //处理读循环
}

func (c *ClientMap) Delete(conn *websocket.Conn) {
	c.lock.Lock()
	defer c.lock.Unlock()
	if len(c.data) < 1 {
		log.Println("WebSocket Core: no record")
		return
	}
	for _, conns := range c.data {
		if _, ok := conns[conn.RemoteAddr().String()]; ok {
			delete(conns, conn.RemoteAddr().String())
			return
		}
	}
	log.Println("WebSocket Core: no record")
	return
}

func (c *ClientMap) SendAllClass(class string, v interface{}) {
	if conns, ok := c.data[class]; ok {
		for _, wsClient := range conns {
			err := wsClient.conn.WriteJSON(v)
			if err != nil {
				c.Delete(wsClient.conn)
				log.Println("WebSocket Core:", err, ", Will Delete Conn")
				continue
			}
		}
	} else {
		log.Println("WebSocket Core: no class")
	}
}

func (c *ClientMap) HeartBeat() {
	for _, conns := range c.data {
		for _, wsClient := range conns {
			err := wsClient.conn.WriteMessage(websocket.TextMessage, []byte("ping"))
			if err != nil {
				c.Delete(wsClient.conn)
				continue
			}
		}
	}
}
