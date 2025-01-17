# structure.md

instagram-follow-unfollow/
├── main.go
├── go.mod
├── models/
│   └── user.go
├── parser/
│   └── json_parser.go
├── analyzer/
│   └── analyze.go
├── exporter/
│   └── csv_exporter.go
└── data/
    ├── followers_and_following/
    │   ├── followers_1.json
    │   ├── following.json
    │   └── pending_follow_requests.json
    └── results/
        ├── not_following_back.csv
        ├── not_followed_back.csv
        └── pending_follow_requests.csv
