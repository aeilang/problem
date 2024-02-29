package handler

import (
	"net/http"
	"strconv"
	"time"

	"github.com/a-h/templ"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/lang/problem/internal/repo/gen"
	"github.com/lang/problem/internal/service"
	"github.com/lang/problem/internal/tool"
	"github.com/lang/problem/internal/view/temp"
	"golang.org/x/crypto/bcrypt"
)

type Handler struct {
	serv service.Servicer
}

func NewHandler(serv service.Servicer) Handler {
	return Handler{
		serv: serv,
	}
}

func (h *Handler) Index(c echo.Context) error {
	teml := temp.Index()
	return ServeTemp(c, teml)
}

func (h *Handler) LoginPage(c echo.Context) error {
	temp := temp.Login()
	return ServeTemp(c, temp)
}

func (h *Handler) LoginLogic(c echo.Context) error {
	email := c.FormValue("email")
	password := c.FormValue("password")
	if !h.verifyEmailPass(c, email, password) {
		return c.String(http.StatusBadRequest, "UserName Or Password is Wrong")
	}
	jwtToken, err := tool.GenJwt(email)
	if err != nil {
		return c.String(http.StatusInternalServerError, "faile to gen jwt token")
	}
	cookie := new(http.Cookie)
	cookie.Name = "token"
	cookie.Value = jwtToken
	cookie.Expires = time.Now().Add(72 * time.Hour)
	c.SetCookie(cookie)
	return c.Redirect(302, "/afterlogin")
}

func (h *Handler) AfterLogin(c echo.Context) error {
	email := c.Get("email").(string)

	u, err := h.serv.CheckEmail(c.Request().Context(), email)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	data, err := h.serv.GetProblemsByCreatedBy(c.Request().Context(), u.ID)
	tem := temp.AfterLogin(email, data)
	return ServeTemp(c, tem)
}

func (h *Handler) SignUpPage(c echo.Context) error {
	tem := temp.SignUp()
	return ServeTemp(c, tem)
}

func (h *Handler) SignUpLogic(c echo.Context) error {
	email := c.FormValue("email")
	password := c.FormValue("password")
	if h.HasEmail(c, email) {
		return c.String(http.StatusBadRequest, "email is already used")
	}

	passHash, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	err = h.serv.CreateUser(c.Request().Context(), gen.CreateUserParams{
		ID:           uuid.New(),
		Username:     email,
		Email:        email,
		PasswordHash: string(passHash),
	})
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	return h.LoginPage(c)
}

func (h *Handler) CreateProblem(c echo.Context) error {
	title := c.FormValue("title")
	description := c.FormValue("description")

	email := c.Get("email").(string)
	u, err := h.serv.CheckEmail(c.Request().Context(), email)
	if err != nil {
		c.String(http.StatusInternalServerError, "failed to get user")
	}

	err = h.serv.CreateProblem(c.Request().Context(), gen.CreateProblemParams{
		Title:       title,
		Description: description,
		IsDone:      false,
		CreatedBy:   u.ID,
	})
	if err != nil {
		c.String(http.StatusInternalServerError, "failed to create problem")
	}

	return h.AfterLogin(c)
}

func (h *Handler) CreatePage(c echo.Context) error {
	tem := temp.CreatePage()
	return ServeTemp(c, tem)
}

func (h *Handler) HasEmail(c echo.Context, email string) bool {
	_, err := h.serv.CheckEmail(c.Request().Context(), email)
	return err == nil
}

func (h *Handler) verifyEmailPass(c echo.Context, email, password string) bool {
	u, err := h.serv.CheckEmail(c.Request().Context(), email)
	if err != nil {
		return false
	}

	err = bcrypt.CompareHashAndPassword([]byte(u.PasswordHash), []byte(password))
	return err == nil
}

func (h *Handler) Delete(c echo.Context) error {
	id := c.Param("id")
	email := c.Get("email").(string)
	u, err := h.serv.CheckEmail(c.Request().Context(), email)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	pid, err := strconv.Atoi(id)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	err = h.serv.DeleteProblem(c.Request().Context(), gen.DeleteProblemParams{
		CreatedBy: u.ID,
		ID:        int32(pid),
	})
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	return h.AfterLogin(c)
}

func ServeTemp(c echo.Context, tem templ.Component) error {
	return tem.Render(c.Request().Context(), c.Response().Writer)
}

func Auth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		cookie, err := c.Cookie("token")
		if err != nil {
			return err
		}

		email, err := tool.VerifyToken(cookie.Value)
		if err != nil {
			return err
		}
		c.Set("email", email)
		return next(c)
	}
}
