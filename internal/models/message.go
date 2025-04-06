package models

import (
	"context"
	"realtime-chat/internal/db"
)

type Message struct {
	Username  string
	Content   string
}

func SaveMessage(msg Message) error {
	_, err := db.DB.Exec(context.Background(), `
		INSERT INTO messages (username, content) VALUES ($1, $2)
	`, msg.Username, msg.Content)

	return err
}
