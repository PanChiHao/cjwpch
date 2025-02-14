package app

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go-svc-tpl/app/Controller"
	"go-svc-tpl/app/Controller/LinkControl"
	"go-svc-tpl/app/Controller/UserControl"
	"go-svc-tpl/model"
)

func addRoutes() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/pong", Controller.Pong)
	//重定向
	e.POST("/", LinkControl.Redirect)

	e.POST("/create", LinkControl.CreateLink)
	e.POST("/register", UserControl.Register)
	e.POST("/login", UserControl.Login)

	login := e.Group("/login")
	config := middleware.JWTConfig{
		Claims:      &model.JwtCustomClaims{},
		SigningKey:  []byte("secret"),
		TokenLookup: "header:token",
	}
	login.Use(middleware.JWTWithConfig(config))
	login.POST("/info", UserControl.Info)
	login.POST("/logout", UserControl.LogOut)
	login.POST("/info", UserControl.Info)
	login.POST("/url/get", UserControl.GetAllUrl)

	login.POST("/create", LinkControl.CreateLinkLogin)
	login.POST("/query", LinkControl.QueryLink)
	login.POST("/update", LinkControl.UpdateLink)
	login.POST("/delete", LinkControl.DeleteLink)
	login.POST("/pause", LinkControl.PauseLink)

	e.Logger.Fatal(e.Start(":1232"))
}
