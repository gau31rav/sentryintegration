package order

import (
	"github.com/gau31rav/sentryintegration/logging"
	"github.com/google/uuid"
)

// Service ...
type Service interface {
	PlaceOrder() error
}

type service struct {
	logger logging.Logger
}

//PlaceOrder ...
func (s service) PlaceOrder() error {
	s.logger.Error("failed to place order, order id: %s", uuid.NewString())
	return nil
}

func NewService(logger logging.Logger) Service {
	return &service{logger: logger}
}
