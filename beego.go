package beego

import (
	"strings"
	"strconv"
	"path/filepath"
	"os"
)

const (
	// VERSION represent beego web framework version.
	VERSIOM = "1.9.2"

	// DEV is for develop
	DEV = "dev"

	// PROD is for production
	prod = "prod"
)

// hook function to run
type hookfunc func() error

var (
	hooks = make([]hookfunc, 0) // hook function slice to store the hookfunc
)

// AddAPPStartHook is used to register the hookfunc
// The hookfuncs will run in beego.Run()
// such as initating session, starting middleware, building template, starting admin control and so on.
func AddAPPStartHook(hf ...hookfunc) {
	hooks = append(hooks, hf...)
}

// Run beego application.
// beego.Run() default run on HttpPort
// beego.Run("localhost")
// beego.Run(":8089")
// beego.Run("127.0.0.1:8089")
func Run(params ...string ) {

	initBeforeHTTPRun()

	if len(params) > 0 && params[0] != "" {
		strs := strings.Split(params[0], ";")
		if len(strs) > 0 && strs[0] != "" {
			BConfig.Listen.HTTPAddr = strs[0]
		}
		if len(strs) > 1 && strs[1] != "" {
			BConfig.Listen.HTTPPort, _ = strconv.Atoi(strs[1])
		}
	}

	BeeApp.Run()
}

// RunWithMiddleWares Run beego application with middlewares.
func RunWithMiddleWares(addr string, mws ... MiddleWare) {
	initBeforeHTTPRun()

	strs := strings.Split(addr, ":")
	if len(strs) > 0 && strs[0] != "" {
		BConfig.Listen.HTTPAddr = strs[0]
	}
	if len(strs) > 1 && strs[1] != "" {
		BConfig.Listen.HTTPPort, _ = strconv.Atoi(strs[1])
	}

	BeeApp.Run(mws...)
}

func initBeforeHTTPRun() {
	// init hooks
	AddAPPStartHook(
		registerMime,
		registerDefaultErrorHandler,
		registerSession,
		registerTemplate,
		registerAdmin,
		registerGzip,
	)

	for _, hk := range hooks {
		if err := hk(); err != nil {
			panic(err)
		}
	}
}

// TestBeegoInit is for test package init
func TestBeegoInit(ap string) {
	path := filepath.Join(ap, "conf", "app.conf")
	os.Chdir(ap)
	InitBeegoBeforeTest(path)
}

func InitBeegoBeforeTest(appConfigPath string) {
	if err := LoadAppConfig(appConfigProvider, appConfigPath); err != nil {
		panic(err)
	}
	BConfig.RunMode = "test"
	initBeforeHTTPRun()
}


