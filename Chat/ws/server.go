package ws

import (
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)
//var clients = make(map[*websocket.Conn]bool)

type SocketsMap struct{
	clients map[*websocket.Conn]bool
}

func NewSocketsMap() (*SocketsMap, error) {
	SocketsMap := new(SocketsMap)
	SocketsMap.clients = make(map[*websocket.Conn]bool)
	return SocketsMap, nil
}

func(s *SocketsMap) Add(conn *websocket.Conn) {
	s.clients[conn] = true
}

func (s *SocketsMap) Remove(conn *websocket.Conn) {
	s.clients[conn] = false
	delete (s.clients, conn)
}

func (s *SocketsMap) IsConnected(conn *websocket.Conn) bool {
 	if _, ok := s.clients[conn]; ok {
 		return true
	}
	return false
}

var Sockets SocketsMap

type fnWsHandler func(w http.ResponseWriter, r *http.Request) (bool, error)
var mCmdToFunc map[string]fnWsHandler

var broadcast = make(chan Payload)
var upgrader = websocket.Upgrader{}

type Payload struct
{
	Command string `json:"Command"`
	Data string `json:"Data"`
}

func Start() error{

	mCmdToFunc = make(map[string]fnWsHandler)
	mCmdToFunc["JR"] = JoinRoom

	NewSocketsMap()

	http.HandleFunc ("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "index.html")
	})
	http.HandleFunc("/ws", HandleRequest)
	err := http.ListenAndServe(":8088", nil)
	if err != nil {
		log.Fatal("error starting http server : ", err)
		return err
	}
	fmt.Println("Start OK")
	return nil
}

func HandleRequest(w http.ResponseWriter, r *http.Request) {
	//go broadcastMessagesToClients()
	websocket, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal("error upgrading GET request to a websocket :: ", err)
	}

	defer websocket.Close()

	if Sockets.IsConnected(websocket) == false {
		Sockets.Add(websocket)
	}

	for
	{
		var payload Payload
		err := websocket.ReadJSON(&payload)
		if err != nil {
			log.Printf("error occurred while reading message : %v", err)
			//delete(clients, websocket)
			break
		}

		Cmd := payload.Command
		if Cmd == "" {
			log.Printf("Cmd is not evaluated")
			break
		}

		fnWsHandler := mCmdToFunc[Cmd]
		if fnWsHandler == nil {
			log.Printf("fnHandler == nil")
			break
		}

		_, err = fnWsHandler(w, r)
		if err != nil {
			log.Printf("fnHandler == nil")
			break
		}
	}
}

func JoinRoom(w http.ResponseWriter, r *http.Request) (bool, error) {

	return true, nil
}
