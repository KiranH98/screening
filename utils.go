package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"sync"
)

// createConnection creates a connection to mysql database
/* func createConnection() *sql.DB {
	db, err := sql.Open("mysql", "root:password@tcp(127.0.0.1:3306)/test")
	fmt.Println("sql open " + err.Error())
	return db
} */

var db *sql.DB
var dbOnce sync.Once

type Config struct {
	DBUser     string `json:"db_user"`
	DBPassword string `json:"db_password"`
	DBName     string `json:"db_name"`
}

func initDB() {
	dbOnce.Do(func() {
		config, err := loadConfig("config.json")
		if err != nil {
			fmt.Println("Error loading config:", err)
			return
		}

		connectionString := fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/%s", config.DBUser, config.DBPassword, config.DBName)
		db, err = sql.Open("mysql", connectionString)
		if err != nil {
			fmt.Println("Error opening database:", err)
		}
	})
}

func loadConfig(filename string) (Config, error) {
	var config Config

	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return config, err
	}

	err = json.Unmarshal(data, &config)
	if err != nil {
		return config, err
	}

	return config, nil
}
