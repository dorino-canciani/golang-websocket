package ws
/*
import (
	"chat/trace"
	"flag"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"os"
)

const (
	socketBufferSize  = 1024
	messageBufferSize = 256
)

type HandlerInterface interface {
	ServeHTTP(w http.ResponseWriter, req *http.Request)
	JoinRoom()
}

type Handler struct {

}

func NewHandler() (*Handler, error) {

}

func Run(h HandlerInterface) {
	var addr = flag.String("addr", ":8089", "The addr of the application.")
	flag.Parse() // parse the flags

	r := newRoom()
	r.tracer = trace.New(os.Stdout)

	http.Handle("/", &templateHandler{filename: "chat.html"})
	http.Handle("/room", h.JoinRoom)

	// get the room going
	go r.run()

	// start the web server
	log.Println("Starting web server on", *addr)
	if err := http.ListenAndServe(*addr, nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}

func (h *Handler) JoinRoom(w http.ResponseWriter, r *http.Request) error {
	return nil
}


func (h *Handler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	var upgrader = &websocket.Upgrader{ReadBufferSize: socketBufferSize, WriteBufferSize: socketBufferSize}

	socket, err := upgrader.Upgrade(w, req, nil)
	if err != nil {
		log.Fatal("ServeHTTP:", err)
		return
	}
	client := &client{
		socket: socket,
		send:   make(chan []byte, messageBufferSize),
		room:   h.r,
	}

	client.descriptor = client.Encrypt([]byte("test"), "test")

	h.r.join <- client
	defer func() { h.r.leave <- client }()
	go client.write()
	client.read()
}
*/