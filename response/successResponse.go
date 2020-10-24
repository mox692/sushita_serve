package response

import (
	"encoding/json"
	"net/http"

	"golang.org/x/xerrors"
)

func SuccessResponse(w http.ResponseWriter, response interface{}) {
	data, err := json.Marshal(response)
	if err != nil {
		ErrResponse(w, xerrors.Errorf(" :%w", err), 500)
		return
	}
	w.Write(data)
}
