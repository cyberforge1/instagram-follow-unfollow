// exporter/csv_exporter.go

package exporter

import (
	"encoding/csv"
	"os"

	"github.com/yourusername/instagram-follow-unfollow/models"
)

// ExportToCSV exports a list of users to a CSV file.
func ExportToCSV(filePath string, users []models.User) error {
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Write header
	writer.Write([]string{"Username"})

	// Write data
	for _, user := range users {
		writer.Write([]string{user.Username})
	}

	return nil
}
