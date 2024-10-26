package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/jcastel2014/quiz3/internal/data"
)

func (c SignUpModel) Delete(signUp *data.SignUp) error {
	query := `
		DELETE FROM signup
		WHERE id = $1
		
		`

	args := []any{signUp.Row}
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	return c.DB.QueryRowContext(ctx, query, args...).Scan(
		&signUp.ID,
		&signUp.Created_at,
		&signUp.Version,
	)
}

func (a *applicationDependencies) signUpDelete(w http.ResponseWriter, r *http.Request) {
	var incomingData struct {
		Row int `json:"delete"`
	}

	err := a.readJSON(w, r, &incomingData)
	if err != nil {
		a.badRequestResponse(w, r, err)
		return
	}

	delete := &data.SignUp{
		Row: incomingData.Row,
	}

	// v := validator.NEW()

	// data.ValidateComment(v, delete)
	// if !v.IsEmpty() {
	// 	a.failedValidationResponse(w, r, v.Errors)
	// 	return
	// }

	signUpModel := SignUpModel{DB: a.DB}

	err = signUpModel.Delete(delete)

	if err != nil {
		log.Println("Error Deleting from Database", err)
	}

	log.Println(delete)
	fmt.Fprintf(w, "%+v\n", incomingData)

}
