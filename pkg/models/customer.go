package models

type Customer struct {
	ID        int
	Name      string `json: "name,omitempty"`
	Role      string `json: "name,omitempty"`
	Email     string `json: "name,omitempty"`
	Phone     string `json: "name,omitempty"`
	Contacted bool   `json: "name,omitempty"`
}
