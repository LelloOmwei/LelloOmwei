package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"
)

type SemanticToken struct {
	UID       string  "json:\"uid\""
	Entity    string  "json:\"ent\""
	State     string  "json:\"st\""
	Certainty float64 "json:\"crt\""
	Context   string  "json:\"ctx\""
	Timestamp int64   "json:\"ts\""
}

func ingestHandler(w http.ResponseWriter, r *http.Request) {
	var t SemanticToken
	if err := json.NewDecoder(r.Body).Decode(&t); err != nil {
		http.Error(w, "Bad Request", 400)
		return
	}

	timestamp := time.Unix(t.Timestamp, 0).Format("2006-01-02 15:04:05")
	entry := fmt.Sprintf("[%s] NODE: %s | STATE: %s | CTX: %s\n", timestamp, t.UID, t.State, t.Context)
	
	fmt.Print("📥 " + entry)

	f, err := os.OpenFile("history.ow", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err == nil {
		defer f.Close()
		f.WriteString(entry)
	}

	w.WriteHeader(http.StatusOK)
}

func main() {
	fmt.Println("🏛️ OM-WEI Gateway (with Memory) is LISTENING on :8080")
	http.HandleFunc("/ingest", ingestHandler)
	http.ListenAndServe(":8080", nil)
}
