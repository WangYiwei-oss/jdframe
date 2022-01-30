package wscore

import (
	"github.com/gorilla/websocket"
	"log"
	"sync"
)

var WebSocketFactory *ClientMap

func init() {
	WebSocketFactory = newClientMap()
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

func (c *ClientMap) Store(class string, conn *websocket.Conn) {
	c.lock.Lock()
	defer c.lock.Unlock()
	wsClient := NewWsClient(conn)
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
				log.Println("WebSocket Core:", err)
			}
		}
	} else {
		log.Println("WebSocket Core: no class")
	}
}
