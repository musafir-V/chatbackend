package internal

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func WriteJson(w http.ResponseWriter, jsonData interface{}, status int) {
	mJson, err := json.Marshal(jsonData)
	if err != nil {
		fmt.Println("Error marshalling JSON:", err)
	} else {
		w.Header()["Content-Type"] = []string{"application/json"}
		w.WriteHeader(status)
		w.Write(mJson)
	}
}
