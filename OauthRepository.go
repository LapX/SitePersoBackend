package main

import (
	"database/sql"
	"log"
	"os"
)

func getOauthInfo() (string, string) {
	var clientid string
	var clientSecret string
	dataSourceName := getDataSourceName()
	database, err := sql.Open("postgres", dataSourceName)
	if err != nil {
		log.Fatal(err)
	}
	row := database.QueryRow("")
	row.Scan(&clientid, &clientSecret)
	database.Close()
	return clientid, clientSecret
}

func getDataSourceName() string {
	dataSourceName := os.Getenv("DATABASE_URL")
	if dataSourceName == "" {
		dataSourceName = "PASSWORD"
	}
	return dataSourceName
}
