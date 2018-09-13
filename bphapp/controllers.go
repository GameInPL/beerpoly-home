package bphapp

import (
	"github.com/gameinpl/beerpoly-home/bphapp/controllers/bphctrl"
	"github.com/goatcms/goatcore/app"
)

// InitControllers add controllers to application
func InitControllers(a app.App) (err error) {
	return bphctrl.InitDependencies(a)
}
