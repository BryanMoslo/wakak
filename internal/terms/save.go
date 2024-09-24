package terms

import (
	"io"
	"log"
	"net/http"
)

func SaveFindings(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Unable to read request body", http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()

	log.Printf("Received Body: %s", string(body))
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Body received"))
}
