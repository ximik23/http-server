package server

import (
	"encoding/json"
	"fmt"
	"http_server/data"
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

			txtBody := fmt.Sprintf("\nBody:\nHealthy")
			textResult(&w, http.StatusOK, "CT", "Code", txtBody)
		},
	},
	{
		url:    "/contacts",
		method: "POST",
		handleFunc: func(w http.ResponseWriter, req *http.Request) {
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
			jsonResult(&w, "Контакт создан!", new_contact)
		},
	},
	{
		url:    "/contacts",
		method: "GET",
		handleFunc: func(w http.ResponseWriter, req *http.Request) {
			contacts := data.GetContacts()
			jsonResult(&w, "", contacts)
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
				txtBody := fmt.Sprintf("\nBody:\n%v", "Контакта с таким Id не существует")
				textResult(&w, http.StatusNotFound, "Code", txtBody)
			} else {
				jsonResult(&w, "", data.GetContactById(int(id)))
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

			textResult(&w, http.StatusNoContent, "Code")
		},
	},
}
