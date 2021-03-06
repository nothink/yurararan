// yurararan api サーバ
package main

import (
	"fmt"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/nothink/yurararan/handler"
	"github.com/nothink/yurararan/shelf"
)

// メインエントリポイント
func main() {
	// TODO: いずれshelfの初期化周りをよしなに整形すること
	fmt.Println(time.Now().Format("2006/01/02 15:04:05"), " : begin of init a shelf...")
	shelf.Init()
	fmt.Println(time.Now().Format("2006/01/02 15:04:05"), " : end of init a shelf.")

	e := echo.New()

	e.Use(middleware.CORSWithConfig(middleware.DefaultCORSConfig))

	// Resources APIs
	e.GET("/api/v0/resources", handler.GetResources)
	e.POST("/api/v0/resources", handler.PostResources)
	e.GET("/api/v0/cardhashes", handler.GetCardHashes)

	// // Card APIs
	// e.GET("/api/v0/girls", handler.GetGirls)
	// e.POST("/api/v0/girls", handler.PostGirls)
	// e.GET("/api/v0/scenes", handler.GetScenes)
	// e.POST("/api/v0/scenes", handler.PostScenes)
	// e.GET("/api/v0/cards", handler.GetCards)
	// e.POST("/api/v0/cards", handler.PostCards)

	e.Logger.Fatal(e.Start(":1323"))
}
