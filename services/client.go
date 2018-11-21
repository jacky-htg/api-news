package services

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
	"websocket/models"

	"github.com/adjust/rmq"
	"github.com/gorilla/websocket"
	"github.com/jacky-htg/api-news/config"
)

const (
	writeWait      = 10 * time.Second
	pongWait       = 60 * time.Second
	pingPeriod     = (pongWait * 9) / 10
	maxMessageSize = 1024 * 1024
)

type client struct {
	ws   *websocket.Conn
	send chan []byte
}

func NewsListHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)

	if err != nil {
		log.Println(err)
		return
	}

	c := &client{
		send: make(chan []byte, maxMessageSize),
		ws:   conn,
	}

	h.register <- c
	go c.writePump()
	c.readPump()
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  maxMessageSize,
	WriteBufferSize: maxMessageSize,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

func (c *client) readPump() {
	defer func() {
		h.unregister <- c
		c.ws.Close()
	}()

	c.ws.SetReadLimit(maxMessageSize)
	c.ws.SetReadDeadline(time.Now().Add(pongWait))
	c.ws.SetPongHandler(func(string) error {
		c.ws.SetReadDeadline(time.Now().Add(pongWait))
		return nil
	})

	for {
		_, message, err := c.ws.ReadMessage()
		if err != nil {
			break
		}

		h.broadcast <- string(message)
	}
}

var connection = rmq.OpenConnection("consumer", config.GetString("database.redis.protocol"), config.GetString("database.redis.address"), config.GetInt("database.redis.db"))

func (c *client) writePump() {
	ticker := time.NewTicker(pingPeriod)

	defer func() {
		ticker.Stop()
		c.ws.Close()
	}()

	for {
		select {
		case _, ok := <-c.send:
			if !ok {
				c.write(websocket.CloseMessage, []byte{})
				return
			}
			/*if err := c.write(websocket.TextMessage, message); err != nil {
				return
			}*/
		case <-ticker.C:
			if err := c.write(websocket.PingMessage, []byte{}); err != nil {
				return
			}
		}

		queue := connection.OpenQueue("news")
		queue.StartConsuming(1000, 500*time.Millisecond)
		queue.AddConsumerFunc("queue", func(delivery rmq.Delivery) {
			m := models.MessageStruct{}
			if err := json.Unmarshal([]byte(delivery.Payload()), &m); err != nil {
				// handle error
				delivery.Reject()
				return
			}

			if err := c.write(websocket.TextMessage, []byte(fmt.Sprintf("%v", m))); err != nil {
				return
			}

			delivery.Ack()
		})
	}
}

func (c *client) write(mt int, message []byte) error {
	c.ws.SetWriteDeadline(time.Now().Add(writeWait))
	return c.ws.WriteMessage(mt, message)
}
