// parser/json_parser.go

package parser

import (
	"encoding/json"
	"os"

	"github.com/yourusername/instagram-follow-unfollow/models"
)

// ParseJSON parses a JSON file and returns a slice of Users.
func ParseJSON(filePath string) ([]models.User, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var users []models.User
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&users)
	if err != nil {
		return nil, err
	}

	return users, nil
}
