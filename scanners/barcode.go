package scanners

import (
	"log"

	"github.com/anuragrao04/maaya-concert-qr/database"
	"github.com/anuragrao04/maaya-concert-qr/googleSheets"
	"github.com/anuragrao04/maaya-concert-qr/models"
	"github.com/gin-gonic/gin"
)

type scanBarcodeRequestFormat struct {
	PRN string `json:"prn"`
}

func ScanBarcode(c *gin.Context) {
	var request scanBarcodeRequestFormat
	c.BindJSON(&request)
	var user models.User
	user, err := database.GetUser(request.PRN)
	if err != nil {
		c.JSON(404, gin.H{
			"message": "User not found",
		})
		return
	}
	err = database.SetPresent(&user)
	if err != nil {
		log.Println("Error updating user:", err)
		c.JSON(500, gin.H{
			"message": "Failed to update user",
		})
		return
	}
	go googleSheets.UpdateRowColorByID(user.ID)
	c.JSON(200, gin.H{
		"user": user,
	})
}
