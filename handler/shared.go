package handler

import (
	"context"

	"github.com/gin-gonic/gin"
)

// withContext はGinのContextからcontext.Contextを取り出して後続の関数に渡します
func withContext(c *gin.Context, f func(ctx context.Context)) {
	ctx := c.Request.Context()
	f(ctx)
}
