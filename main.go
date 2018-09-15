package main

import (
	"log"
	"os"

	"github.com/gameinpl/beerpoly-home/bphapp"
	"github.com/gameinpl/beerpoly-home/modules/goatcms/cmsapp"
	"github.com/goatcms/goatcore/app/bootstrap"
	"github.com/goatcms/goatcore/app/goatapp"
	"github.com/goatcms/goatcore/app/modules/terminal"
)

func main() {
	errLogs := log.New(os.Stderr, "", 0)
	app, err := goatapp.NewGoatApp("BeerpolyHome", "1.0.0", "./")
	if err != nil {
		errLogs.Println(err)
		return
	}
	bootstrap := bootstrap.NewBootstrap(app)
	if err = bootstrap.Register(terminal.NewModule()); err != nil {
		errLogs.Println(err)
		os.Exit(1)
		return
	}
	if err = bootstrap.Register(cmsapp.NewModule()); err != nil {
		errLogs.Println(err)
		os.Exit(1)
		return
	}
	if err = bootstrap.Register(bphapp.NewModule()); err != nil {
		errLogs.Println(err)
		os.Exit(1)
		return
	}
	if err = bootstrap.Init(); err != nil {
		errLogs.Println(err)
		os.Exit(1)
		return
	}
	if err = bootstrap.Run(); err != nil {
		errLogs.Println(err)
		os.Exit(1)
		return
	}
}
