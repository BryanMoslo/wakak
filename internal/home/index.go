package home

import (
	"net/http"
	"wakak/internal/model"

	"github.com/leapkit/core/render"
)

func Index(w http.ResponseWriter, r *http.Request) {
	terms := r.Context().Value("terms").(model.TermsService)
	findings, err := terms.All()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	rw := render.FromCtx(r.Context())
	rw.Set("findings", findings)

	err = rw.Render("home/index.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
