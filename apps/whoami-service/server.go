package main

import (
	"fmt"
	"net/http"
)

func whoami(w http.ResponseWriter, req *http.Request) {
	var name = req.Header.Get("X-User")

	if name == "" {
		name = "Unknown"
	}

	fmt.Fprint(w, name)
}

func main() {
	fmt.Println("whoami listening on http://localhost:5000")
	http.HandleFunc("/whoami", whoami)
	http.ListenAndServe(":5000", nil)
}
