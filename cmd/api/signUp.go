package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/jcastel2014/quiz3/internal/data"
	_ "github.com/jcastel2014/quiz3/internal/data"
	"github.com/jcastel2014/quiz3/internal/validator"
	_ "github.com/jcastel2014/quiz3/internal/validator"
)

type SignUpModel struct {
	DB *sql.DB
}

func (c SignUpModel) Insert(signUp *data.SignUp) error {
	query := `
		INSERT INTO signup (email, firstName ,middleName ,lastName)
		VALUES($1, $2, $3, $4)
		`

	args := []any{signUp.Email, signUp.FirstName, signUp.MiddleName, signUp.LastName}
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	return c.DB.QueryRowContext(ctx, query, args...).Scan(
		&signUp.ID,
		&signUp.Created_at,
		&signUp.Version,
	)
}

func (a *applicationDependencies) signUp(w http.ResponseWriter, r *http.Request) {
	var incomingData struct {
		Email      string `json:"email"`
		FirstName  string `json:"firstName"`
		MiddleName string `json:"middleName`
		LastName   string `json:"lastName`
	}

	err := a.readJSON(w, r, &incomingData)
	if err != nil {
		a.badRequestResponse(w, r, err)
		return
	}

	signUp := &data.SignUp{
		Email:      incomingData.Email,
		FirstName:  incomingData.FirstName,
		MiddleName: incomingData.MiddleName,
		LastName:   incomingData.LastName,
	}

	v := validator.NEW()

	data.ValidateComment(v, signUp)
	if !v.IsEmpty() {
		a.failedValidationResponse(w, r, v.Errors)
		return
	}

	signUpModel := SignUpModel{DB: a.DB}

	err = signUpModel.Insert(signUp)

	if err != nil {
		log.Println("Error Inserting into Database", err)
	}

	log.Println(signUp)
	fmt.Fprintf(w, "%+v\n", incomingData)
}
