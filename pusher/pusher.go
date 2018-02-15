package pusher

import (
	"time"

	"github.com/kochman/hotshots/config"

	"github.com/kochman/hotshots/log"
)

type Pusher struct {
	cfg config.Config
}

func New(cfg *config.Config) (*Pusher, error) {
	p := &Pusher{
		cfg: *cfg,
	}

	return p, nil
}

func (p *Pusher) Run() {
	ticker := time.Tick(time.Second)
	for range ticker {
		log.Info("Pusher looped")
	}
}
