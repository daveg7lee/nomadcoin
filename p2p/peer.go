package p2p

import (
	"fmt"

	"github.com/gorilla/websocket"
)

var Peers map[string]*peer = make(map[string]*peer)

type peer struct {
	conn  *websocket.Conn
	inbox chan []byte
}

func (p *peer) read() {
	for {
		_, message, err := p.conn.ReadMessage()
		if err != nil {
			break
		}
		fmt.Printf("%s", message)
	}
}

func (p *peer) write() {
	for {
		message := <-p.inbox
		p.conn.WriteMessage(websocket.TextMessage, message)
	}
}

func initPeer(conn *websocket.Conn, address, port string) *peer {
	p := &peer{
		conn,
		make(chan []byte),
	}
	key := fmt.Sprintf("%s:%s", address, port)

	go p.read()
	go p.write()

	Peers[key] = p

	return p
}
