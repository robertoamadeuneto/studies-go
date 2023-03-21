package main

import (
	"emailn/internal/entrypoint/configuration"
	"net/http"
)

func main() {
	router := configuration.BuildRouter()

	http.ListenAndServe(":3000", router)
}
