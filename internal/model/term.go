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
