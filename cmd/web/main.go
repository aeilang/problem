package main

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/lang/problem/internal/handler"
	"github.com/lang/problem/internal/repo"
	"github.com/lang/problem/internal/repo/gen"
	"github.com/lang/problem/internal/service"
	_ "github.com/lib/pq"
)

func main() {
	app := echo.New()
	dbString := LoadConfig()
	srv := Prepare(dbString)
	app.Use(middleware.Logger())
	app.Use(middleware.Recover())

	app.GET("/", srv.Index)
	app.GET("/login", srv.LoginPage)
	app.POST("/login", srv.LoginLogic)
	app.GET("/signup", srv.SignUpPage)
	app.POST("/signup", srv.SignUpLogic)

	app.GET("/afterlogin", srv.AfterLogin, handler.Auth)
	app.GET("/create", srv.CreatePage, handler.Auth)
	app.POST("/create", srv.CreateProblem, handler.Auth)
	app.DELETE("/delete/:id", srv.Delete, handler.Auth)

	app.Logger.Fatal(app.Start(":8888"))
}

func LoadConfig() string {
	err := godotenv.Load()
	if err != nil {
		panic("failed to load .env file")
	}
	dbString := os.Getenv("POSTGRES")
	return dbString
}

func Prepare(dbString string) *handler.Handler {
	db := repo.NewDB(dbString)
	querys := gen.New(db)
	serv := service.NewService(querys)
	srv := handler.NewHandler(serv)
	return &srv
}
