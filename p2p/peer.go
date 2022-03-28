package p2p

import (
	"fmt"
	"sync"

	"github.com/gorilla/websocket"
)

type peers struct {
	value map[string]*peer
	m     sync.Mutex
}

var Peers peers = peers{
	value: make(map[string]*peer),
}

type peer struct {
	key     string
	address string
	port    string
	conn    *websocket.Conn
	inbox   chan []byte
}

func AllPeers(p *peers) []string {
	p.m.Lock()
	defer p.m.Unlock()

	var keys []string

	for key := range p.value {
		keys = append(keys, key)
	}

	return keys
}

func (p *peer) close() {
	Peers.m.Lock()
	defer Peers.m.Unlock()

	p.conn.Close()
	delete(Peers.value, p.key)
}

func (p *peer) read() {
	defer p.close()
	for {
		_, message, err := p.conn.ReadMessage()
		if err != nil {
			break
		}
		fmt.Printf("%s", message)
	}
}

func (p *peer) write() {
	defer p.close()
	for {
		message, ok := <-p.inbox
		if !ok {
			break
		}
		p.conn.WriteMessage(websocket.TextMessage, message)
	}
}

func initPeer(conn *websocket.Conn, address, port string) *peer {
	key := fmt.Sprintf("%s:%s", address, port)

	p := &peer{
		key:     key,
		address: address,
		port:    port,
		conn:    conn,
		inbox:   make(chan []byte),
	}

	go p.read()
	go p.write()

	Peers.value[key] = p

	return p
}
