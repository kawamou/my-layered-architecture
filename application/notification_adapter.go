package application

import "context"

// NotificationAdapter は通知に関するインターフェースです
type NotificationAdapter interface {
	Notify(context.Context, string) error
}
