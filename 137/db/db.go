package db

type Post struct {
	ID    int
	Title string
	Body  string
}

var Posts = []*Post{
	{
		ID:    1,
		Title: "first post",
		Body:  "first body",
	},
	{
		ID:    2,
		Title: "second post",
		Body:  "second body",
	},
}
