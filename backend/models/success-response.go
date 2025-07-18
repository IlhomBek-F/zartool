package models

type SuccessResponseWithMeta[T any] struct {
	Status  int       `json:"status"`
	Message string    `json:"message"`
	Data    T         `json:"data,omitzero"`
	Meta    MetaModel `json:"meta,omitzero"`
}

type SuccessResponseWithData[T any] struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    T      `json:"data,omitzero"`
}

type SuccessResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}
