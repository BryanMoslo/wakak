package terms

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"wakak/internal/model"
)

func SaveFindings(w http.ResponseWriter, r *http.Request) {
	terms := r.Context().Value("terms").(model.TermsService)
	term := model.Term{}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Unable to read request body", http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()

	err = json.Unmarshal(body, &term)
	if err != nil {
		http.Error(w, "Unable to parse request body", http.StatusBadRequest)
		return
	}

	if err := terms.Save(term); err != nil {
		fmt.Println(err)
	}

	w.Write([]byte("Request received"))
}
