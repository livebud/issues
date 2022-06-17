package controller

import (
	context "context"

	"github.com/livebud/issues/137/db"
	model "github.com/livebud/issues/137/model/post"
)

type Controller struct {
}

// Index of posts
// GET
func (c *Controller) Index(ctx context.Context) (posts []*db.Post, err error) {
	model := &model.Model{}
	return model.FindMany()
}

// Show post
// GET /:id
func (c *Controller) Show(ctx context.Context, id int) (post *db.Post, err error) {
	model := &model.Model{}
	return model.Find(id)
}
