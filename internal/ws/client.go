package ws

import (
	"log"
)

func (c *Client) Read(hub *Hub) {
	defer func() {
		hub.Unregister <- c
		c.Conn.Close()
	}()

	for {
		_, msg, err := c.Conn.ReadMessage()
		if err != nil {
			log.Println("❌ erro leitura:", err)
			break
		}
		hub.Broadcast <- []byte(c.Username + ": " + string(msg))
	}

	msgStr := string(msg)
	models.SaveMessage(models.Message{
		Username: c.Username,
		Content:  msgStr,
	})
}

func (c *Client) Write() {
	defer c.Conn.Close()

	for msg := range c.Send {
		err := c.Conn.WriteMessage(1, msg)
		if err != nil {
			log.Println("❌ erro escrita:", err)
			break
		}
	}
}
