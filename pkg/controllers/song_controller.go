package controllers

import (
	"net/http"
	"text/template"

	"github.com/google/uuid"

	"github.com/afchatfield/SongWebService/pkg/models"
	_ "github.com/go-sql-driver/mysql"
)

var tmpl = template.Must(template.ParseGlob("pkg/form/*"))

// standard handler, returns records
func Index(w http.ResponseWriter, r *http.Request) {
	// w.Header().Set("Content-Type", "application/json")
	results := models.GetSongs()
	tmpl.ExecuteTemplate(w, "Index", results)
}

func Search(w http.ResponseWriter, r *http.Request) {
	search := r.FormValue("search_field")
	results := models.GetSongsBySearch("SELECT * FROM song WHERE SONG_NAME LIKE ? OR ARTIST LIKE ? ORDER BY ID ASC LIMIT 100;", search, search)
	tmpl.ExecuteTemplate(w, "Index", results)
}

func Show(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	song := models.GetSongByID(id)
	tmpl.ExecuteTemplate(w, "Show", song)
}

func New(w http.ResponseWriter, r *http.Request) {
	guid := uuid.New()
	tmpl.ExecuteTemplate(w, "New", guid)
}

func Insert(w http.ResponseWriter, r *http.Request) {
	// get song from New page
	// song := &models.Song{}
	// utils.ParseBody(r, song)
	// models.CreateSong(song)
	http.Redirect(w, r, "/", http.StatusMovedPermanently) //301
}

func Edit(w http.ResponseWriter, r *http.Request) {
	Id := r.URL.Query().Get("id")
	song := models.GetSongByID(Id)
	tmpl.ExecuteTemplate(w, "Edit", song)
}

func Update(w http.ResponseWriter, r *http.Request) {
	// get song from Edit page
	// song := &models.Song{}
	// utils.ParseBody(r, song)
	// models.UpdateSong(song)
	http.Redirect(w, r, "/", http.StatusMovedPermanently) //301
}

func Delete(w http.ResponseWriter, r *http.Request) {
	Id := r.URL.Query().Get("id")
	models.DeleteSongByID(Id)
	http.Redirect(w, r, "/", http.StatusMovedPermanently) //301
}
