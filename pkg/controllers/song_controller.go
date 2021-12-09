package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"text/template"

	"github.com/google/uuid"

	"github.com/afchatfield/SongWebService/pkg/models"
	"github.com/afchatfield/SongWebService/pkg/utils"
	_ "github.com/go-sql-driver/mysql"
)

var tmpl = template.Must(template.ParseGlob("pkg/form/*"))

// standard handler, returns records
func IndexJson(w http.ResponseWriter, r *http.Request) {
	songs := models.GetSongs()
	res, _ := json.Marshal(songs)
	w.Header().Set("Content-Type", "aplication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func Index(w http.ResponseWriter, r *http.Request) {
	songs := models.GetSongs()
	tmpl.ExecuteTemplate(w, "Index", songs)
}

func SearchJson(w http.ResponseWriter, r *http.Request) {
	search := r.FormValue("search_field")
	results := models.GetSongsBySearch("SELECT * FROM song WHERE SONG_NAME LIKE ? OR ARTIST LIKE ? ORDER BY ID ASC LIMIT 100;", search, search)
	res, _ := json.Marshal(results)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func Search(w http.ResponseWriter, r *http.Request) {
	search := r.FormValue("search_field")
	results := models.GetSongsBySearch("SELECT * FROM song WHERE SONG_NAME LIKE ? OR ARTIST LIKE ? ORDER BY ID ASC LIMIT 100;", search, search)
	tmpl.ExecuteTemplate(w, "Index", results)
}

func ShowJson(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	song := models.GetSongByID(id)
	res, _ := json.Marshal(song)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func Show(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	song := models.GetSongByID(id)
	tmpl.ExecuteTemplate(w, "Show", song)
}

func CreateSongJson(w http.ResponseWriter, r *http.Request) {
	CreateSong := &models.Song{}
	utils.ParseBody(r, CreateSong)

}

func New(w http.ResponseWriter, r *http.Request) {
	guid := uuid.New()
	tmpl.ExecuteTemplate(w, "New", guid)
}

func Insert(w http.ResponseWriter, r *http.Request) {
	song := &models.Song{
		Name:   r.FormValue("name"),
		Artist: r.FormValue("artist"),
		GUID:   r.FormValue("guid")}
	models.CreateSong(song)
	http.Redirect(w, r, "/", http.StatusMovedPermanently) //301
}

func Edit(w http.ResponseWriter, r *http.Request) {
	Id := r.URL.Query().Get("id")
	song := models.GetSongByID(Id)
	tmpl.ExecuteTemplate(w, "Edit", song)
}

func Update(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(r.FormValue("id"))
	song := &models.Song{
		Id:     id,
		Name:   r.FormValue("name"),
		Artist: r.FormValue("artist")}
	models.UpdateSong(song)
	http.Redirect(w, r, "/", http.StatusMovedPermanently) //301
}

func Delete(w http.ResponseWriter, r *http.Request) {
	Id := r.URL.Query().Get("id")
	models.DeleteSongByID(Id)
	http.Redirect(w, r, "/", http.StatusMovedPermanently) //301
}
