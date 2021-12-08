package models

import (
	"database/sql"
	"log"

	"github.com/afchatfield/SongWebService/pkg/app"
)

// Record type (Song)
type Song struct {
	Id     int    `yaml:"id"`
	Name   string `yaml:"name"`
	Artist string `yaml:"artist"`
	GUID   string `yaml:"guid"`
}

var db *sql.DB

func init() {
	app.Connect()
	db = app.GetDB()
}

func CollectResults(results *sql.Rows) []Song {
	var Songs []Song
	// scan results
	for results.Next() {
		var song Song
		err := results.Scan(&song.Id, &song.Name, &song.Artist, &song.GUID)
		if err != nil {
			panic(err.Error())
		}
		Songs = append(Songs, song)
	}

	return Songs
}

func GetSongs() []Song { // query string
	query := "SELECT * FROM song ORDER BY ID ASC LIMIT 100;"
	results, err := db.Query(query)
	if err != nil {
		panic(err.Error())
	}
	return CollectResults(results)
}

func GetSongsBySearch(query string, search1 string, search2 string) []Song {
	results, err := db.Query(query, search1, search2)
	if err != nil {
		panic(err.Error())
	}
	return CollectResults(results)
}

func GetSongByID(id string) Song {
	result, err := db.Query("SELECT * FROM song WHERE id=?", id)
	if err != nil {
		panic(err.Error())
	}
	var song Song
	for result.Next() {
		err = result.Scan(&song.Id, &song.Name, &song.Artist, &song.GUID)
		if err != nil {
			panic(err.Error())
		}
	}
	return song
}

func CreateSong(s *Song) {
	// guid := uuid.New().String()
	insForm, err := db.Prepare("INSERT INTO song(SONG_NAME, ARTIST, GUID) VALUES(?,?,?)")
	if err != nil {
		panic(err.Error())
	}
	insForm.Exec(s.Name, s.Artist, s.GUID)
	log.Println("INSERT: Name: " + s.Name + " | Artist: " + s.Artist + " | GUID: " + s.GUID)
}

func DeleteSongByID(id string) {
	delForm, err := db.Prepare("DELETE FROM song WHERE id=?")
	if err != nil {
		panic(err.Error())
	}
	delForm.Exec(id)
	log.Printf("DELETE: ID: %v", id)
}

func UpdateSong(s Song) {
	insForm, err := db.Prepare("UPDATE song SET SONG_NAME=?, ARTIST=? WHERE id=?")
	if err != nil {
		panic(err.Error())
	}
	insForm.Exec(s.Name, s.Artist, s.Id)
	log.Println("UPDATE: Name: " + s.Name + " | Artist: " + s.Artist)
}
