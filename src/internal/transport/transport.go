package transport

import (
	"log"
	"net/http"
	"tcc-test/internal/entity"
	"tcc-test/internal/usecase"

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

func (usecase userTransport) CreateUser(c echo.Context) error {
	req := entity.User{}
	err := c.Bind(&req)
	log.Println("req", req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "The information entered is incomplete. Please enter again.")
	}

	return c.JSON(http.StatusCreated, "user")
}

func (usecase userTransport) GetUsers(c echo.Context) (err error) {
	users, err := usecase.usecase.GetUsers()
	if err != nil {
		return
	}
	return c.JSON(http.StatusCreated, users)
}

func (usecase userTransport) GetUser(c echo.Context) error {
	req := entity.User{}
	err := c.Bind(&req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "The information entered is incomplete. Please enter again.")
	}

	return c.JSON(http.StatusCreated, "user")
}
