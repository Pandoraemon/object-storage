package main

import (
	"fmt"
	"github.com/pandoraemon/object-storage/dataService/heartbeat"
	"github.com/pandoraemon/object-storage/dataService/locate"
	"github.com/pandoraemon/object-storage/dataService/objects"
	"github.com/pandoraemon/object-storage/dataService/temp"
	"log"
	"net/http"
	"os"
)

func main() {
	locate.CollectObjects()
	os.Setenv("RABBITMQ_SERVER", "amqp://guest:guest@10.128.137.42:5672")
	go heartbeat.StartHeartbeat()
	go locate.StartLocate()
	http.HandleFunc("/objects/", objects.Handler)
	http.HandleFunc("/temp/", temp.Handler)
	fmt.Println(os.Getenv("LISTEN_ADDRESS"))
	log.Fatal(http.ListenAndServe(os.Getenv("LISTEN_ADDRESS"), nil))
}
