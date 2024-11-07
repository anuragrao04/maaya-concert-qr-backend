package googleSheets

import (
	"fmt"
	"log"

	"google.golang.org/api/sheets/v4"
)

func UpdateRowColorByID(ID uint) {
	if SRV == nil {
		log.Printf("Sheets service not initialized. Call Connect() first.")
		return
	}

	readRange := "A:A" // Assuming ID is in column A. If you populated the sheet using populate script, it will be
	resp, err := SRV.Spreadsheets.Values.Get(SpreadsheetID, readRange).Do()
	if err != nil {
		log.Printf("Unable to read sheet data: %v", err)
		return
	}

	var rowNumber int64 = -1
	stringID := fmt.Sprintf("%d", ID)
	for i, row := range resp.Values {
		if len(row) > 0 && row[0] == stringID {
			rowNumber = int64(i + 1)
			break
		}
	}

	if rowNumber == -1 {
		log.Printf("ID '%d' not found in the sheet", ID)
		return
	}

	greenColor := &sheets.Color{
		Red:   0.0,
		Green: 1.0,
		Blue:  0.0,
		Alpha: 1.0,
	}

	colorUpdateRequest := &sheets.Request{
		RepeatCell: &sheets.RepeatCellRequest{
			Range: &sheets.GridRange{
				SheetId:       SheetID,
				StartRowIndex: rowNumber - 1,
				EndRowIndex:   rowNumber,
			},

			Cell: &sheets.CellData{
				UserEnteredFormat: &sheets.CellFormat{
					BackgroundColor: greenColor,
				},
			},
			Fields: "userEnteredFormat.backgroundColor",
		},
	}

	batchUpdateRequest := &sheets.BatchUpdateSpreadsheetRequest{
		Requests: []*sheets.Request{colorUpdateRequest},
	}

	_, err = SRV.Spreadsheets.BatchUpdate(SpreadsheetID, batchUpdateRequest).Context(CTX).Do()
	if err != nil {
		log.Printf("Unable to update sheet: %v", err)
	}
}
