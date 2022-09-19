package v1

import (
	"backChannel/dto"
	"backChannel/helper"
	"github.com/labstack/echo/v4"
	"net/http"
)

type BackChannelInf interface {
	Create(e echo.Context) error
}

type BackChannelInstance struct {
}

func BackChannelController() BackChannelInf {
	return new(BackChannelInstance)
}

func (c BackChannelInstance) Create(e echo.Context) error {
	//Business Logic
	var x dto.InputFormStr
	err := e.Bind(&x) //?
	if err != nil {
		return e.JSON(http.StatusBadRequest, "bad input data")
	}
	ansString := helper.Reverse(x.InputString)
	return e.JSON(http.StatusOK, ansString)
}
