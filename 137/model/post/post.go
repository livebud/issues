package post

import (
	"fmt"

	"github.com/livebud/issues/137/db"
)

type Model struct {
}

func (m *Model) FindMany() ([]*db.Post, error) {
	return db.Posts, nil
}

func (m *Model) Find(id int) (*db.Post, error) {
	for _, p := range db.Posts {
		if p.ID == id {
			return p, nil
		}
	}
	return nil, fmt.Errorf("post %d not found", id)
}
