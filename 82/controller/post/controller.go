package post

import (
	context "context"
)

type Controller struct {
	// Dependencies...
}

// Post struct
type Post struct {
	ID int `json:"id"`
	// Fields...
}

// Index of posts
// GET /post
func (c *Controller) Index(ctx context.Context) (posts []*Post, err error) {
	return []*Post{
		{
			ID: 1,
		},
		{
			ID: 2,
		},
	}, nil
}

// Show post
// GET /post/:id
func (c *Controller) Show(ctx context.Context, id int) (post *Post, err error) {
	return &Post{}, nil
}
