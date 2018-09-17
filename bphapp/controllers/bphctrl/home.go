package bphctrl

import (
	"html/template"

	"github.com/gameinpl/beerpoly-home/modules/goatcms/cmsapp/services"
	"github.com/gameinpl/beerpoly-home/modules/goatcms/cmsapp/services/requestdep"
	"github.com/goatcms/goatcore/app"
	"github.com/goatcms/goatcore/dependency"
	"github.com/goatcms/goatcore/goathtml"
)

// Home is home controller struct
type Home struct {
	deps struct {
		Template       services.Template `dependency:"TemplateService"`
		OpenSourceLink string            `config:"menu.OpenSourceLink"`
		GameLink       string            `config:"menu.GameLink"`
		DownloadLink   string            `config:"menu.DownloadLink"`
	}
	view *template.Template
}

// NewHome create a new instance of Home controller
func NewHome(dp dependency.Provider) (*Home, error) {
	var err error
	ctrl := &Home{}
	if err = dp.InjectTo(&ctrl.deps); err != nil {
		return nil, err
	}
	ctrl.view, err = ctrl.deps.Template.View(goathtml.DefaultLayout, "custom/home/main", nil)
	if err != nil {
		return nil, err
	}
	return ctrl, nil
}

// Get is endpoint for HTTP GET request
func (c *Home) Get(requestScope app.Scope) (err error) {
	var (
		requestDeps struct {
			Responser requestdep.Responser `request:"ResponserService"`
		}
	)
	if err = requestScope.InjectTo(&requestDeps); err != nil {
		return err
	}
	return requestDeps.Responser.Execute(c.view, map[string]string{
		"OpenSourceLink": c.deps.OpenSourceLink,
		"GameLink":       c.deps.GameLink,
		"DownloadLink":   c.deps.DownloadLink,
	})
}
