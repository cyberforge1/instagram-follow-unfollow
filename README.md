# Instagram Follow-Unfollow Analyzer

## Project Summary

The **Instagram Follow-Unfollow Analyzer** is a local tool that helps users analyze their Instagram followers and following lists without violating Instagram’s terms of service. The tool allows you to:
1. **Manually export and import data**: Analyze followers and following information using Instagram’s account data download feature.
2. **Compare and analyze lists**:
   - Identify users who don’t follow you back.
   - Identify users you don’t follow back.
   - Transfer pending follow requests into a structured format.
3. **Output actionable results**: Export results to CSV files for easy manual review and actions, such as unfollowing.

---

## Key Features

- **Manual Data Handling**:
  - Import JSON files downloaded from Instagram.
  - Avoid automation or scraping to comply with Instagram’s policies.
- **Comparison Logic**:
  - Identify discrepancies between your followers and following lists.
  - Parse pending follow requests for easy reference.
- **CSV Export**:
  - Generate CSV files listing:
    - Users you follow but who don’t follow you back.
    - Users who follow you but you don’t follow back.
    - Pending follow requests.
  - Include direct profile links for easy manual review and navigation.
- **No Automation of Account Actions**:
  - You must manually perform any follow/unfollow actions to avoid account bans.

---

## How to Run the Project

1. **Prepare the Environment**:
   - Ensure [Go](https://golang.org/) is installed.
   - Clone or download the project repository.
   - Place JSON files (`followers.json`, `following.json`, `pending_follow_requests.json`) in the `data/followers_and_following` directory.

2. **Run the Project**:
   - Use `go run main.go` or build the executable with `go build -o instagram-follow-unfollow` and run `./instagram-follow-unfollow`.

3. **View Results**:
   - Check the generated CSV files:
     - `not_following_back.csv`: Users you follow but who don’t follow you back.
     - `not_followed_back.csv`: Users who follow you but you don’t follow back.
     - `pending_follow_requests.csv`: Pending follow requests with clickable profile links.

---

## Future Enhancements

1. **GUI Interface**:
   - Add a graphical user interface for easier data import and analysis.
2. **Enhanced Data Analysis**:
   - Include insights like mutual followers or inactive accounts.

---

## License

This project is provided as-is and intended for personal use only. Ensure you follow Instagram's terms of service when using this tool.
