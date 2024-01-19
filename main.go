package main

import (
	"net/http"
)

func main() {
	initDB()
	setupJSONAPI()
	http.ListenAndServe(":80", nil)
}
