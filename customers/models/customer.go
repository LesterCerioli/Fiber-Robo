package models

import(
	"github.com/google/uuid"
)

type Customer {
	ID		string	'json:"id"'
	Name	string  'json: "name"'
	Email	string	'json: "email"'
	Phone string `json:"phone"`

}

func NewCustomer(name, email, phone string) *Customer {
    return &Customer{
        ID:    uuid.New().String(),
        Name:  name,
        Email: email,
        Phone: phone,
    }
}

