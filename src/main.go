package main

import (
	"./profile"
	"./users"
	"github.com/labstack/echo"
)

// "github.com/labstack/echo/middleware"

func main() {
	// APIサーバーにする。
	maingPage()
}

func maingPage() {
	// Echoのインスタンス作る
	e := echo.New()

	// e.Use(middleware.Logger())
	// e.Use(middleware.Recover())
	// Basic認証
	// e.Use(interceptor.BasicAuth())
	// TODO: path綺麗に
	e.GET("/profile/profile", profile.returnProfile())
	e.GET("/profile/career", profile.returnCareer())
	e.GET("/profile/skil", profile.returnSkil())
	// TODO: RDB使ったなんか
	e.GET("/users", users.ReturnUsers())
	// サーバー起動
	e.Start(":8080")
	// e.Run(standard.New(":" + os.Getenv("PORT")))
}
