package client

import (
	"time"
)

type IronfishClient interface {
	Connect(timeout time.Duration) error
	Request(path string, data []byte, timeout time.Duration) ([]byte, error)
	Close() error
}
