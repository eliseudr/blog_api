package db

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Open a connection to the MySQL database (localhost)
func Open(user, pass, host, port, name string) (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, pass, host, port, name)
	return gorm.Open(mysql.Open(dsn), &gorm.Config{})
}

// Open a connection to the MySQL server (localhost)
func OpenServer(user, pass, host, port string) (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/?charset=utf8mb4&parseTime=True&loc=Local", user, pass, host, port)
	return gorm.Open(mysql.Open(dsn), &gorm.Config{})
}
