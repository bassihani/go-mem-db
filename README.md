go-memdb/
├── cmd/
│   └── server/
│       └── main.go         # Entry point: Initializes the DB and runs the app
├── internal/
│   ├── db/
│   │   ├── storage.go        # The Interface definition (The "Contract")
│   │   ├── engine_map.go   # Implementation 1: Using map[string]*Record
│   │   ├── engine_slice.go # Implementation 2: Using []Record (for ordered data)
│   │   └── database_manager.go     # The "Manager" struct that coordinates history/logging
│   └── models/
│       └── record.go       # Shared Structs (Record, HistoryEntry)
├── pkg/
│   └── utils/              # Helper functions (Time formatting, ID generation)
├── go.mod                  # Go module file
└── Makefile                # Shortcuts for 'go run' and 'go test'