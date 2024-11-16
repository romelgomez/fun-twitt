package helper

import (
	"encoding/json"
	"net/http"
)

type WebResponse struct {
	Code    int         `json:"code"`
	Status  string      `json:"status"`
	Data    interface{} `json:"data,omitempty"`
	Message string      `json:"message"`
}

func ReadRequestBody(r *http.Request, result interface{}) error {
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(result); err != nil {
		return err
	}
	return nil
}

func WriteResponseBody(write http.ResponseWriter, response interface{}) {
	write.Header().Add("Content-Type", "application/json")
	encoder := json.NewEncoder(write)
	err := encoder.Encode(response)
	if err != nil {
		panic(err)
	}
}

func WriteErrorResponse(writer http.ResponseWriter, message string, code int) {
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(code)
	json.NewEncoder(writer).Encode(map[string]string{
		"error": message,
	})
}

func Response(writer http.ResponseWriter, result interface{}, err error) {
	if err != nil {
		WriteErrorResponse(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	webResponse := WebResponse{
		Code:   200,
		Status: "OK",
		Data:   result,
	}

	WriteResponseBody(writer, webResponse)
}
