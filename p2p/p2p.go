package p2p

import (
	"fmt"
	"net/http"

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

	initPeer(conn, ip, openPort)

}

func AddPeer(address, port, openPort string) {
	url := fmt.Sprintf("ws://%s:%s/ws?openPort=%s", address, port, openPort[1:])
	conn, _, err := websocket.DefaultDialer.Dial(url, nil)
	utils.HandleErr(err)

	initPeer(conn, address, port)
}
