package data

import (
	"github.com/jcastel2014/quiz3/internal/validator"
)

type SignUp struct {
	ID         int64
	Email      string
	FirstName  string
	MiddleName string
	LastName   string
	Created_at string
	Version    string
	Row        int
}

func ValidateComment(v *validator.Validator, comment *SignUp) {

	if comment.Row != 0 {
		v.Check(comment.Row != 0, "Row", "number must be provided")
		v.Check(comment.Row < 100, "Row", "integer too big")
	} else {
		v.Check(comment.Email != "", "content", "must be provided")
		v.Check(comment.FirstName != "", "author", "must be provided")
		v.Check(comment.MiddleName != "", "middleName", "must be provided")
		v.Check(comment.LastName != "", "lastName", "must be provided")
		v.Check(len(comment.Email) <= 100, "content", "must not be more than 100 bytes long")
		v.Check(len(comment.FirstName) <= 100, "author", "must not be more than 100 butes longs")
		v.Check(len(comment.MiddleName) <= 100, "middleName", "must not be more than 100 butes longs")
		v.Check(len(comment.LastName) <= 100, "lastName", "must not be more than 100 butes longs")
	}
}
