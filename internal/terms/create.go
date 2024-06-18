package terms

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"wakak/internal/model"

	"github.com/leapkit/core/render"
)

func New(w http.ResponseWriter, r *http.Request) {
	rw := render.FromCtx(r.Context())

	err := rw.Render("terms/create.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func Create(w http.ResponseWriter, r *http.Request) {
	rw := render.FromCtx(r.Context())
	termsInput := strings.Split(r.FormValue("Terms"), ",")
	terms := cleanTerms(termsInput)

	if len(terms) == 0 {
		rw.Set("responseMessages", []string{"This field can't be blank"})
		rw.Set("responseClass", "text-red-600")
		err := rw.RenderClean("terms/response.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	request := model.TermRequest{
		Keywords:    terms,
		CallbackURL: "http://localhost:4000/findings",
	}

	response, err := register(request)
	if err != nil {
		rw.Set("responseMessages", []string{"Internal issue, not your fault."})
		rw.Set("responseClass", "text-red-600")
		err = rw.RenderClean("terms/response.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	rw.Set("responseMessages", []string{response.Message})
	dataResponse := response.Data.(map[string]interface{})
	if dataResponse["errors"] != nil {
		rw.Set("responseMessages", dataResponse["errors"].([]interface{}))
	}

	rw.Set("responseClass", "text-green-600")
	if response.Status != http.StatusCreated {
		rw.Set("responseClass", "text-red-600")
	}

	err = rw.RenderClean("terms/response.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func register(termReq model.TermRequest) (model.TermResponse, error) {
	url := "http://localhost:3000/darkweb/register"
	method := "POST"

	reqBody, err := json.Marshal(termReq)
	if err != nil {
		return model.TermResponse{}, err
	}

	client := &http.Client{}
	req, err := http.NewRequest(method, url, bytes.NewBuffer(reqBody))

	if err != nil {
		fmt.Println(err)
		return model.TermResponse{}, err
	}

	// Adding headers
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "sss")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return model.TermResponse{}, err
	}
	defer res.Body.Close()

	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return model.TermResponse{}, err
	}

	response := model.TermResponse{}
	err = json.Unmarshal(resBody, &response)
	if err != nil {
		fmt.Println("Error parsing response:", err)
		return model.TermResponse{}, err
	}

	return response, nil
}

func cleanTerms(termsInput []string) []string {
	terms := []string{}

	for _, term := range termsInput {
		trimmedTerm := strings.TrimSpace(term)
		if trimmedTerm == "" {
			continue
		}
		terms = append(terms, trimmedTerm)
	}

	return terms
}
