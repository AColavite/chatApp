package ws

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetActiveRoomsHandler(c *gin.Context) {
	rooms := GetActiveRoomsMap()
	c.JSON(http.StatusOK, gin.H{"active_rooms": rooms})
}
