package main

import (
	"fmt"
	"testing"

	"github.com/afchatfield/SongWebService/pkg/app"
	"github.com/afchatfield/SongWebService/pkg/models"
	"github.com/stretchr/testify/assert"
)

func TestApp(t *testing.T) {
	app.Connect()
}

func TestModels(t *testing.T) {
	fmt.Println("QUERY test")
	Songs := models.GetSongs()
	assert.Equal(t, len(Songs), 100, "100 Songs returned")
	search := "Ross From Friends"
	Songs = models.GetSongsBySearch("SELECT * FROM song WHERE SONG_NAME LIKE ? OR ARTIST LIKE ? ORDER BY ID ASC LIMIT 100;", search, search)
	fmt.Println(len(Songs))
	song := models.GetSongByID("262243")
	fmt.Println(song)

	fmt.Println("\nINSERT test")
	newsong := new(models.Song)
	newsong.Name = "Billie (Loving Arms)"
	newsong.Artist = "Fred Again.."
	models.CreateSong(newsong)
	search = "Fred Again.."
	Songs = models.GetSongsBySearch("SELECT * FROM song WHERE SONG_NAME LIKE ? OR ARTIST LIKE ? ORDER BY ID ASC LIMIT 100;", search, search)
	fmt.Println(Songs)

	fmt.Println("\nUPDATE test")
	updatesong := Songs[0]
	updatesong.Name = "Faisal (Envelops Me)"
	models.UpdateSong(updatesong)
	Songs = models.GetSongsBySearch("SELECT * FROM song WHERE SONG_NAME LIKE ? OR ARTIST LIKE ? ORDER BY ID ASC LIMIT 100;", search, search)
	fmt.Println(Songs)

	fmt.Println("\nDELETE test")
	models.DeleteSongByID(fmt.Sprintf("%v", Songs[0].Id))
	Songs = models.GetSongsBySearch("SELECT * FROM song WHERE SONG_NAME LIKE ? OR ARTIST LIKE ? ORDER BY ID ASC LIMIT 100;", search, search)
	fmt.Println(Songs)
}
