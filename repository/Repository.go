package repository

import (
	"database/sql"
	"log"
	"os"
)

type UserInfo struct {
	ID      string
	Email   string
	Picture string
	Token   string
}

func StoreUserInfo(userInfo UserInfo) {
	var count int
	database := initDatabaseConnection()
	row := database.QueryRow("select count(id) from userinfo where id = $1", userInfo.ID)
	row.Scan(&count)
	if count == 0 {
		queryString := `insert into userinfo(id, email, picture, token) values ($1, $2, $3, $4)`
		database.Exec(queryString, userInfo.ID, userInfo.Email, userInfo.Picture, userInfo.Token)
	} else if count == 1 {
		queryString := `update userinfo set token = $1 where id=$2`
		database.Exec(queryString, userInfo.Token, userInfo.ID)
	}
	database.Close()
}

func FetchUser(token string) (string, string) {
	var email string
	var picture string

	database := initDatabaseConnection()
	database.QueryRow("select email, picture from userinfo").Scan(&email, &picture)

	database.Close()

	return email, picture
}

func initDatabaseConnection() *sql.DB {
	dataSourceName := os.Getenv("DATABASE_URL")

	database, err := sql.Open("postgres", dataSourceName)
	if err != nil {
		log.Fatal(err)
	}

	return database
}
