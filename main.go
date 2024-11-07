package main

import (
	"log"
	"os"

	"github.com/anuragrao04/maaya-concert-qr/creators"
	"github.com/anuragrao04/maaya-concert-qr/database"
	"github.com/anuragrao04/maaya-concert-qr/googleSheets"
	"github.com/anuragrao04/maaya-concert-qr/mailers"
	"github.com/anuragrao04/maaya-concert-qr/scanners"
	"github.com/anuragrao04/maaya-concert-qr/senders"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	// load envs
	err := godotenv.Load()
	if err != nil {
		log.Println(err)
		panic(err)
	}
	err = database.Connect("maaya-concert.db")

	if err != nil {
		log.Panic("Error connecting to database:", err)
		return
	}

	// connect to email server
	mailers.Connect()

	// connect to sheets API
	googleSheets.Connect()

	// mkdir the tempTickets folder if it doesn't exist
	os.Mkdir("tempTickets", 0777)

	router := gin.Default()
	router.POST("/create", creators.CreateUser)
	router.POST("/scan-barcode", scanners.ScanBarcode)
	router.POST("/scan-qr", scanners.ScanQR)
	router.POST("/send-ticket", senders.SendTicket)
	router.GET("/populate-sheet", googleSheets.PopulateSheetWithDBValues)
	router.Run(":6969")
}
