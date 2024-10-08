package server

import (
	"dominant/infrastructure/utils/log_utils"
	"encoding/json"
	"github.com/gorilla/websocket"
	"net/http"
	"sync"
	"time"
)

//
// @Author yfy2001
// @Date 2024/8/20 19 36
//

var WSServer = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		//验证
		return true
	},
}

type WebSocket struct {
	Conn   *websocket.Conn
	Done   chan struct{}
	ticker *time.Ticker
	mu     sync.Mutex
}

func NewWebSocket(w http.ResponseWriter, r *http.Request) *WebSocket {
	conn, err := WSServer.Upgrade(w, r, nil)
	if err != nil {
		log_utils.Error.Println(err)
		return nil
	}
	interval := 1000
	return &WebSocket{
		Conn:   conn,
		Done:   make(chan struct{}),
		ticker: time.NewTicker(time.Millisecond * time.Duration(interval)),
	}
}

func (ws *WebSocket) OnMessage(handler func([]byte)) {
	defer close(ws.Done)
	for {
		_, message, err := ws.Conn.ReadMessage()
		if err != nil {
			log_utils.Error.Println("Websocket Read ERROR:", err)
			return
		}
		if handler != nil {
			handler(message)
		}
		log_utils.Info.Printf("Websocket Receive: %s", message)
	}
}

func (ws *WebSocket) PushMessage(data func() []any) {
	for {
		select {
		case <-ws.Done:
			return
		case <-ws.ticker.C:
			ws.mu.Lock()
			payload, _ := json.Marshal(data())
			err := ws.Conn.WriteMessage(websocket.TextMessage, payload)
			ws.mu.Unlock()
			if err != nil {
				log_utils.Error.Println("Websocket Write ERROR:", err)
				return
			}
		}
	}
}

func (ws *WebSocket) Close() {
	ws.mu.Lock()
	defer ws.mu.Unlock()
	ws.Conn.Close()
}
