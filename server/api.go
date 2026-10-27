package server

import (
	"fmt"
	"net/http"
)

type layuotApi struct {
	url        string
	method     string
	handleFunc func(w http.ResponseWriter, r *http.Request)
}

var api = []layuotApi{
	{
		url:    "/api/v1/health",
		method: "GET",
		handleFunc: func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "text/plain")
			status := http.StatusOK
			w.WriteHeader(status)

			ct := (w.Header().Get("Content-Type"))
			fmt.Fprintf(w, "Content-Type: %v.", ct)
			fmt.Fprintf(w, "\nCode: HTTP %v (%v)\n\n", status, http.StatusText(status))
			fmt.Fprintf(w, "\nBody:\n")
			w.Write([]byte("HEALTHY"))
		},
	},
}
