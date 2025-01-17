// analyzer/analyze.go

package analyzer

import "github.com/yourusername/instagram-follow-unfollow/models"

// CompareLists compares two user lists and returns differences.
func CompareLists(followers, following []models.User) (notFollowingBack, notFollowedBack []models.User) {
	followingMap := make(map[string]bool)

	// Create a map for faster lookups of the following list
	for _, user := range following {
		followingMap[user.Username] = true
	}

	// Find people who follow you but you don't follow back
	for _, user := range followers {
		if !followingMap[user.Username] {
			notFollowedBack = append(notFollowedBack, user)
		}
	}

	// Reset the map for the reverse check
	followersMap := make(map[string]bool)
	for _, user := range followers {
		followersMap[user.Username] = true
	}

	// Find people you follow but who don't follow you back
	for _, user := range following {
		if !followersMap[user.Username] {
			notFollowingBack = append(notFollowingBack, user)
		}
	}

	return
}
