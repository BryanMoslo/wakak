package model

type Term struct {
	ID      int    `db:"id" json:"-"`
	Source  string `db:"source" json:"source"`
	Keyword string `db:"keyword" json:"keyword"`
	Content string `db:"content" json:"content"`
}

type TermsService interface {
	Save(finding Term) error
	All() ([]Term, error)
}

type TermRequest struct {
	Keywords    []string `json:"keywords"`
	CallbackURL string   `json:"callback_url"`
}

// {"status":201,"message":"keyword(s) registered successfully","data":{"callback_url":"http://localhost:4000/findings","keywords":"BryanMoslo"}}
type TermResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}
