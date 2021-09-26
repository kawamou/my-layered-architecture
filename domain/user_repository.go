package domain

import (
	"context"
)

// UserRepository はUserRepositoryのインタフェースです
type UserRepository interface {
	Add(ctx context.Context, user User) error               // 作成
	FindById(ctx context.Context, id string) (*User, error) // 取得
	// Update(ctx context.Context, id string, user User) error // 更新
	// FindByName(ctx context.Context, name string) (User, error) // 取得
	// FindAll(ctx context.Context) ([]User, error) //全件取得
}
