package cache

import "database/sql"

type Cache struct {
	conn *sql.DB
}

func New() *Cache {
	return &Cache{}
}
