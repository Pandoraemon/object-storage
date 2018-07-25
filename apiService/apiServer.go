package main

import (
	"github.com/pandoraemon/object-storage/apiService/heartbeat"
	"github.com/pandoraemon/object-storage/apiService/locate"
	"github.com/pandoraemon/object-storage/apiService/objects"
	"github.com/pandoraemon/object-storage/apiService/temp"
	"github.com/pandoraemon/object-storage/apiService/versions"
	"log"
	"net/http"
	"os"
)

func main() {
	os.Setenv("RABBITMQ_SERVER", "amqp://guest:guest@10.128.137.42:5672")
	os.Setenv("ES_SERVER", "10.128.137.42:9200")
	go heartbeat.ListenHeartbeat()
	http.HandleFunc("/objects/", objects.Handler)
	http.HandleFunc("/locate/", locate.Handler)
	http.HandleFunc("/versions/", versions.Handler)
	http.HandleFunc("/temp/", temp.Handler)
	log.Fatal(http.ListenAndServe(os.Getenv("LISTEN_ADDRESS"), nil))
}
