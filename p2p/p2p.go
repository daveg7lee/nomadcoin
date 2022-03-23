package p2p

import (
	"fmt"
	"net/http"

	"github.com/daveg7lee/nomadcoin/utils"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{}

func Upgrade(w http.ResponseWriter, r *http.Request) {
	upgrader.CheckOrigin = func(r *http.Request) bool {
		return true
	}
	_, err := upgrader.Upgrade(w, r, nil)
	utils.HandleErr(err)

}

func AddPeer(address, port string) {
	url := fmt.Sprintf("ws://%s:%s/ws", address, port)
	conn, _, err := websocket.DefaultDialer.Dial(url, nil)
	utils.HandleErr(err)

	initPeer(conn, address, port)

}
