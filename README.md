go-memdb
A modular, in-memory storage engine written in Go. This project demonstrates the use of interfaces, dependency injection, and test-driven development to create a flexible key-value store that supports multiple underlying storage strategies.

🏗 ArchitectureThe project follows a layered architecture to separate concerns between data models, storage logic, and management:Manager (DBManager): The central coordinator that handles input validation, logging/history, and delegates storage operations to an engine.Storage Interface (Storage): A "contract" that defines how any storage engine must behave (Save, Fetch, Delete).Storage Engines:MapEngine: High-performance storage using Go's native map.SliceEngine: Ordered storage using a slice, useful for specific data traversal needs.Models: Shared structures for Record and HistoryEntry.Utils: Shared helpers for ID generation and key validation.🚀 
FeaturesPluggable Engines: Easily switch between Map and Slice implementations at runtime using dependency injection.Operation History: Every action (Set, Get, Remove) is logged with a unique UUID, timestamp, and status.Thread-Safety Ready: Includes a test_race target in the Makefile to detect concurrency issues.
Clean API: Simple Set, Get, and Remove methods.📂 Project StructurePlaintextgo-memdb/

├── cmd/server/main.go         # Entry point: Initializes the DB and runs the app
├── internal/
│   ├── db/
│   │   ├── storage.go         # Interface definition
│   │   ├── engine_map.go      # map[string]*Record implementation
│   │   ├── engine_slice.go    # []Record implementation
│   │   └── database_manager.go# Manager coordinating history and logging
│   └── models/
│       └── record.go          # Shared Structs (Record, HistoryEntry)
├── pkg/utils/                 # Helpers (Time, ID generation)
├── go.mod                     # Go module definition
└── Makefile                   # Build and test shortcuts

(Structure derived from)🛠 Getting StartedPrerequisitesGo 1.25.6 or higher InstallationClone the repository and install dependencies:Bashgo mod download
Makefile CommandsThe project includes a Makefile for common tasks:CommandDescriptionmake buildCompiles the project into bin/go-memdbmake runBuilds and runs the applicationmake testRuns unit tests for the DB and Utilsmake test_raceRuns tests with the Go race detector enabledmake cleanRemoves binaries and clears Go caches💻 Usage ExampleYou can inject different engines into the DBManager depending on your needs:Go// Using the Map Engine
engineMap := db.NewEngineMap()
dbmanager := db.NewDBManager(engineMap)

// Set a value
dbmanager.Set("Emp1", 196000)

// Get a value
val, err := dbmanager.Get("Emp1")
(Example based on)🧪 TestingThe project uses testify for assertions. Tests include table-driven tests for the Slice Engine and Mock-based testing for the Database Manager.To run all tests:Bashmake test