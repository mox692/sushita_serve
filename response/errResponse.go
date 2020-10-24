package response

import (
	"fmt"
	"net/http"
)

// Todo: errorが増えた時に、ここで切り分け。
func ErrResponse(w http.ResponseWriter, err error, code int) {
	switch err.(type) {
	default:
		{
			writeErr(w, err, code)
		}
	}
}

func writeErr(w http.ResponseWriter, err error, code int) {
	fmt.Fprintf(w, "%+v\n", err)
	w.WriteHeader(code)
}
