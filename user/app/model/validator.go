package model

import (
	"encoding/json"
	"net/http"
)

func DecodeAndValidate(r *http.Request, v InputValidation) error {
	// json decode the payload - obviously this could be abstracted
	// to handle many content types
	if err := json.NewDecoder(r.Body).Decode(v); err != nil {
		return err
	}
	defer r.Body.Close()
	// peform validation on the InputValidation implementation
	return v.Validate(r)
}