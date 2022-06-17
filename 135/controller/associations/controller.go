package associations

import (
	context "context"
)

type Controller struct {
	// Dependencies...
}

// Association struct
type Association struct {
	// Fields...
}

// Index of associations
// GET /associations
func (c *Controller) Index(ctx context.Context) (associations []*Association, err error) {
	return []*Association{}, nil
}

// Show association
// GET /associations/:id
func (c *Controller) Show(ctx context.Context, id int) (association *Association, err error) {
	return &Association{}, nil
}
