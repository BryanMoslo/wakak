package internal

import (
	"fmt"
	"wakak/internal/terms"

	"github.com/leapkit/core/server"
)

// AddServices is a function that will be called by the server
// to inject services in the context.
func AddServices(r server.Router) error {
	db, err := DB()
	if err != nil {
		return fmt.Errorf("connecting to the database: %w", err)
	}

	// Services that will be injected in the context
	r.Use(server.InCtxMiddleware("terms", terms.NewService(db)))

	return nil
}
