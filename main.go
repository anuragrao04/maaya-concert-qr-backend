package main

import (
	"log"

	"github.com/anuragrao04/maaya-concert-qr/creators"
	"github.com/anuragrao04/maaya-concert-qr/database"
	"github.com/anuragrao04/maaya-concert-qr/scanners"
	"github.com/anuragrao04/maaya-concert-qr/senders"
	"github.com/gin-gonic/gin"
)

func main() {
	err := database.Connect("maaya-concert.db")
	if err != nil {
		log.Panic("Error connecting to database:", err)
		return
	}

	router := gin.Default()
	router.POST("/create", creators.CreateUser)
	router.POST("/scan-barcode", scanners.ScanBarcode)
	router.POST("/scan-qr", scanners.ScanQR)
	router.POST("/send-ticket", senders.SendTicket
	router.Run(":6969")
}
