// GoPrint project main.go
package main

import (
	"net/http"
)

func main() {
	http.Handle("/app/", http.StripPrefix("/app", http.FileServer(http.Dir("app"))))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "app/index.html")
	})
	http.ListenAndServe(":8080", nil)
}
