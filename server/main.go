package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/luytbq/frontendlead-knockoff/api"
)

func main() {

	http.HandleFunc("/api/", api.HandleAPI)

	handleStatic()

	// Start the server
	port := ":5130"
	fmt.Println("Server running at http://localhost" + port)
	log.Fatal(http.ListenAndServe(port, nil))
}

func handleStatic() {
	staticDir := "static"

	// Serve static files
	staticHandler := http.FileServer(http.Dir(staticDir))
	http.Handle("/static/", http.StripPrefix("/static/", staticHandler))

	// Redirect root path "/" to "/home.html"
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/" {
			http.Redirect(w, r, "/home.html", http.StatusFound)
			return
		}
		http.ServeFile(w, r, staticDir+r.URL.Path)
	})
}
