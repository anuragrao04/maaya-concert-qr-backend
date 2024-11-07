package creators

import (
	"log"

	"github.com/anuragrao04/maaya-concert-qr/backups"
	"github.com/anuragrao04/maaya-concert-qr/database"
	"github.com/anuragrao04/maaya-concert-qr/models"
	"github.com/gin-gonic/gin"
)

type createUserRequestFormat struct {
	PRN      string `json:"prn"`
	SRN      string `json:"srn"`
	Name     string `json:"name"`
	Semester string `json:"semester"`
	Branch   string `json:"branch"`
}

func CreateUser(c *gin.Context) {
	var request createUserRequestFormat
	c.BindJSON(&request)

	newUser := models.User{
		PRN:      request.PRN,
		SRN:      request.SRN,
		Name:     request.Name,
		Semester: request.Semester,
		Branch:   request.Branch,
	}

	log.Println(newUser)

	err := database.CreateUser(&newUser)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "Error creating user",
		})
		return
	}

	go backups.IncrementWriteCount()

	c.JSON(200, gin.H{
		"message": "User Created Successfully",
	})
}
