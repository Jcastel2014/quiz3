package main

import (
	"context"
	"net/http"
	"time"

	"github.com/jcastel2014/quiz3/internal/data"
)

func (a *applicationDependencies) signUpRead(w http.ResponseWriter, r *http.Request) {
	signUpModel := SignUpModel{DB: a.DB}
	query := `
		SELECT id, email, firstName, middleName, lastName
		FROM signup
		`
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	rows, err := signUpModel.DB.QueryContext(ctx, query)

	if err != nil {
		a.badRequestResponse(w, r, err)
		return
	}

	defer rows.Close()

	var signUps []data.SignUp
	for rows.Next() {
		var signUp data.SignUp

		if err := rows.Scan(
			&signUp.ID,
			&signUp.Email,
			&signUp.FirstName,
			&signUp.MiddleName,
			&signUp.LastName,
		); err != nil {
			a.badRequestResponse(w, r, err)
			return
		}
		signUps = append(signUps, signUp)
	}

	if err := rows.Err(); err != nil {
		a.badRequestResponse(w, r, err)
		return
	}

	a.writeJSON(w, http.StatusOK, envelop{"signups": signUps}, nil)

}
