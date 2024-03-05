package rest

import (
	"echo-api/domain"
	"net/http"
	"os"
	"time"

	"github.com/labstack/echo/v4"
)

type UserService interface {
	SignUp(user domain.User) (domain.UserResponse, error)
	LogIn(user domain.User) (string, error)
}

type UserHandler struct {
	Service UserService
}

func NewUserHandler(e *echo.Echo, us UserService) {
	handler := &UserHandler{
		Service: us,
	}

	e.POST("/signup", handler.SignUp)
	e.POST("/login", handler.LogIn)
	e.POST("/logout", handler.LogOut)
	e.GET("/csrf", handler.CsrfToken)
}

func (u *UserHandler) SignUp(c echo.Context) error {
	user := domain.User{}
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, ResponseError{Message: err.Error()})
	}

	resUser, err := u.Service.SignUp(user)
	if err != nil {
		// TODO: バリデーションエラーの場合は400を返す
		return c.JSON(http.StatusInternalServerError, ResponseError{Message: err.Error()})
	}

	return c.JSON(http.StatusCreated, resUser)
}

func (u *UserHandler) LogIn(c echo.Context) error {
	user := domain.User{}
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, ResponseError{Message: err.Error()})
	}

	token, err := u.Service.LogIn(user)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, ResponseError{Message: err.Error()})
	}

	cookie := new(http.Cookie)
	cookie.Name = "token"
	cookie.Value = token
	cookie.Expires = time.Now().Add(24 * time.Hour)
	cookie.Path = "/"
	cookie.Domain = os.Getenv("API_DOMAIN")
	cookie.Secure = true
	cookie.HttpOnly = true
	cookie.SameSite = http.SameSiteNoneMode
	c.SetCookie(cookie)

	return c.NoContent(http.StatusOK)
}

func (u *UserHandler) LogOut(c echo.Context) error {
	cookie := new(http.Cookie)
	cookie.Name = "token"
	cookie.Value = ""
	cookie.Expires = time.Now()
	cookie.Path = "/"
	cookie.Domain = os.Getenv("API_DOMAIN")
	cookie.Secure = true
	cookie.HttpOnly = true
	cookie.SameSite = http.SameSiteNoneMode
	c.SetCookie(cookie)

	return c.NoContent(http.StatusOK)
}

func (u *UserHandler) CsrfToken(c echo.Context) error {
	token := c.Get("csrf").(string)
	return c.JSON(http.StatusOK, echo.Map{"csrf_token": token})
}
