package main

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/kawamou/my-layered-architecture/application"
	"github.com/kawamou/my-layered-architecture/handler"
	"github.com/kawamou/my-layered-architecture/infra"
)

// Router
type Router struct {
	Engine *gin.Engine
}

// NewRouter ...
func NewRouter() *Router {
	engine := gin.New()
	return &Router{
		Engine: engine,
	}
}

// Run ...
func (r *Router) Run(port string) error {
	ctx := context.Background()
	f := infra.NewFirestoreClient(ctx)
	pingHandler := handler.NewPingHandler(*application.NewPingService(infra.NewSlackLocalNotifier()))
	createUserHandler := handler.NewCreateUserHandler(*application.NewCreateUserService(infra.NewUserRepositoryImpl(ctx, f)))
	findUserHandler := handler.NewFindUserHandler(*application.NewFindUserService(infra.NewUserRepositoryImpl(ctx, f)))
	api := r.Engine.Group("")
	api.GET("/ping", pingHandler.Handle)          // ヘルスチェック
	api.POST("/users", createUserHandler.Handle)  // 作成
	api.GET("/users/:id", findUserHandler.Handle) // 閲覧
	// api.POST("/users/:id") // 更新
	// api.POST("/users/:id/delete") // 削除
	// api.GET("/users/:id") // 閲覧
	// api.GET("/users") // 全件閲覧、またはクエリパラメータ？
	return r.Engine.Run(port)
}
