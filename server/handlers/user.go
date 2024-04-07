package handlers

import (
	"net/http"
	"small_go_project/server/httpmodels"
	"strconv"

	"github.com/labstack/echo/v4"
)

func (h *Handlers) CreateUser(c echo.Context) error {
	h.log.Info("called CreateBook handler")

	req := new(httpmodels.CreateUserRequest)

	err := c.Bind(req)
	if err != nil {
		h.log.Error("failed to parse request body", err)
		return c.JSON(http.StatusBadRequest, httpmodels.IncorrectReqBodyErr)
	}

	err = req.Validate()
	if err != nil {
		h.log.Error("validation failed: ", err)
		return c.JSON(http.StatusBadRequest, httpmodels.ValidationFailedErr)
	}

	return c.NoContent(http.StatusCreated)
}

func (h *Handlers) GetUser(c echo.Context) error {
	h.log.Info("called GetUser handler")
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		h.log.Error("wrong parameter: ", err)
		return c.JSON(http.StatusBadRequest, httpmodels.IncorrectParameter)
	}

	resp := httpmodels.GetUserRequest{
		Name:    "Ivan",
		Surname: "Ivanov",
		Age:     21,
		Id:      id,
	}

	return c.JSON(http.StatusOK, resp)
}

func (h *Handlers) DeleteUser(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		h.log.Error("wrong parameter: ", err)
		return c.JSON(http.StatusBadRequest, httpmodels.IncorrectParameter)
	}
	h.log.Infof("user %d deleted", id)
	return c.NoContent(http.StatusAccepted)
}

func (h *Handlers) UpdateUser(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		h.log.Error("wrong parameter: ", err)
		return c.JSON(http.StatusBadRequest, httpmodels.IncorrectParameter)
	}
	req := new(httpmodels.CreateUserRequest)

	err = c.Bind(req)
	if err != nil {
		h.log.Error("failed to parse request body", err)
		return c.JSON(http.StatusBadRequest, httpmodels.IncorrectReqBodyErr)
	}

	err = req.Validate()
	if err != nil {
		h.log.Error("validation failed: ", err)
		return c.JSON(http.StatusBadRequest, httpmodels.ValidationFailedErr)
	}
	h.log.Infof("user id: %d updated", id)

	return c.JSON(http.StatusAccepted, req)
}
