package data

import (
	"fmt"
	"time"
)

type Contact struct {
	Id          int       `json:"id"`
	First_name  string    `json:"first_name"`
	Middle_name string    `json:"middle_name"`
	Last_name   string    `json:"last_name"`
	Birthday    time.Time `json:"birthday"`
	Description string    `json:"description"`
}

var contactMap map[int]Contact = make(map[int]Contact)

func createIdForContact() int {
	var idUser = 0
	if len(contactMap) == 0 {
		idUser = 1
	} else {
		var idsList []int
		for key := range contactMap {
			idsList = append(idsList, key)
		}
		idUser = idsList[len(idsList)-1] + 1
	}
	return idUser
}

func existInContacts(c Contact) bool {
	var exist bool
	for _, contact := range contactMap {
		if contact.First_name == c.First_name &&
			contact.Last_name == c.Last_name &&
			contact.Middle_name == c.Middle_name {
			exist = true
			break
		}
	}
	return exist
}

func AddContact(contact Contact) (error, Contact) {
	var err error

	if len(contact.First_name) < 1 ||
		len(contact.Last_name) < 1 ||
		len(contact.Middle_name) < 1 ||
		contact.Birthday.IsZero() {
		err = fmt.Errorf("Одно из полей не заполнено")
		return err, Contact{}
	}

	if existInContacts(contact) {
		err = fmt.Errorf("Такой контакт уже существует!")
		return err, Contact{}
	}

	contact.Id = createIdForContact()
	contactMap[contact.Id] = contact
	return nil, contact
}

func GetContacts() []Contact {
	var listContact []Contact
	for _, contact := range contactMap {
		listContact = append(listContact, contact)
	}
	return listContact
}

func GetContactById(id int) Contact {
	if contact, ok := contactMap[id]; ok {
		return contact
	}
	return Contact{}
}

func DeleteContact(id int) (err error) {
	if _, ok := contactMap[id]; ok {
		delete(contactMap, id)
		err = nil
	} else {
		err = fmt.Errorf("Такой контакта на существует!")
	}
	return
}
