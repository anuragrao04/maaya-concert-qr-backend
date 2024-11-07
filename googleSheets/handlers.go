package googleSheets

import (
	"github.com/anuragrao04/maaya-concert-qr/database"
	"github.com/gin-gonic/gin"
)

func PopulateSheetWithDBValues(c *gin.Context) {
	users, err := database.GetAllUsers()
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	go AddUsersToSheet(users)
	c.JSON(202, gin.H{"message": "Populating sheet with DB values. This may take some time."})
}
