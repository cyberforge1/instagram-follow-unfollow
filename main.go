// #main.go

package main

import (
	"fmt"
	"log"
	"os"

	"github.com/yourusername/instagram-follow-unfollow/analyzer"
	"github.com/yourusername/instagram-follow-unfollow/exporter"
	"github.com/yourusername/instagram-follow-unfollow/parser"
)

func main() {
	// Ensure the results directory exists
	resultsDir := "data/results"
	err := os.MkdirAll(resultsDir, os.ModePerm)
	if err != nil {
		log.Fatalf("Failed to create results directory: %v", err)
	}

	// Parse followers and following JSON files
	followers, err := parser.ParseJSON("data/followers_and_following/followers_1.json")
	if err != nil {
		log.Fatalf("Failed to parse followers: %v", err)
	}

	following, err := parser.ParseJSON("data/followers_and_following/following.json")
	if err != nil {
		log.Fatalf("Failed to parse following: %v", err)
	}

	// Analyze the lists
	notFollowingBack, notFollowedBack := analyzer.CompareLists(followers, following)

	// Export "Not Following Back" to a CSV file
	err = exporter.ExportToCSV(resultsDir+"/not_following_back.csv", notFollowingBack)
	if err != nil {
		log.Fatalf("Failed to export 'not following back': %v", err)
	}
	fmt.Println("Exported 'not following back' to data/results/not_following_back.csv")

	// Export "Not Followed Back" to a CSV file
	err = exporter.ExportToCSV(resultsDir+"/not_followed_back.csv", notFollowedBack)
	if err != nil {
		log.Fatalf("Failed to export 'not followed back': %v", err)
	}
	fmt.Println("Exported 'not followed back' to data/results/not_followed_back.csv")

	// Parse pending follow requests JSON file
	pendingFollowRequests, err := parser.ParseJSON("data/followers_and_following/pending_follow_requests.json")
	if err != nil {
		log.Fatalf("Failed to parse pending follow requests: %v", err)
	}

	// Export pending follow requests to a CSV file
	err = exporter.ExportToCSV(resultsDir+"/pending_follow_requests.csv", pendingFollowRequests)
	if err != nil {
		log.Fatalf("Failed to export pending follow requests: %v", err)
	}
	fmt.Println("Exported pending follow requests to data/results/pending_follow_requests.csv")
}
