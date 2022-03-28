package p2p

import (
	"fmt"
	"net/http"
	"time"

	"github.com/daveg7lee/nomadcoin/utils"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{}

func Upgrade(w http.ResponseWriter, r *http.Request) {
	openPort := r.URL.Query().Get("openPort")
	ip := utils.Splitter(r.RemoteAddr, ":", 0)

	upgrader.CheckOrigin = func(r *http.Request) bool {
		return openPort != "" && ip != ""
	}
	conn, err := upgrader.Upgrade(w, r, nil)
	utils.HandleErr(err)

	peer := initPeer(conn, ip, openPort)

	time.Sleep(2 * time.Second)
	peer.inbox <- []byte("Hello too!")
	peer.inbox <- []byte("Hello 2!")
	peer.inbox <- []byte("Hello 3!")
	peer.inbox <- []byte("Hello 4!")
	peer.inbox <- []byte("Hello 5!")
}

func AddPeer(address, port, openPort string) {
	url := fmt.Sprintf("ws://%s:%s/ws?openPort=%s", address, port, openPort[1:])
	conn, _, err := websocket.DefaultDialer.Dial(url, nil)
	utils.HandleErr(err)

	peer := initPeer(conn, address, port)

	time.Sleep(time.Second)
	peer.inbox <- []byte("Hello!")
}
