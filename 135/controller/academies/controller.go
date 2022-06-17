package academies

import (
	context "context"
)

type Controller struct {
	// Dependencies...
}

// Academy struct
type Academy struct {
	// Fields...
}

// Index of academies
// GET /academies
func (c *Controller) Index(ctx context.Context) (academies []*Academy, err error) {
	return []*Academy{}, nil
}

// Show academy
// GET /academies/:id
func (c *Controller) Show(ctx context.Context, id int) (academy *Academy, err error) {
	return &Academy{}, nil
}
