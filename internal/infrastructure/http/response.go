package handlers

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
)

type ResponseBody struct {
	Data  interface{} `json:"data"`
	Error interface{} `json:"error"`
}

func SuccessResponse(w http.ResponseWriter, data interface{}, status int) {
	body := ResponseBody{
		Data: data,
	}
	writeResponse(w, body, status)
}

func ErrorResponse(w http.ResponseWriter, err error, status int) {
	body := ResponseBody{
		Error: err.Error(),
	}
	writeResponse(w, body, status)
}

func writeResponse(w http.ResponseWriter, body ResponseBody, status int) {
	w.Header().Set("Content-Type", "application/json")

	var buf bytes.Buffer
	tempEncoder := json.NewEncoder(&buf)
	if err := tempEncoder.Encode(body); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		errorBody := ResponseBody{
			Error: "Failed to encode writeResponse: " + err.Error(),
		}
		if err = json.NewEncoder(w).Encode(errorBody); err != nil {
			log.Printf("Critical: Failed to encode error writeResponse: %v", err)

			w.Header().Set("Content-Type", "text/plain")
			if _, err = w.Write([]byte("Internal server error")); err != nil {
				panic(err)
			}
		}

		return
	}

	w.WriteHeader(status)
	if _, err := w.Write(buf.Bytes()); err != nil {
		panic(err)
	}
}
