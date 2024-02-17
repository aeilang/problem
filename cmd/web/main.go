package main

import (
	"net/http"
	"time"

	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
	"github.com/lang/problem/internal/view/temp"
	_ "github.com/lib/pq"
)

func main() {
	// mux := http.NewServeMux()
	// mux.HandleFunc("/home", func(w http.ResponseWriter, r *http.Request) {
	// 	tem := temp.Index()
	// 	tem.Render(r.Context(), w)
	// })

	// mux.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
	// 	tem := temp.Login()
	// 	tem.Render(r.Context(), w)
	// })
	// http.ListenAndServe(":8888", mux)
	app := echo.New()
	app.GET("/", func(c echo.Context) error {
		tem := temp.Index()
		return ServeTemp(c, tem)
	})
	app.GET("/login", func(c echo.Context) error {
		tem := temp.Login()
		return ServeTemp(c, tem)
	})
	app.POST("/login", func(c echo.Context) error {
		email := c.FormValue("email")
		password := c.FormValue("password")
		cookie := new(http.Cookie)
		cookie.Name = "token"
		cookie.Value = email + password
		cookie.Expires = time.Now().Add(24 * time.Second)
		c.SetCookie(cookie)
		tem := temp.Nav()
		return ServeTemp(c, tem)
	})

	app.Logger.Fatal(app.Start(":8888"))
}

func ServeTemp(c echo.Context, tem templ.Component) error {
	return tem.Render(c.Request().Context(), c.Response().Writer)
}
