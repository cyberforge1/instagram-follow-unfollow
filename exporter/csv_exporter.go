// #exporter/csv_exporter.go

package exporter

import (
	"encoding/csv"
	"os"

	"github.com/yourusername/instagram-follow-unfollow/models"
)

// ExportToCSV exports a list of users to a CSV file with a clickable hyperlink.
func ExportToCSV(filePath string, users []models.User) error {
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Write header
	writer.Write([]string{"Profile Link", "Username"})

	// Write data
	for _, user := range users {
		if len(user.StringListData) > 0 {
			stringData := user.StringListData[0]
			writer.Write([]string{
				stringData.Href,
				stringData.Value,
			})
		}
	}

	return nil
}
