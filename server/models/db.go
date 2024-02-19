package models

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func InitDb() *gorm.DB {
	var (
		db  *gorm.DB
		err error
	)

	db, err = gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
	})

	if err != nil {
		println("Error connecting to database", err)
	}

	// Migrate database
	db.AutoMigrate(&User{})

	// Seed database
	if err = UsersSeed(db); err != nil {
		panic(err)
	}

	return db
}
