package model

type Term struct {
	ID      int    `db:"id"`
	Source  string `db:"source"`
	Keyword string `db:"keyword"`
	Content string `db:"content"`
}

type TermsService interface {
	Save(finding Term) error
	All() ([]Term, error)
}
