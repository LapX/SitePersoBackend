package database

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
	defer database.Close()

	database.QueryRow("select count(id) from userinfo where id = $1", userInfo.ID).Scan(&count)
	if count == 0 {
		queryString := `insert into userinfo(id, email, picture, token) values ($1, $2, $3, $4)`
		database.Exec(queryString, userInfo.ID, userInfo.Email, userInfo.Picture, userInfo.Token)
	} else if count == 1 {
		queryString := `update userinfo set token = $1, picture = $2 where id=$3`
		database.Exec(queryString, userInfo.Token, userInfo.Picture, userInfo.ID)
	}
}

func ModifyNumberOfGraphs(token string, amount int) {
	var id string
	var count int
	numberOfGraphs := 0
	database := initDatabaseConnection()
	defer database.Close()

	database.QueryRow("select id from userinfo where token=$1", token).Scan(&id)
	database.QueryRow("select count(id) from userdashboard where id = $1", id).Scan(&count)

	if count == 0 {
		if amount > 0 {
			numberOfGraphs = numberOfGraphs + amount
		}
		if numberOfGraphs > 20 {
			numberOfGraphs = 20
		}
		queryString := `insert into userdashboard(id, numberofgraphs) values ($1, $2)`
		database.Exec(queryString, id, numberOfGraphs)
	} else if count == 1 {
		database.QueryRow("select numberofgraphs from userdashboard where id=$1", id).Scan(&numberOfGraphs)
		numberOfGraphs = numberOfGraphs + amount
		if numberOfGraphs < 0 {
			numberOfGraphs = 0
		} else if numberOfGraphs > 20 {
			numberOfGraphs = 20
		}
		queryString := `update userdashboard set numberofgraphs = $1 where id=$2`
		database.Exec(queryString, numberOfGraphs, id)
	}
}

func FetchUser(token string) (string, string) {
	var email string
	var picture string

	database := initDatabaseConnection()
	defer database.Close()

	database.QueryRow("select email, picture from userinfo where token=$1", token).Scan(&email, &picture)

	return email, picture
}

func FetchNumberOfEarningsGraphs(token string) int {
	var id string
	var numberOfEarningsGraphs int
	database := initDatabaseConnection()
	defer database.Close()

	database.QueryRow("select id from userinfo where token=$1", token).Scan(&id)
	database.QueryRow("select numberofgraphs from userdashboard where id=$1", id).Scan(&numberOfEarningsGraphs)

	return numberOfEarningsGraphs
}

func FetchNumberOfEarningsGraph() int {
	noAuthId := 0
	var numberOfEarningsGraphs int
	database := initDatabaseConnection()
	defer database.Close()

	database.QueryRow("select numberofgraphs from userdashboard where id=$1", noAuthId).Scan(&numberOfEarningsGraphs)

	return numberOfEarningsGraphs
}

func initDatabaseConnection() *sql.DB {
	dataSourceName := os.Getenv("DATABASE_URL")
	database, err := sql.Open("postgres", dataSourceName)
	if err != nil {
		log.Fatal(err)
	}

	return database
}
