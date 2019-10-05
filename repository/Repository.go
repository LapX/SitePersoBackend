package repository

import (
	"database/sql"
	"fmt"
	"log"
	"os"
)

type UserInfo struct {
	ID      string
	Email   string
	Picture string
}

func StoreUserInfo(userInfo UserInfo) {
	var id string
	dataSourceName := os.Getenv("DATABASE_URL")

	database, err := sql.Open("postgres", dataSourceName)
	if err != nil {
		log.Fatal(err)
	}

	row := database.QueryRow("select id from userinfo").Scan(&id)

	if row != nil {
		if row == sql.ErrNoRows {
			fmt.Println("Inserting into database")
			fmt.Println(userInfo.ID)
			fmt.Println(userInfo.Email)
			fmt.Println(userInfo.Picture)
			queryString := `insert into userinfo(id, email, picture) values ($1, $2, $3)`
			database.Exec(queryString, userInfo.ID, userInfo.Email, userInfo.Picture)
		}
	}

	database.Close()
}
