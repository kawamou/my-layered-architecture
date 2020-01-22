package infrastructure

import (
	"github.com/labstack/echo"
	"go-cleanarchitecture-restapi/interfaces/controllers"
)

func Init() {
	e := echo.New()

	memoController := controllers.NewMemoController(NewSqlHandler())

	e.GET("/memos", func(c echo.Context) error {return memoController.Index(c)})
	e.GET("/memo/:id", func(c echo.Context) error { return memoController.Show(c) })
	e.POST("/create", func(c echo.Context) error { return memoController.Create(c) })

	e.Logger.Fatal(e.Start(":8888"))
}

