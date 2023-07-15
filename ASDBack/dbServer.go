package ASDBack

import (
	"database/sql"
	"fmt"
	_ "modernc.org/sqlite"
)

func SetUp() {
	db, err := sql.Open("sqlite", "test.db")
	if err != nil {
		fmt.Println("Error opening database:", err)
	}
	defer db.Close()
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS keys (id INTEGER PRIMARY KEY AUTOINCREMENT, apikey TEXT)`)
	if err != nil {
		fmt.Println("Error creating table:", err)
		return
	}

}

func CheckApiKey(apiKey string) bool {

	db, err := sql.Open("sqlite", "test.db")
	if err != nil {
		fmt.Println("Error opening database:", err)
		return false
	}
	defer db.Close()

	var exists bool
	err = db.QueryRow("SELECT EXISTS(SELECT 1 FROM keys WHERE apikey = ?)", apiKey).Scan(&exists)
	if err != nil {
		fmt.Println("Error querying data:", err)
		return false
	}

	if exists {
		return true
	} else {
		fmt.Println("API key does not exist in the database")
	}
	return false
}
