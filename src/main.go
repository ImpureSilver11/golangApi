package main

import (
	"github.com/labstack/echo"
	"./profile"
	"./users"
)

func main() {
	// JSONファイル取り出し
	// raw, err := ioutil.ReadFile("./json/profile/skil.json")
	
	// APIサーバーにする。
	maingPage()
}

func maingPage() {
	// Echoのインスタンス作る
	e := echo.New()
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
}
