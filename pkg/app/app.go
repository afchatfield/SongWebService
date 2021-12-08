package app

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v3"
)

var db *sql.DB

func Connect() {
	configFilepath := "pkg/config/database.yaml"
	yfile, err := ioutil.ReadFile(configFilepath)
	if err != nil {
		log.Fatal(err)
	}

	var dbConfig Database

	err2 := yaml.Unmarshal([]byte(yfile), &dbConfig)
	if err2 != nil {
		log.Fatal(err2)
	}

	dbAddr := dbConfig.User + ":" + dbConfig.Pass + "@" + dbConfig.Protocol +
		"(" + dbConfig.Host + ":" + dbConfig.Port + ")/" + dbConfig.Name
	trydb, err3 := sql.Open("mysql", dbAddr)

	// if there is an error opening the connection, handle it
	if err3 != nil {
		panic(err.Error())
	}

	db = trydb

	fmt.Println("Successfully connected to database")
}

func GetDB() *sql.DB {
	return db
}
