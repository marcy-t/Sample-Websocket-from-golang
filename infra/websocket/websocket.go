package websocket

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/gorilla/websocket"
	"github.com/pkg/errors"
)

type WebSocketClient struct {
	*websocket.Conn
}

type Interface interface {
	Con()
	Close()
}

var (
	endpoint string = "wss://121-85-7-220f1.hyg1.eonet.ne.jp:3001/ws"
)

func NewClient() (Interface, error) {
	conn, _, err := websocket.DefaultDialer.Dial(endpoint, nil)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	conn.SetReadDeadline(time.Now().Add(time.Second * 300)) // DEBUG

	return &WebSocketClient{conn}, nil
}

// sample
func (w *WebSocketClient) Con() {
	fmt.Println("$$$$$$$$$$$$$$$$$$$$$$")
	fmt.Println("$$$$$$$$$$$$$$$$$$$$$$")

	// websocket : uidなど受け取り
	_, receiveMsg, err := w.ReadMessage()
	if err != nil {
		log.Fatal(err)
	}

	var mapMsg map[string]interface{}
	if err := json.Unmarshal(receiveMsg, &mapMsg); err != nil {
		log.Fatal(err)
	}
	// receive := map[string]interface{}{"uid": mapMsg["uid"], "subscribe": "block"}
	receive := map[string]interface{}{"uid": mapMsg["uid"], "subscribe": "unconfirmedAdded"}

	// 送信
	w.WriteJSON(receive)

	_, receiveMsg01, err := w.ReadMessage()
	if err != nil {
		log.Fatal(err)
	}

	var mapMsg01 map[string]interface{}
	if err := json.Unmarshal(receiveMsg01, &mapMsg01); err != nil {
		log.Fatal(err)
	}

	for _, v := range mapMsg01 {
		fmt.Println("v : ", v)
	}

	fmt.Println("$$$$$$$$$$$$$$$$$$$$$$")
	fmt.Println("$$$$$$$$$$$$$$$$$$$$$$")
}

// Send ...
// func (w *WebSocketClient) Send(msg interface{}) error {
// 	return w.WriteJSON(msg)
// }

// Close ...
func (w *WebSocketClient) Close() {
	w.Conn.Close()
}
