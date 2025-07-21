package models

type (
	RentTools struct {
		Base
		RentToolsPayload
	}

	RentToolsPayload struct {
		Name     string `json:"name"`
		Size     string `json:"size"`
		Quantity uint   `json:"quantity"`
		UserId   uint   `json:"-"`
	}
)
