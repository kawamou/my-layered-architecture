package application

import (
	"context"

	"github.com/pkg/errors"
)

type PingService struct {
	n NotificationAdapter
}

func NewPingService(n NotificationAdapter) *PingService {
	return &PingService{
		n: n,
	}
}

func (p *PingService) Handle(ctx context.Context) error {
	if err := p.n.Notify(ctx, "pong"); err != nil {
		return errors.Wrap(err, "")
	}
	return nil
}
