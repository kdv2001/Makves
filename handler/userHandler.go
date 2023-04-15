package handler

import (
	"Makves/appErrors"
	"encoding/json"
	"errors"
	"strconv"
	"strings"

	"Makves/usecase"
	"github.com/gofiber/fiber/v2"
)

type userHandler struct {
	userUseCase usecase.UserUseCase
}

func NewHandler(userUseCase usecase.UserUseCase) userHandler {
	return userHandler{userUseCase: userUseCase}
}

// GetUserByIds godoc
// @Summary      get items by ids
// @Description  get items by ids
// @Tags         user
// @Produce      json
// @Param        id   query []int  true "items ids"
// @Success      200  {object} []model.User
// @Router       /get-items [get]
func (h userHandler) GetUserByIds(ctx *fiber.Ctx) error {
	queryParam := ctx.Query("id")
	userIds := strings.Split(queryParam, ",")

	ids := make([]int64, 0, len(userIds))
	for _, id := range userIds {
		intId, err := strconv.ParseInt(id, 10, 64)
		if err != nil {
			return fiber.ErrInternalServerError
		}

		ids = append(ids, intId)
	}

	result, err := h.userUseCase.GetUsersByIds(ids)
	if errors.Is(err, appErrors.ErrItemNotFound) {
		return fiber.ErrBadRequest
	} else if err != nil {
		return err
	}

	byteResult, err := json.Marshal(result)
	if err != nil {
		return fiber.ErrInternalServerError
	}

	ctx.Context().Success("application:json", byteResult)

	return nil
}
