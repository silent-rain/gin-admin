package main

import (
	"fmt"

	"gorm.io/driver/sqlite" // Sqlite driver based on CGO

	// "github.com/glebarez/sqlite" // Pure go SQLite driver, checkout https://github.com/glebarez/sqlite for details
	"gorm.io/gorm"
)

func main() {
	// github.com/mattn/go-sqlite3
	db, err := gorm.Open(sqlite.Open("data.db"), &gorm.Config{})
	if err != nil {
		fmt.Printf("err: %#v\n", err)
		return
	}

	fmt.Printf("%#v\n", db)
}
