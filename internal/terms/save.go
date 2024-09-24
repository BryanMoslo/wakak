package terms

import (
	"fmt"
	"io"
	"net/http"
)

func SaveFindings(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Unable to read request body", http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()

	fmt.Println(string(body))
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Body received"))
}
