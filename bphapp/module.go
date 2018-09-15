package bphapp

import (
	"github.com/goatcms/goatcore/app"
)

// Module is slots app module structure
type Module struct{}

// NewModule create new module instance
func NewModule() app.Module {
	return &Module{}
}

// RegisterDependencies is init callback to register module dependencies
func (m *Module) RegisterDependencies(a app.App) (err error) {
	return nil
}

// InitDependencies is init callback to inject dependencies inside modules
func (m *Module) InitDependencies(a app.App) error {
	if err := InitControllers(a); err != nil {
		return err
	}
	return nil
}

// Run is run event callback
func (m *Module) Run() error {
	return nil
}
