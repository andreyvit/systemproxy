package systemproxy

import (
	"errors"
)

type Settings struct {
	Enabled bool
	Server  string
}

var ErrNotImpl = errors.New("systemproxy not implemented on this platform")
