package main

import (
	"log"
	"net/http"
	"os"

	"github.com/dikyayodihamzah/realtime-doc-collaboration/app/service/hubsrv"
	"github.com/dikyayodihamzah/realtime-doc-collaboration/app/service/wssrv"
	"github.com/dikyayodihamzah/realtime-doc-collaboration/pkg/config"
	"github.com/dikyayodihamzah/realtime-doc-collaboration/pkg/model"

	_ "github.com/joho/godotenv/autoload"
)

// Main function to set up the server
func main() {
	doc := new(model.Document)
	hub := config.NewHub(doc)
	go hubsrv.Run(hub)

	wss := wssrv.New(hub)
	wsTopic := os.Getenv("WS_TOPIC")
	http.HandleFunc(wsTopic, func(w http.ResponseWriter, r *http.Request) {
		wss.Serve(w, r)
	})

	addr := os.Getenv("SERVER_PORT")
	log.Println("Starting server on", addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}
