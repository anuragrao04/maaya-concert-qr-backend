package scanners

import (
	"github.com/anuragrao04/maaya-concert-qr/tokens"
	"github.com/gin-gonic/gin"
)

type scanQRRequestFormat struct {
	JWT string `json:"prn"`
}

func ScanQR(c *gin.Context) {
	var request scanQRRequestFormat
	err := c.BindJSON(&request)

	if err != nil || request.JWT == "" {
		c.JSON(400, gin.H{
			"message": "Invalid Request",
		})
		return
	}

	// verify jwt
	valid, claims, err := tokens.VerifyQR(request.JWT)

	if err != nil {
		c.JSON(500, gin.H{
			"message": "Error verifying token",
		})
		return
	}

	if !valid {
		c.JSON(401, gin.H{
			"message": "Invalid Token",
		})
		return
	}

	// return the claims if everything went well
	c.JSON(200, claims)
}
