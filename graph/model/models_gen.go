// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type Comment struct {
	ID        int        `json:"id" db:"id"`
	PostID    int        `json:"postId" db:"post_id"`
	ParentID  *int       `json:"parentId,omitempty" db:"parent_id"`
	UserID    int        `json:"userId" db:"user_id"`
	Text      string     `json:"text" db:"text"`
	Level     int        `json:"level" db:"level"`
	CreatedAt string     `json:"createdAt" db:"created_at"`
	Comments  []*Comment `json:"comments"`
}

type Mutation struct {
}

type NewComment struct {
	ParentID *int   `json:"parentId,omitempty"`
	Text     string `json:"text"`
	UserID   int    `json:"userId"`
	PostID   int    `json:"postId"`
	Level    int    `json:"level"`
}

type NewPost struct {
	Header string `json:"header"`
	Text   string `json:"text"`
	UserID int    `json:"userId"`
}

type NewUser struct {
	Name     string `json:"name"`
	Login    string `json:"login"`
	Password string `json:"password"`
}

type Pagination struct {
	Offset int `json:"offset"`
	Limit  int `json:"limit"`
}

type Post struct {
	ID        int        `json:"id"`
	UserID    int        `json:"userId" db:"user_id"`
	Text      string     `json:"text"`
	Header    string     `json:"header"`
	CreatedAt string     `json:"createdAt" db:"created_at"`
	IsClosed bool `json:"is_closed" db:"is_closed"`
	Comments  []*Comment `json:"comments"`
}

type Query struct {
}
type Subscription struct {
}
type User struct {
	ID       int    `json:"id"`
	Login    string `json:"login"`
	Password string `json:"password"`
	Name     string `json:"name"`
}

type Notification struct {
	Text     string `json:"text"`
	IssuerID int    `json:"issuerId"`
}