package utils

import (
	"e-comm-auth/model"
	"github.com/gofiber/fiber/v2"
	"reflect"
	"time"
)

func SendSliceResponse(ctx *fiber.Ctx, data interface{}) {
	slice := reflect.ValueOf(data)
	if slice.Len() == 0 {
		ctx.Response().SetStatusCode(204)
		return
	}
	sliceResp := &model.SliceResponse{
		Status: fiber.StatusOK,
		Time:   time.Now().Unix(),
		Count:  slice.Len(),
		Data:   &data,
	}
	ctx.JSON(sliceResp)
}

func SendStructResponse(ctx *fiber.Ctx, model interface{}, code int) {
	ctx.Response().SetStatusCode(code)
	ctx.JSON(model)
}
