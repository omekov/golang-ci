package healthy

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/omekov/golang-ci/pkg/getenv"
)

// OK ...
const OK = `{"alive": true}`

// HealthCheckHandler - ...
func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Call Health Check")
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	io.WriteString(w, OK)
}

// Check - http request healthy
func Check() error {
	client := http.Client{}
	port := getenv.GetVar("PORT", "8081")
	resp, err := client.Get(fmt.Sprintf("http://localhost:%s/health", port))
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	answer := string(body)
	if answer != OK {
		return fmt.Errorf("Invalid Answer %s expected %s", answer, OK)
	}
	return nil
}
