package auth

import (
	"net/http"
	"realtime-chat/internal/db"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func registerHandler(c *gin.Context)
	var user User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid payload"})
		return
	}

	hashed, _ := bcrypt.GenerateFromPassword([]byte(user.Password), 14)
	_, err := db.DB.Exec(c, "INSERT INTO users (username, password) VALUES ($1, $2)", user.Username, string(hashed))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "erro ao registrar"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "registrado com sucesso"})

func loginHandler(c *gin.Context) {
	var user User
	if err := c.ShouldBindJSON(&user); err != nil{
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid payload"})
		return
	}

	var dbUser User
	err := db.DB.QueryRow(c, "SELECT id, password FROM users WHERE username=$1", user.Username).Scan(&dbUser.ID, &dbUser.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "usuário inválido"})
		return
	}
	err = bcrypt.CompareHashAndPassword([]byte(dbUser.Password), []byte(user.Password))
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "senha incorreta"})
		return
	}

	token, _ := GenerateJWT(dbUser.ID, user.Username)
	c.JSON(http.StatusOK, gin.H{"token": token})