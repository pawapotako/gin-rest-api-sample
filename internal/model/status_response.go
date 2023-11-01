package model

type StatusResponse struct {
	IsSuccess bool   `json:"isSuccess"`
	Code      uint   `json:"code"`
	Message   string `json:"message"`
	Data      string `json:"data"`
}
