package responses

import (
	"encoding/json"
	"log"
	"net/http"
)

func JSON(writer http.ResponseWriter, statusCode int, data interface{}) {
	writer.WriteHeader(statusCode)

	if err := json.NewEncoder(writer).Encode(data); err != nil {
		log.Fatal(err)
	}
}

func Error(writer http.ResponseWriter, statusCode int, err error) {
	JSON(writer, statusCode, struct {
		Error string `json:"error"`
	}{
		Error: err.Error(),
	})
}
