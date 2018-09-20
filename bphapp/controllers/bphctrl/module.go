package bphctrl

import (
	"github.com/gameinpl/beerpoly-home/modules/goatcms/cmsapp/services"
	"github.com/goatcms/goatcore/app"
)

// InitDependencies init all dependency modules
func InitDependencies(a app.App) (err error) {
	var (
		deps struct {
			Router services.Router `dependency:"RouterService"`
		}
		rules *Rules
		home  *Home
	)
	if err = a.DependencyProvider().InjectTo(&deps); err != nil {
		return err
	}
	// home
	if home, err = NewHome(a.DependencyProvider()); err != nil {
		return err
	}
	deps.Router.OnGet("/", home.Get)
	// rules
	if rules, err = NewRules(a.DependencyProvider()); err != nil {
		return err
	}
	deps.Router.OnGet("/rules", rules.Get)
	return nil
}
