package controller

import (
	context "context"
)

type Controller struct {
	// Dependencies...
}

// Hello struct
type Hello struct {
	// Fields...
}

// Index of hellos
// GET /
func (c *Controller) Index(ctx context.Context) (hellos []*Hello, err error) {
	return []*Hello{}, nil
}
