package creators

import (
	"log"

	"github.com/anuragrao04/maaya-concert-qr/database"
	"github.com/anuragrao04/maaya-concert-qr/models"
	"github.com/gin-gonic/gin"
)

type createUserRequestFormat struct {
	PRN      string `json:"prn"`
	SRN      string `json:"srn"`
	Email    string `json:"email"`
	Name     string `json:"name"`
	Semester string `json:"semester"`
	Branch   string `json:"branch"`

	// only for outside people
	IsPesticide           bool   `json:"isPesticide"`
	PesticideReferralSRN  string `json:"pesticideReferralSRN"`
	PesticideReferralName string `json:"pesticideReferralName"`
	AadharNumber          string `json:"aadharNumber"`
}

func CreateUser(c *gin.Context) {
	var request createUserRequestFormat
	c.BindJSON(&request)

	newUser := models.User{
		PRN:      request.PRN,
		SRN:      request.SRN,
		Email:    request.Email,
		Name:     request.Name,
		Semester: request.Semester,
		Branch:   request.Branch,

		IsPesticide:           request.IsPesticide,
		PesticideReferralSRN:  request.PesticideReferralSRN,
		PesticideReferralName: request.PesticideReferralName,
		AadharNumber:          request.AadharNumber,
	}

	log.Println(newUser)

	err := database.CreateUser(&newUser)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "Error creating user",
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "User Created Successfully",
	})
}
