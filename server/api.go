package server

import (
	"encoding/json"
	"fmt"
	"http_server/data"
	"log"
	"net/http"
	"strconv"
)

type layuotApi struct {
	url        string
	method     string
	handleFunc func(w http.ResponseWriter, req *http.Request)
}

var api = []layuotApi{
	{
		url:    "/health",
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
	{
		url:    "/contacts",
		method: "POST",
		handleFunc: func(w http.ResponseWriter, req *http.Request) {
			w.Header().Set("Content-Type", "application/json")

			var contact = data.Contact{}

			err := json.NewDecoder(req.Body).Decode(&contact)
			if err != nil {
				http.Error(w, "Некорректные данные: "+err.Error(), http.StatusBadRequest)
				return
			}

			err, new_contact := data.AddContact(contact)
			if err != nil {
				http.Error(w, "Ошибка создания сонтакта: "+err.Error(), http.StatusInternalServerError)
				return
			}
			w.WriteHeader(http.StatusCreated)
			w.Write([]byte("Контакт создан!\n"))

			encoder := json.NewEncoder(w)
			encoder.SetIndent("", "  ")
			err = encoder.Encode(new_contact)

			if err != nil {
				http.Error(w, "Ошибка сбора данных JSON", http.StatusInternalServerError)
				log.Printf("Ошибка сбора json: ", err)
				return
			}
		},
	},
	{
		url:    "/contacts",
		method: "GET",
		handleFunc: func(w http.ResponseWriter, req *http.Request) {
			w.Header().Set("Content-Type", "application/json")

			contacts := data.GetContacts()

			err := json.NewEncoder(w).Encode(contacts)
			if err != nil {
				http.Error(w, "Ошибка сбора данных JSON", http.StatusInternalServerError)
				log.Printf("Ошибка сбора json: ", err)
				return
			}
		},
	},
	{
		url:    "/contacts/{id}",
		method: "GET",
		handleFunc: func(w http.ResponseWriter, req *http.Request) {

			id, err := strconv.Atoi(req.PathValue("id"))
			if err != nil {
				http.Error(w, "Невозможно взять пользователя по такому id", http.StatusInternalServerError)
				return
			}

			contact := data.GetContactById(int(id))
			if contact.Id == 0 {
				w.Header().Set("Content-Type", "text/plain; charset=utf-8")
				w.WriteHeader(http.StatusNotFound)
				fmt.Fprintf(w, "Code: HTTP %v (%v)\n\n", http.StatusNotFound, http.StatusText(http.StatusNotFound))
				fmt.Fprintf(w, "\nBody:\n")
				w.Write([]byte("Контакта с таким Id не существует"))
			} else {
				w.Header().Set("Content-Type", "application/json")
				encoder := json.NewEncoder(w)
				encoder.SetIndent("", "  ")
				encoder.Encode(data.GetContactById(int(id)))
			}
		},
	},
	{
		url:    "/contacts/{id}",
		method: "DELETE",
		handleFunc: func(w http.ResponseWriter, req *http.Request) {
			id, err := strconv.Atoi(req.PathValue("id"))
			if err != nil {
				http.Error(w, "Невозможно удалить пользователя по такому id", http.StatusInternalServerError)
				return
			}

			err = data.DeleteContact(id)
			if err != nil {
				http.Error(w, "", http.StatusNotFound)
				return
			}

			w.Header().Set("Content-Type", "text/plain; charset=utf-8")
			w.WriteHeader(http.StatusNoContent)
			fmt.Fprintf(w, "Code: HTTP %v (%v)\n\n", http.StatusNoContent, http.StatusText(http.StatusNoContent))
		},
	},
}
