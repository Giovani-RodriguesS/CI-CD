package utils

import (
	"errors"
	"github.com/gin-gonic/gin"
	"strconv"
)

func IdValidate(ctx *gin.Context) (uint, error) {
	Rawid := ctx.Param("Id")

	if Rawid == "" {
		return 0, errors.New("id não deve ser nulo")
	}

	Id, err := strconv.Atoi(Rawid)

	if err != nil {
		return 0, errors.New("id deve ser inteiro")
	}

	if Id < 1 {
		return 0, errors.New("id deve ser positivo")
	}

	uId := uint(Id)
	return uId, nil
}
