package controllers

import (
	"context"
	"net/http"
	"sqlc/api/models"
	"sqlc/api/repository"

	"sqlc/internal/database"

	"github.com/gin-gonic/gin"
)

type AuthorController struct {
	repo *repository.AuthorRepository
}

func NewAuthorController(repo *repository.AuthorRepository) *AuthorController {
	return &AuthorController{
		repo: repo,
	}
}

func (c *AuthorController) CreateAuthor(ctx *gin.Context) {
	var input models.AuthorInput
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	params := database.CreateAuthorParams{
		Name: input.Name,
		Bio:  input.Bio,
	}

	author, err := c.repo.CreateAuthor(context.Background(), params)
	if err != nil {
		ctx.JSON(http.StatusServiceUnavailable, gin.H{"error": err.Error()})
		return
	}

	var response string = author.Message
	// response, err := fromDB(author)
	// if err != nil {
	// 	ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	// 	return
	// }
	ctx.JSON(http.StatusOK, response)

}

func (c *AuthorController) ListAuthors(ctx *gin.Context) {
	authors, err := c.repo.ListAuthors(context.Background())
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, authors)

}

// func fromDB(author database.Author) (*models.AuthorResponse, error) {

// 	if author.ID != 0 {
// 		return &models.AuthorResponse{
// 			Message: "Author created successfully",
// 		}, nil
// 	} else {
// 		return nil, errors.New("author not created")
// 	}
// }
