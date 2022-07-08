package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/go-http-utils/headers"
	"hash/fnv"
	"log"
	"net/http"
	"restServer/model"
	"strconv"
)

func DefaultHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Serving:", r.URL.Path, "from", r.Host)
	w.WriteHeader(http.StatusNotFound)
	Body := `<head></head><body>
     <h>Go REST Service Example</h>
	 <p></p>` +
		"<p><a href =\"http://localhost" + model.Port + "/docs\">Link to Swagger</a></p>" +
		"<p><a href =\"http://localhost" + model.Port + "/swagger.yaml\">Link to YAML</a></p>" +
		`</body>`
	_, _ = fmt.Fprintf(w, "%s", Body)
}

func writeJSONValue(w http.ResponseWriter, v any) {
	writeJSONValueWithCode(w, v, http.StatusOK, false)
}

func writeJSONValueHeadersOnly(w http.ResponseWriter, v any) {
	writeJSONValueWithCode(w, v, http.StatusOK, true)
}

func writeJSONValueWithCode(w http.ResponseWriter, v any, code int, headersOnly bool) {
	jsonStr, err := json.Marshal(v)
	if err == nil {
		w.Header().Set(headers.ContentType, "application/json; charset=utf-8")
		w.Header().Set(headers.ETag, hash(jsonStr))
		w.WriteHeader(code)
		if !headersOnly {
			_, err := fmt.Fprint(w, string(jsonStr))
			if err != nil {
				log.Println("Error while writing response:", err)
			}
		}
	} else {
		log.Println(err)
		http.Error(w, "Error:", http.StatusInternalServerError)
	}
}

func faultResponseWithError(w http.ResponseWriter, message string, code int, err error) {
	log.Println(message, "caused by", err)
	errorResponse := model.ErrorResponse{message}
	writeJSONValueWithCode(w, errorResponse, code, false)
}

func faultResponse(w http.ResponseWriter, message string, code int) {
	log.Println(message)
	errorResponse := model.ErrorResponse{message}
	writeJSONValueWithCode(w, errorResponse, code, false)
}

func hash(data []byte) string {
	h := fnv.New64a()
	_, err := h.Write(data)
	if err == nil {
		return strconv.FormatUint(h.Sum64(), 10)
	} else {
		return "-1"
	}
}
