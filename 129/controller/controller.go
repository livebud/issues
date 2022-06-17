package controller

import (
	context "context"
	"fmt"
)

type Controller struct {
	// Dependencies...
}

// Hello struct
type Hello struct {
	// Fields...
}

// Index of hellos
// GET
func (c *Controller) Index(ctx context.Context) (hellos []*Hello, err error) {
	fmt.Println("HI!!......")
	return []*Hello{}, nil
}
