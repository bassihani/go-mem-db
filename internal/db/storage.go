package db

// primary interface of the database
type Storage interface {
	Save(key string, value any) error
	Fetch(key string) (any, bool)
	Delete(key string) error
}
