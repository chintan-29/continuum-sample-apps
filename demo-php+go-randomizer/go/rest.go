package main

import (
	"io"
	"net/http"
	"log"
	"os"
	"math/rand"
	"encoding/json"
)

// Randomizer REST API
func RandoServer(w http.ResponseWriter, req *http.Request) {

	// Build response data
	type response struct {
		API_Instance	string
		API_FQN		string
		Number     int
	}
	rstruct := response {
		API_Instance: os.Getenv("CNTM_INSTANCE_UUID"),
		API_FQN: os.Getenv("CNTM_JOB_FQN"),
		Number: rand.Intn(100),
	}

	// Reformat response into JSON
	byteout, err := json.MarshalIndent(rstruct, "", "  ")
	if err != nil {
		log.Fatal("error:", err)
	}
	out := string(byteout[:])

	// Display data
	io.WriteString(w, out)
}

func main() {
	http.HandleFunc("/", RandoServer)
	err := http.ListenAndServe(":"+os.Getenv("PORT"), nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
