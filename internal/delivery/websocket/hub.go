package websocket

import (
	"sync"

	"github.com/Adibayuluthfiansyah/Go-LiveChat/internal/domain"
)

type Hub struct {
	clients     map[string]map[*Client]bool
	register    chan *Client
	unregister  chan *Client
	broadcast   chan *domain.Message
	chatUsecase domain.ChatUseCase
	mu          sync.RWMutex
}

func NewHub(chatUsecase domain.ChatUseCase) *Hub {
	return &Hub{
		clients:     make(map[string]map[*Client]bool),
		register:    make(chan *Client),
		unregister:  make(chan *Client),
		broadcast:   make(chan *domain.Message),
		chatUsecase: chatUsecase,
	}
}
func (h *Hub) Run() {
	for {
		select {
		case client := <-h.register:
			h.mu.Lock()
			if h.clients[client.StreamID] == nil {
				h.clients[client.StreamID] = make(map[*Client]bool)
			}
			h.clients[client.StreamID][client] = true
			h.mu.Unlock()
		case client := <-h.unregister:
			h.mu.Lock()
			if _, ok := h.clients[client.StreamID]; ok {
				delete(h.clients[client.StreamID], client)
				close(client.Send)
				if len(h.clients[client.StreamID]) == 0 {
					delete(h.clients, client.StreamID)
				}
			}
			h.mu.Unlock()
		case message := <-h.broadcast:
			h.chatUsecase.SendMessage(message.StreamID, message.UserID, message.Content)
			h.mu.Lock()
			connection := h.clients[message.StreamID]
			for client := range connection {
				select {
				case client.Send <- message:
				default:
					close(client.Send)
					delete(h.clients[message.StreamID], client)
				}
			}
			h.mu.Unlock()
		}
	}
}
