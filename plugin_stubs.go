//go:build (!linux && !freebsd && !darwin) || !cgo

package nexp_plugin

import "errors"

func lookup(p *Plugin, symName string) (Symbol, error) {
	return nil, errors.New("plugin: not implemented")
}

func open(name string) (*Plugin, error) {
	return nil, errors.New("plugin: not implemented")
}