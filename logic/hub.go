package logic

// Hub maintains the set of active clients and broadcasts messages to the
// clients.
type Hub struct {
	// Registered clients.
	clients map[string]*Client

	// Inbound messages from the clients.
	msgChan chan *Msg

	// Register requests from the clients.
	register chan *Client

	// Unregister requests from clients.
	unregister chan string
}

func NewHub() *Hub {
	return &Hub{
		msgChan:    make(chan *Msg),
		register:   make(chan *Client),
		unregister: make(chan string),
		clients:    make(map[string]*Client),
	}
}

func (h *Hub) Send(msg *Msg) {
	h.msgChan <- msg
}

func (h *Hub) Run() {
	for {
		select {
		case client := <-h.register:
			h.clients[client.identifyCode] = client
		case uid := <-h.unregister:
			if client, ok := h.clients[uid]; ok {
				delete(h.clients, uid)
				close(client.send)
			}
		case message := <-h.msgChan:
			if client, exist := h.clients[message.IdentifyCode]; exist {
				select {
				case client.send <- message.Msg:
				default:
					close(client.send)
					delete(h.clients, message.IdentifyCode)
				}
			}
		}
	}
}
