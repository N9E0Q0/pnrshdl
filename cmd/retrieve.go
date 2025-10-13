package main

import (
	"net/http"

	// aeromexico "github.com/N9E0Q0/pnrshdl/pkg/aeromexico/pnr"
	// aircanada "github.com/N9E0Q0/pnrshdl/pkg/aircanada/pnr"
	delta "github.com/N9E0Q0/pnrshdl/pkg/delta/pnr"
	// united "github.com/N9E0Q0/pnrshdl/pkg/united/pnr"
)

func DeltaRetrieveHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		w.Header().Add("Location", "/delta?error=t")
		w.WriteHeader(302)
		return
	}

	firstName := r.Form.Get("first_name")
	lastName := r.Form.Get("last_name")
	confirmationCode := r.Form.Get("confirmation_code")

	if len(confirmationCode) != 6 || len(firstName) == 0 || len(lastName) == 0 {
		w.Header().Add("Location", "/delta?error=t")
		w.WriteHeader(302)
		return
	}

	retrievedPNR, err := delta.Retrieve(delta.DeltaEndpoint, firstName, lastName, confirmationCode)
	if err != nil {
		w.Header().Add("Location", "/delta?error=t")
		w.WriteHeader(302)
		return
	}

	t := Parse("DLTERM32.html")

	t.Execute(w, struct {
		PNR              delta.PNR
		ConfirmationCode string
		CommitHash       string
	}{
		retrievedPNR,
		confirmationCode,
		commitHash,
	})
}

