package scanners

import (
	// "github.com/anuragrao04/maaya-concert-qr/googleSheets"
	"github.com/anuragrao04/maaya-concert-qr/database"
	"github.com/anuragrao04/maaya-concert-qr/googleSheets"
	"github.com/anuragrao04/maaya-concert-qr/tokens"
	"github.com/gin-gonic/gin"
)

type scanQRRequestFormat struct {
	JWT string `json:"jwt"`
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
	// googleSheets.UpdateRowColorByPRN(claims)

	userIDFloat, _ := claims["userID"].(float64)
	userID := uint(userIDFloat)

	user, err := database.GetUserByID(userID)
	if err != nil {
		c.JSON(404, gin.H{
			"message": "User not found",
		})
		return
	}

	if user.IsPresent {
		// this user is already marked present
		c.JSON(401, gin.H{
			"user":    user,
			"message": "User already marked present",
		})
		return
	}

	err = database.SetPresent(&user)
	go googleSheets.UpdateRowColorByID(user.ID)

	c.JSON(200, user)
}
