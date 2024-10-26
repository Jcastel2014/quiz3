package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/jcastel2014/quiz3/internal/data"
	"github.com/jcastel2014/quiz3/internal/validator"
)

func (c SignUpModel) Patch(signUp *data.SignUp) error {
	query := `
		UPDATE signup
		SET email = $1
		WHERE id = $2
		`

	args := []any{signUp.Email, signUp.Row}
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	return c.DB.QueryRowContext(ctx, query, args...).Scan(
		&signUp.ID,
		&signUp.Created_at,
		&signUp.Version,
	)
}

func (a *applicationDependencies) signUpUpdate(w http.ResponseWriter, r *http.Request) {
	var incomingData struct {
		Email string `json:"email"`
		Row   int    `json:"row"`
	}

	err := a.readJSON(w, r, &incomingData)
	if err != nil {
		a.badRequestResponse(w, r, err)
		return
	}

	update := &data.SignUp{
		Email: incomingData.Email,
		Row:   incomingData.Row,
	}

	v := validator.NEW()

	data.ValidateComment(v, update)
	if !v.IsEmpty() {
		a.failedValidationResponse(w, r, v.Errors)
		return
	}

	signUpModel := SignUpModel{DB: a.DB}

	err = signUpModel.Patch(update)

	if err != nil {
		log.Println("Error Deleting from Database", err)
	}

	log.Println(update)
	fmt.Fprintf(w, "%+v\n", incomingData)

}
