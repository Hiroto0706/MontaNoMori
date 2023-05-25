package main

import (
	"todo/app/handler"
	"todo/app/models"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Message struct {
	Text string `json:"text"`
}

func main() {
	models.InitDB()

	// Echoのインスタンスを作成
	e := echo.New()

	// e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
	// 	AllowOrigins: []string{"http://localhost:8080"},
	// }))

	// // CORS設定を有効にする
	e.Use(middleware.CORS())

	// ルートの定義
	e.GET("/", handler.GetImages)
	e.GET("/search", handler.SearchImages)
	e.GET("/get/:id", handler.GetImage)
	e.GET("/get/next/:id", handler.GetNextImage)
	e.GET("/get/prev/:id", handler.GetPrevImage)
	e.GET("/recommended", handler.GetImages)
	e.GET("/tag/:name", handler.GetImagesByTag)
	e.POST("/login", handler.Login)
	e.POST("/login/check", handler.LoginCheck)

	// admin関係はここ
	admin := e.Group("/admin")
	admin.POST("/upload", handler.UploadImage)
	admin.PUT("/edit/:id", handler.EditImage)
	admin.DELETE("/delete/:id", handler.DeleteImage)
	admin.GET("/tag/get", handler.GetTags)

	tag := admin.Group("/tag")
	tag.POST("/new", handler.CreateTag)
	tag.PUT("/edit", handler.UpdateTag)
	tag.DELETE("/delete/:id", handler.DeleteTag)

	// Start server
	e.Logger.Fatal(e.Start(":8000"))
}
