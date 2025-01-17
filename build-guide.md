# Build Guide: Instagram Follow-Unfollow Analyzer

## 1. Manual Data Export

1. Go to **Instagram Settings > Privacy and Security > Download Data**.
2. Request your data, which includes followers and following lists in a JSON format.
3. Once the data is emailed to you, download and place the JSON files in the `data/` directory of the project.

---

## 2. Setting Up the Tool

### Dependencies

- Install Go (Golang) on your system. Follow the [Go installation guide](https://go.dev/dl/).

### Directory Structure

```plaintext
instagram-follow-unfollow/
├── main.go
├── models/
│   └── user.go
├── parser/
│   └── json_parser.go
├── analyzer/
│   └── analyze.go
├── exporter/
│   └── csv_exporter.go
└── data/
    ├── followers.json   # JSON file containing followers
    └── following.json   # JSON file containing following
