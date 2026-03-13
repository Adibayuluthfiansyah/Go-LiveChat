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
