package senders

import (
	"github.com/anuragrao04/maaya-concert-qr/database"
	"github.com/anuragrao04/maaya-concert-qr/mailers"
	"github.com/anuragrao04/maaya-concert-qr/tokens"
	"github.com/gin-gonic/gin"
	"log"
)

type sendTicketRequest struct {
	PRN string `json:"prn"`
}

func SendTicket(c *gin.Context) {
	var req sendTicketRequest
	err := c.BindJSON(&req)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid request"})
		return
	}

	user, err := database.GetUser(req.PRN)
	if err != nil {
		c.JSON(404, gin.H{"error": "User not found"})
		return
	}

	qr, err := tokens.CreateQR(&user)
	if err != nil {
		log.Println("Error creating QR:", err)
		c.JSON(500, gin.H{"error": "Internal server error"})
		return
	}

	err = mailers.SendTicket(qr, user.Email)
	if err != nil {
		log.Println("Error sending ticket:", err)
		c.JSON(500, gin.H{"error": "Error sending ticket"})
	}
}
