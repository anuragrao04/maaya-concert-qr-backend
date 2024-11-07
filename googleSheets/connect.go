package googleSheets

import (
	"context"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/option"
	"google.golang.org/api/sheets/v4"
)

var SRV *sheets.Service
var CTX = context.Background()

const (
	SpreadsheetID = "1IJDAr1kV0LzC2koJCYhB0s6Gx1gn0lfMOcJ1A4N-ul0"
	SheetID       = 0
)

func Connect() {
	client, err := google.FindDefaultCredentials(CTX, "https://www.googleapis.com/auth/spreadsheets")
	if err != nil {
		panic("Unable to find default credentials: " + err.Error())
	}

	SRV, err = sheets.NewService(CTX, option.WithCredentials(client))
	if err != nil {
		panic("Unable to retrieve Sheets client: " + err.Error())
	}
	return
}
