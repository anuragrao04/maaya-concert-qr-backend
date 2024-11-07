package googleSheets

import (
	"fmt"
	"log"

	"github.com/anuragrao04/maaya-concert-qr/models"
	"google.golang.org/api/sheets/v4"
)

func AddUsersToSheet(users []models.User) {
	if SRV == nil {
		log.Printf("Sheets service not initialized. Call Connect() first.")
		return
	}

	var valueRange sheets.ValueRange
	var values [][]interface{}

	// Prepare the data for the spreadsheet
	for _, user := range users {
		values = append(values, []interface{}{
			user.ID,
			user.PRN,
			user.SRN,
			user.Email,
			user.Name,
			user.Semester,
			user.Branch,
			user.IsPesticide,
			user.PesticideReferralSRN,
			user.PesticideReferralName,
			// Exclude IsPresent
		})
	}

	valueRange.Values = values

	writeRange := fmt.Sprintf("Sheet1!A1")

	_, err := SRV.Spreadsheets.Values.Append(SpreadsheetID, writeRange, &valueRange).
		ValueInputOption("USER_ENTERED"). // Important: interprets data as user-entered
		InsertDataOption("INSERT_ROWS").  // Inserts new rows
		Context(CTX).
		Do()

	if err != nil {
		log.Printf("Unable to write data to sheet: %v", err)
	}
}
