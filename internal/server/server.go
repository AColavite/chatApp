package server

import (
	"log"
	"net/http"
	"realtime-chat/internal/auth"
	"realtime-chat/internal/common"

	"github.com/gin-gonic/gin"
	"realtime-chat/internal/ws"
)

func Start() error {
	router := gin.Default()

	// 
	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})

	router.GET("/ws/chat/:room", ws.WebSocketHandler)
	router.GET("/api/rooms/active", ws.GetActiveRoomsHandler)
	router.GET("/ws/chat", ws.WebSocketHandler)
	router.GET("/api/rooms/active", ws.GetActiveRoomsHandler)


	// (login e registro)
	auth.RegisterRoutes(router)

	// Grupo protegido por JWT
	private := router.Group("/api")
	private.Use(auth.JWTMiddleware())

	private.GET("/profile", func(c *gin.Context) {
		userID := c.GetInt("user_id")
		username := c.GetString("username")

		c.JSON(http.StatusOK, gin.H{
			"user_id":  userID,
			"username": username,
		})
	})

	// Iniciar servidor
	port := common.GetEnv("PORT", "8080")
	log.Println("🚀 Server running at http://localhost:" + port)
	return router.Run(":" + port)

	ws.InitHub()

}
