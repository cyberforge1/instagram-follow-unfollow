// #analyzer/analyze.go

package analyzer

import "github.com/yourusername/instagram-follow-unfollow/models"

// CompareLists compares two user lists and returns differences.
func CompareLists(followers, following []models.User) (notFollowingBack, notFollowedBack []models.User) {
	followingMap := make(map[string]bool)

	// Create a map for faster lookups of the following list
	for _, user := range following {
		if len(user.StringListData) > 0 {
			followingMap[user.StringListData[0].Value] = true
		}
	}

	// Find people who follow you but you don't follow back
	for _, user := range followers {
		if len(user.StringListData) > 0 {
			if !followingMap[user.StringListData[0].Value] {
				notFollowedBack = append(notFollowedBack, user)
			}
		}
	}

	// Reset the map for the reverse check
	followersMap := make(map[string]bool)
	for _, user := range followers {
		if len(user.StringListData) > 0 {
			followersMap[user.StringListData[0].Value] = true
		}
	}

	// Find people you follow but who don't follow you back
	for _, user := range following {
		if len(user.StringListData) > 0 {
			if !followersMap[user.StringListData[0].Value] {
				notFollowingBack = append(notFollowingBack, user)
			}
		}
	}

	return
}
