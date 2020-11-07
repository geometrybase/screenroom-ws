package main

type Channel struct {
	// Registered clients.
	clients map[*Client]bool

	// Inbound messages from the clients.
	broadcast chan *BroadcastMsg

	// Register requests from the clients.
	register chan *Client

	// Unregister requests from clients.
	unregister chan *Client
}

type BroadcastMsg struct {
	Data []byte
	From *Client
}

func newChannel() *Channel {
	return &Channel{
		broadcast:  make(chan *BroadcastMsg),
		register:   make(chan *Client),
		unregister: make(chan *Client),
		clients:    make(map[*Client]bool),
	}
}

func (c *Channel) run() {
	for {
		select {
		case client := <-c.register:
			c.clients[client] = true
		case client := <-c.unregister:
			if _, ok := c.clients[client]; ok {
				delete(c.clients, client)
				close(client.send)
			}
		case message := <-c.broadcast:
			for client := range c.clients {
				if message.From == client {
					continue
				}
				select {
				case client.send <- message.Data:
				default:
					close(client.send)
					delete(c.clients, client)
				}
			}
		}
	}
}
