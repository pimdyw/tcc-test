package transport

import (
	"net/http"
	"tcc-test/internal/entity"
	"tcc-test/internal/usecase"
	"time"

	"github.com/labstack/echo/v4"
)

type userTransport struct {
	usecase usecase.UserUsecaseInterface
}

func NewUserTransport(
	usecase usecase.UserUsecaseInterface,
) *userTransport {
	return &userTransport{
		usecase: usecase,
	}
}

func (usecase userTransport) CreateUser(c echo.Context) (err error) {
	req := entity.UserRequest{}
	err = c.Bind(&req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "The information entered is incomplete. Please enter again.")
	}

	// convert time from fe
	layout := "2006-01-02"
	birthTime, err := time.Parse(layout, req.BirthDate)
	if err != nil {
		return
	}

	user := entity.User{
		ID:        req.ID,
		Name:      req.Name,
		Surname:   req.Surname,
		Address:   req.Address,
		BirthDate: birthTime,
		Age:       req.Age,
	}
	err = usecase.usecase.CreateUser(c.Request().Context(), user)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Create user failed.")
	}

	return c.JSON(http.StatusCreated, "Create user success")
}

func (usecase userTransport) GetUsers(c echo.Context) (err error) {
	users, err := usecase.usecase.GetUsers(c.Request().Context())
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Get users failed.")
	}
	return c.JSON(http.StatusCreated, users)
}

func (usecase userTransport) GetUser(c echo.Context) error {
	id := c.Param("id")
	user, err := usecase.usecase.GetUser(c.Request().Context(), id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Get user failed.")
	}

	return c.JSON(http.StatusCreated, user)
}
