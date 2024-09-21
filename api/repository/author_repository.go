package repository

import (
	"context"
	"sqlc/api/models"
	"sqlc/internal/database"

	"github.com/jackc/pgx/v5/pgxpool"
)

type AuthorRepository struct {
	queries *database.Queries
}

func NewAuthorRepository(pool *pgxpool.Pool) *AuthorRepository {
	// fmt.Print(pool.Stat().TotalConns())
	return &AuthorRepository{
		queries: database.New(pool),
	}
}

func (r *AuthorRepository) CreateAuthor(ctx context.Context, arg database.CreateAuthorParams) (*models.AuthorResponse, error) {
	_, err := r.queries.CreateAuthor(ctx, arg)
	if err != nil {
		return nil, err
	}
	return &models.AuthorResponse{Message: "Author created successfully"}, nil
}

func (r *AuthorRepository) ListAuthors(ctx context.Context) ([]database.Author, error) {
	return r.queries.ListAuthors(ctx)
	// res, err := r.queries.ListAuthors(ctx)
	// if err != nil {
	// 	return nil, err
	// }
	// return res, nil
}

func (r *AuthorRepository) GetAuthor(ctx context.Context, arg int64) (database.Author, error) {
	return r.queries.GetAuthor(ctx, arg)
	// _, err := r.queries.GetAuthor(ctx, arg.ID)
	// if err != nil {
	// return nil, err
	// }
	// return &models.AuthorResponse{Message: "Author updated successfully"}, nil
}

// func (r *UserRepository) CreateAuthor(ctx context.Context,  arg database.CreateUserParams) error {
// 	return r.queries.CreateAuthor(ctx,)
// }

// func (r *UserRepository) GetUser(ctx context.Context, id int32) (*models.User, error) {
// 	return r.queries.GetUser(ctx, id)
// }

// func (r *UserRepository) ListUsers(ctx context.Context) ([]models.User, error) {
// 	return r.queries.ListUsers(ctx)
// }

// func (r *UserRepository) UpdateUser(ctx context.Context, id int32, name, email string) error {
// 	return r.queries.UpdateUser(ctx, name, email, id)
// }

// func (r *UserRepository) DeleteUser(ctx context.Context, id int32) error {
// 	return r.queries.DeleteUser(ctx, id)
// }
