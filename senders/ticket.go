package senders

import (
	"log"
	"encoding/base64"

	"github.com/anuragrao04/maaya-concert-qr/database"
	"github.com/anuragrao04/maaya-concert-qr/mailers"
	"github.com/anuragrao04/maaya-concert-qr/models"
	"github.com/anuragrao04/maaya-concert-qr/tokens"
	"github.com/gin-gonic/gin"
)

type sendTicketRequest struct {
	PRN   string `json:"prn"`
	Email string `json:"email"`
}

func SendTicket(c *gin.Context) {
	var req sendTicketRequest
	err := c.BindJSON(&req)
	if err != nil || (req.PRN == "" && req.Email == "") {
		c.JSON(400, gin.H{"error": "Invalid request"})
		return
	}

	var user models.User

	// first preference to PRN
	if req.PRN == "" {
		// find by email
		user, err = database.GetUserByEmail(req.Email)
	} else {
		// some idiots didn't fill the form.
		// So we don't have their email.
		// But they made the payment on pesuacademy
		// They also suddenly decided that outside people are allowed
		// so we either have their prn, or their email
		// That's why this shitty code
		user, err = database.GetUser(req.PRN)
	}
	if err != nil {
		c.JSON(404, gin.H{"error": "User not found"})
		return
	}

	qrFilePath, err := tokens.CreateQR(&user)
	if err != nil {
		log.Println("Error creating QR:", err)
		c.JSON(500, gin.H{"error": "Internal server error"})
		return
	}

	if user.Email == "" {
		// this dude didn't fill the google form
		// we only have his payment ack
		// try sending to srn@pesu.pes.edu
		user.Email = user.SRN + "@pesu.pes.edu"
	}

	err = mailers.SendTicket(qrFilePath, user.Email)
	if err != nil {
		log.Println("Error sending ticket:", err)
		c.JSON(500, gin.H{"error": "Error sending ticket"})
	}

	c.JSON(200, gin.H{"message": "Ticket sent"})
}

func SendTicketToAll(c *gin.Context) {

	users, err := database.GetAllUsers()
	if err != nil {
		log.Println("Error getting users:", err)
		c.JSON(500, gin.H{"error": "Internal server error"})
		return
	}

	for _, user := range users {
		go func(user models.User) {
			qrFilePath, err := tokens.CreateQR(&user)
			if err != nil {
				log.Println("Error creating QR:", err)
				c.JSON(500, gin.H{"error": "Internal server error"})
				return
			}

			if user.Email == "" {
				// this dude didn't fill the google form
				// we only have his payment ack
				// try sending to srn@pesu.pes.edu
				user.Email = user.SRN + "@pesu.pes.edu"
			}

			err = mailers.SendTicket(qrFilePath, user.Email)
			if err != nil {
				log.Printf("Error sending ticket to %s: %s\n", user.Email, err)
			}

		}(user)
	}

	c.JSON(202, gin.H{"message": "Sending tickets"})
}

func GetTicket(c *gin.Context) {
	var req sendTicketRequest
	err := c.BindJSON(&req)
	if err != nil || (req.PRN == "") {
		c.JSON(400, gin.H{"error": "Invalid request"})
		return
	}

	var user models.User

	user, err = database.GetUser(req.PRN)

	if err != nil {
		c.JSON(404, gin.H{"error": "User not found"})
		return
	}

	qrPng, err := tokens.GetQRPng(&user)
	if err != nil {
		log.Println("Error creating QR:", err)
		c.JSON(500, gin.H{"error": "Internal server error"})
		return
	}


	c.JSON(200, gin.H{"image": base64.StdEncoding.EncodeToString(qrPng), "message": "Ticket"})
}
