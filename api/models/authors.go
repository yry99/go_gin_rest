package models

type AuthorInput struct {
	ID   int64
	Name string `json:"name,omitempty" binding:"required,max=32"`
	Bio  string `json:"bio,omitempty" binding:"required"`
}

type AuthorResponse struct {
	Message string `json:"message"`
}
