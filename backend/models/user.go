package models

type User struct {
	Base
	Name        string      `json:"name"`
	Address     string      `json:"address"`
	Pre_payment uint        `json:"pre_payment"`
	Active      bool        `json:"active"`
	RentTools   []RentTools `json:"rent_tools"`
}
