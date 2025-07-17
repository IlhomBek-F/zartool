package models

type SuccessResponse[T any] struct {
	Status  int       `json:"status"`
	Message string    `json:"message"`
	Data    T         `json:"data,omitzero"`
	Meta    MetaModel `json:"meta,omitzero"`
}
