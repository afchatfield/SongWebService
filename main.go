package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/afchatfield/SongWebService/pkg/app"
	"github.com/afchatfield/SongWebService/pkg/routes"
	_ "github.com/go-sql-driver/mysql"
	yaml "gopkg.in/yaml.v3"
)

// Initiate web server
func main() {
	yfile, err := ioutil.ReadFile("pkg/config/connection.yaml")
	if err != nil {
		log.Fatal(err)
	}

	var serveConfig app.Serve
	err2 := yaml.Unmarshal([]byte(yfile), &serveConfig)
	if err2 != nil {
		log.Fatal(err2)
	}

	router := routes.GetRouter()
	address := fmt.Sprintf("%s:%d", serveConfig.Host, serveConfig.Port)
	srv := &http.Server{
		Handler:      router,
		Addr:         address, //"127.0.0.1:9100",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}
