package prodinfra

import (
	"errors"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func NewDB() (*gorm.DB, error) {
	dsn := os.Getenv("APP_MYSQL_DSN")
	if dsn == "" {
		return nil, errors.New("APP_MYSQL_DSN not found")
	}
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	return db, err
}

func NewStubDB() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	_ = db.Migrator().DropTable(&AccountPO{})
	err = db.Migrator().AutoMigrate(&AccountPO{})
	if err != nil {
		return nil, err
	}

	return db, nil
}
