package terms

import (
	"wakak/internal/model"

	"github.com/jmoiron/sqlx"
)

type service struct {
	db *sqlx.DB
}

func NewService(db *sqlx.DB) model.TermsService {
	return &service{db}
}

func (s *service) Save(finding model.Term) error {
	_, err := s.db.NamedExec(`
		INSERT INTO findings
			(source, keyword, content)
		VALUES
			(:source, :keyword, :content)
		ON CONFLICT DO NOTHING`,
		finding,
	)

	return err
}

func (s *service) All() ([]model.Term, error) {
	var terms []model.Term
	return terms, s.db.Select(&terms, "SELECT * FROM findings")
}
