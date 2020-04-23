package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/R-V-S/prstats/subsystem"
	"github.com/R-V-S/prstats/subsystem/router"
)

// @title PR Stats
// @version 1.0.0
// @tag.name PR Stats
// @tag.description A PR Stat is a collection of calculated metrics about a PR or a set of PRs
func main() {
	os.Setenv("TZ", "UTC")

	subsystem.LoadConfig()

	router := router.NewRouter()
	port := os.Getenv("HTTP_API_PORT")

	fmt.Printf("Running server on port " + port)

	http.ListenAndServe(":"+port, router)
}
