package loaders

import (
	"context"
	"net/http"
	"strconv"
	"time"

	"github.com/eeQuillibrium/posts/graph/model"
	"github.com/jmoiron/sqlx"
	"github.com/vikstrous/dataloadgen"
)

type ctxKey string

const (
	loadersKey = ctxKey("dataloaders")
)

type commentReader struct {
	db *sqlx.DB
}

func (u *commentReader) getComments(ctx context.Context, commentIDs []string) ([]*model.Comment, []error) {
	var commentsIDs []int = make([]int, 0, len(commentIDs))
	for i := 0; i < len(commentIDs); i++ {
		id, err := strconv.Atoi(commentIDs[i])
		if err != nil {
			return nil, []error{err}
		}
		commentsIDs = append(commentsIDs, id)
	}

	query, args, err := sqlx.In("SELECT * FROM Comments WHERE id IN (?)", commentsIDs)
	if err != nil {
		return nil, []error{err}
	}
	var comments []*model.Comment
	
	err = u.db.SelectContext(ctx, &comments, u.db.Rebind(query), args...)
	if err != nil {
		return comments, []error{err}
	}
	return comments, nil
}

type Loaders struct {
	CommentLoader *dataloadgen.Loader[string, *model.Comment]
}

func NewLoaders(db *sqlx.DB) *Loaders {
	ur := &commentReader{db: db}
	return &Loaders{
		CommentLoader: dataloadgen.NewLoader(ur.getComments, dataloadgen.WithWait(time.Millisecond)),
	}
}

// Middleware injects data loaders into the context
func Middleware(db *sqlx.DB, next http.Handler) http.Handler {
	// return a middleware that injects the loader to the request context
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		loader := NewLoaders(db)
		r = r.WithContext(context.WithValue(r.Context(), loadersKey, loader))
		next.ServeHTTP(w, r)
	})
}

// For returns the dataloader for a given context
func For(ctx context.Context) *Loaders {
	return ctx.Value(loadersKey).(*Loaders)
}

// GetUser returns single comment by id efficiently
func GetComment(ctx context.Context, commentID string) (*model.Comment, error) {
	loaders := For(ctx)
	return loaders.CommentLoader.Load(ctx, commentID)
}

// GetUsers returns many comment by ids efficiently
func GetComments(ctx context.Context, commentIDs []string) ([]*model.Comment, error) {
	loaders := For(ctx)
	return loaders.CommentLoader.LoadAll(ctx, commentIDs)
}