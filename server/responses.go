package server

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
)

func jsonResult(w *http.ResponseWriter, text string, value any) {
	resText := fmt.Sprintf("%v\n", text)
	(*w).Header().Set("Content-Type", "application/json")
	(*w).WriteHeader(http.StatusCreated)
	(*w).Write([]byte(resText))

	encoder := json.NewEncoder((*w))
	encoder.SetIndent("", "  ")
	err := encoder.Encode(value)
	if err != nil {
		http.Error((*w), "Ошибка сбора данных JSON", http.StatusInternalServerError)
		log.Printf("Ошибка сбора json: ", err)
		return
	}
}

func textResult(w *http.ResponseWriter, status int, text ...string) {
	(*w).Header().Set("Content-Type", "text/plain; charset=utf-8")
	(*w).WriteHeader(status)
	for _, str := range text {
		if strings.Contains(str, "Code") {
			fmt.Fprintf((*w), "Code: HTTP %v (%v)\n", status, http.StatusText(status))
		} else if strings.Contains(str, "CT") {
			ct := ((*w).Header().Get("Content-Type"))
			fmt.Fprintf((*w), "Content-Type: %v.\n", ct)
		} else {
			fmt.Fprintf((*w), str)
		}
	}
}
