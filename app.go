package beego

import (
	"net/http"
	"fmt"
	"net"
	"net/http/fcgi"
	"源码/beego/logs"
)

var (
	// BeeApp is an application instance
	BeeApp *App
)

type App struct {
	Handlers *ControllerRegister
	Server	 *http.Server
}

// MiddWare function for http.Handler
type MiddleWare func(http.Handler) http.Handler

func (app *App) Run(mws ...MiddleWare) {

}