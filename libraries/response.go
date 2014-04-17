package lib

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Response struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

func QuickResponse(w http.ResponseWriter, status string, message string) {
	e := Response{
		Status:  status,
		Message: message,
	}

	bytes, err := json.MarshalIndent(e, "", "\t")
	if err != nil {
		log.Printf("%s", err)
	}

	fmt.Fprintf(w, "%s", bytes)
}
