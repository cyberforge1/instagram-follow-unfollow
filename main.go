// main.go

package main

import (
	"fmt"
	"log"

	"github.com/yourusername/instagram-follow-unfollow/analyzer"
	"github.com/yourusername/instagram-follow-unfollow/exporter"
	"github.com/yourusername/instagram-follow-unfollow/parser"
)

func main() {
	// Parse followers and following JSON files
	followers, err := parser.ParseJSON("data/followers.json")
	if err != nil {
		log.Fatalf("Failed to parse followers: %v", err)
	}

	following, err := parser.ParseJSON("data/following.json")
	if err != nil {
		log.Fatalf("Failed to parse following: %v", err)
	}

	// Analyze the lists
	notFollowingBack, notFollowedBack := analyzer.CompareLists(followers, following)

	// Export results
	err = exporter.ExportToCSV("not_following_back.csv", notFollowingBack)
	if err != nil {
		log.Fatalf("Failed to export not following back: %v", err)
	}
	fmt.Println("Exported not following back to not_following_back.csv")

	err = exporter.ExportToCSV("not_followed_back.csv", notFollowedBack)
	if err != nil {
		log.Fatalf("Failed to export not followed back: %v", err)
	}
	fmt.Println("Exported not followed back to not_followed_back.csv")
}
