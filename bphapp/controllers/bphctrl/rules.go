package bphctrl

import (
	"html/template"

	"github.com/gameinpl/beerpoly-home/modules/goatcms/cmsapp/services"
	"github.com/gameinpl/beerpoly-home/modules/goatcms/cmsapp/services/requestdep"
	"github.com/goatcms/goatcore/app"
	"github.com/goatcms/goatcore/dependency"
	"github.com/goatcms/goatcore/goathtml"
)

// Rules is rules controller struct
type Rules struct {
	deps struct {
		Template services.Template `dependency:"TemplateService"`
	}
	view *template.Template
}

// NewRules create a new instance of Rules controller
func NewRules(dp dependency.Provider) (*Rules, error) {
	var err error
	ctrl := &Rules{}
	if err = dp.InjectTo(&ctrl.deps); err != nil {
		return nil, err
	}
	ctrl.view, err = ctrl.deps.Template.View(goathtml.DefaultLayout, "custom/home/rules", nil)
	if err != nil {
		return nil, err
	}
	return ctrl, nil
}

// Get is endpoint for HTTP GET request
func (c *Rules) Get(requestScope app.Scope) (err error) {
	var (
		requestDeps struct {
			Responser requestdep.Responser `request:"ResponserService"`
		}
	)
	if err = requestScope.InjectTo(&requestDeps); err != nil {
		return err
	}
	return requestDeps.Responser.Execute(c.view, nil)
}
