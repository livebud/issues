package controller

import (
	context "context"

	"github.com/livebud/issues/137/model/post"
)

type Controller struct {
}

type Post struct {
	Model post.Model
}

// Show post
// GET /:id
func (c *Controller) Show(ctx context.Context, id int) (post *Post, err error) {
	return &Post{}, nil
}
