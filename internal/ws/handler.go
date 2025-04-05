package ws

import (
	"log"
	"net/http"
	"realtime-chat/internal/auth"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func WebSocketHandler(c *gin.Context) {
	room := c.Param("room")
	if room == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "sala não especificada"})
		return
	}

	tokenStr := c.Query("token")
	if tokenStr == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "token JWT ausente"})
		return
	}

	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		return auth.JWTKey(), nil
	})
	if err != nil || !token.Valid {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "token inválido"})
		return
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "claims inválidos"})
		return
	}
	username := claims["username"].(string)

	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Println("upgrade erro:", err)
		return
	}

	client := &Client{
		Conn:     conn,
		Username: username,
		Send:     make(chan []byte),
	}

	hub := GetHub(room)
	hub.Register <- client

	go client.Read(hub)
	go client.Write()
}
