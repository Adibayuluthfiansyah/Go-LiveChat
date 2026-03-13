package websocket

import (
	"log"
	"time"

	"github.com/Adibayuluthfiansyah/Go-LiveChat/internal/domain"
	"github.com/gorilla/websocket"
)

const (
	writeWait      = 10 * time.Second
	pongWait       = 60 * time.Second
	pingPeriod     = (pongWait * 9) / 10
	maxMessageSize = 512
)

type Client struct {
	hub      *Hub
	conn     *websocket.Conn
	Send     chan *domain.Message
	StreamID string
	UserID   string
}

// read pump
func (c *Client) ReadPump() {
	defer func() {
		c.hub.unregister <- c
		c.conn.Close()
	}()
	c.conn.SetReadLimit(maxMessageSize)
	c.conn.SetReadDeadline(time.Now().Add(pongWait))
	c.conn.SetPongHandler(func(string) error { c.conn.SetReadDeadline(time.Now().Add(pongWait)); return nil })
	for {
		var payload struct {
			Content string `json:"content"`
		}
		err := c.conn.ReadJSON(&payload)
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("error: %v", err)
			}
			break
		}
		msg := &domain.Message{
			StreamID: c.StreamID,
			UserID:   c.UserID,
			Content:  payload.Content,
		}
		c.hub.broadcast <- msg
	}
}

// write pump
func (c *Client) WritePump() {
	ticker := time.NewTicker(pingPeriod)
}
