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
	findings := []model.Term{}

	data := []struct {
		Content []string
		Source  []string
		Keyword []string
	}{}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Unable to read request body", http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()

	err = json.Unmarshal(body, &data)
	if err != nil {
		http.Error(w, "Unable to parse request body", http.StatusBadRequest)
		return
	}

	for _, v := range data {
		findings = append(findings, model.Term{
			Source:  v.Source[0],
			Content: v.Content[0],
			Keyword: v.Keyword[0],
		})
	}

	for _, finding := range findings {
		if err := terms.Save(finding); err != nil {
			fmt.Println(err)
		}
	}

	w.Write([]byte("Request received"))
}
