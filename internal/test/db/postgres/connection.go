package postgres

import (
	"github.com/jinzhu/gorm"
)

func NewConnection() (*gorm.DB, error) {
	return gorm.Open("postgres", "your_connection_string_here")
}
