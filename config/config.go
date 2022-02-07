package config

import (
	"fmt"
	"strings"
)

// Loader is the interface that all config managers should follow.
type Loader interface {
	LoadConfig(out interface{}) error
}

// LoadWLFromString is used to convert a comma separated list of strings into a list of strings
func LoadWLFromString(wl string) []string {
	if wl == "" {
		return []string{}
	}

	return strings.Split(wl, ",")
}

// ErrConfigNotFound is used when a config manager is unable to locate the config.
// This should be the case whether it is stored in a file or on a remote service.
type ErrConfigNotFound struct {
	File string
	Err  error
}

func (e *ErrConfigNotFound) Error() string {
	return fmt.Sprintf("config not found, %s: %s", e.File, e.Err.Error())
}

func (e *ErrConfigNotFound) Unwrap() error { return e.Err }

type ErrInvalidConfig struct {
	Err error
}

func (e *ErrInvalidConfig) Error() string {
	return fmt.Sprintf("invalid config: %v", e.Err.Error())
}

func (e *ErrInvalidConfig) Unwrap() error { return e.Err }
