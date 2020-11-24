package handlers

import (
	"log"
	"net/http"

	"github.com/omekov/golang-ci/pkg/getenv"
	"github.com/omekov/golang-ci/pkg/healthy"
)

// Mux - start http mux
func Mux() http.Handler {
	mux := http.NewServeMux()
	port := getenv.GetVar("PORT", "8081")
	mux.HandleFunc("/health", healthy.HealthCheckHandler)
	if err := http.ListenAndServe(":"+port, mux); err != nil {
		log.Fatal(err)
		return nil
	}
	return mux
}
