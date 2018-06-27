package beego

import (
	"Test/beego/context"
	"源码/beego/config"
	"源码/beego/session"
)

type Config struct {
	AppName			   	string // Application name
	RunMode				string // Running Mode: dev | prod
	RouterCaseSensitive bool
	ServerName			string
	RecoverPanic		bool
	RecoverFunc			func(*context.Context)
	CopyRequestBody		bool
	EnableGzip			bool
	MaxMemory			int64
	EnableErrorsShow	bool
	EnableErrorsRender  bool
	Listen				Listen
	WebConfig			WebConfig
	Log					LogConfig
}

// Listen holds for http and https related config
type Listen struct {
	Graceful          bool // Graceful means use graceful module to start the server
	ServerTimeOut     int64
	ListenTCP4        bool
	EnavleHTTP		  bool
	HTTPAddr		  string
	HTTPPort		  int
	EnableHTTPS		  bool
	EnableMutualHTTPS bool
	HTTPSAddr		  string
	HTTPSPort		  int
	HTTPSCertFile	  string
	EnableAdmin		  bool
	AdminPort		  int
	EnableFcgi		  bool
	EnableStdIo		  bool // EnavleStdIo works with EnableFcgi Use FCGI via standard I/O
}

// WebConfig holds web related config
type WebConfig struct {
	AutoRender				bool
	EnableDocs				bool
	FlashName				string
	FlashSeparator			string
	DirectoryIndex			bool
	StaticDir				map[string]string
	StaticExtensionsToGzip	[]string
	TemplateLeft			string
	TemplateRight			string
	ViewsPath				string
	EnableXSRF				bool
	XSRFKey					string
	XSRFExpire				int
	Session					SessionConfig
}

// SessionConfig holds session related config
type SessionConfig struct {
	SessionOn					 bool
	SessionProvider				 string
	SessionName					 string
	SessionGCMaxLifetime		 int64
	SessionProviderConfig		 string
	SessionCookieLifeTime		 int
	SessionAutoSetCookie		 bool
	SessionDomain				 string
	SessionDisableHTTPOnly 		 bool  // used to allow for cross domain cookies/javascript cookies.
	SessionEnableSidInHTTPHeader bool	// enavles store/get the sessionId into/from http headers
	SessionNameInHTTPHeader		 string
	SessionEnableSidInURLQuery	 bool  // enable get the sessionId from Url Query params
}

// LogConfig holds Log related config
type LogConfig struct {
	AccessLogs			bool
	AccessLogsFormat	string // access log format: JSON_FORMAT, APACHR_FORMAT or empty string
	FileLineNum			bool
	Outputs				map[string]string // Store Adaptor: config
}
var (
	// BConfig is the default config for Application
	BConfig *Config

	// AppConfig is the instance of Config, store the config information from file
	AppConfig *beegoAppConfig

	// GlobalSessions is the instance for the session manager
	GlobalSessions *session.Manager

	// appConfigProvider is the provider for the config, default is ini
	appConfigProvider = "ini"
)

func LoadAppConfig(adapterName, configPath string) error {

}

type beegoAppConfig struct {
	innerConfig config.Configer
}

func (b *beegoAppConfig) String(key string) string {

}