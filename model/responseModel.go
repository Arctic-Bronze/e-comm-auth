package model

import (
	"github.com/gofiber/fiber/v2"
	"time"
)

type ErrorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Time    int64  `json:"time"`
}

func (resp *ErrorResponse) CreateErrorResponse(err *fiber.Error) *ErrorResponse {
	resp.Code = err.Code
	resp.Message = err.Message
	resp.Time = time.Now().Unix()
	return resp
}

type SliceResponse struct {
	Status int          `json:"status"`
	Time   int64        `json:"time"`
	Count  int          `json:"count"`
	Data   *interface{} `json:"data"`
}
