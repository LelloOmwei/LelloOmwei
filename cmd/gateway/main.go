package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"

	"omwei-core/pkg/semantics" // Uisti sa, že táto cesta k tvojim structom sedí
)

func main() {
	// --- CESTA 1: TU RUST POSIELA DÁTA ---
	http.HandleFunc("/ingest", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		if r.Method == "OPTIONS" {
			return
		}

		var t semantics.Token
		if err := json.NewDecoder(r.Body).Decode(&t); err != nil {
			http.Error(w, "Invalid Signal", 400)
			return
		}

		// Ak Rust neposlal Timestamp, dáme aktuálny
		if t.Ts == 0 {
			t.Ts = time.Now().Unix()
		}

		entry := fmt.Sprintf("[%d] %s -> %s (val: %.2f)\n", t.Ts, t.ID, t.Sign, t.Val)
		fmt.Print("📥 .ow " + entry)

		// Zápis do súboru v aktuálnom priečinku
		f, err := os.OpenFile("history.ow", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			fmt.Println("❌ Chyba zápisu:", err)
			return
		}
		defer f.Close()
		f.WriteString(entry)

		w.WriteHeader(http.StatusAccepted)
	})

	// --- CESTA 2: TU WEB ČÍTA DÁTA ---
	http.HandleFunc("/history", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		
		// Čítame ten istý súbor
		data, err := os.ReadFile("history.ow")
		if err != nil {
			w.Write([]byte("Zatiaľ žiadne záznamy."))
			return
		}
		w.Header().Set("Content-Type", "text/plain")
		w.Write(data)
	})

	fmt.Println("🏛️ OM-WEI Gateway is LISTENING on :8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("❌ Server spadol:", err)
	}
}
