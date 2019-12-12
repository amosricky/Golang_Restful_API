package setting

import (
	"github.com/go-ini/ini"
	"github.com/sirupsen/logrus"
	"time"
)

type App struct {
	PageSize int
	RuntimeRootPath string
	JWTSalt string
}

type Server struct {
	RunMode string
	HttpPort int
	ReadTimeout time.Duration
	WriteTimeout time.Duration
	RandomMin int
	RandomMax int
}

type Database struct {
	Type string
	Path string
	Name string
	TablePrefix string
}

type Log struct {
	LogSavePath string
	LogPrefix string
	LogFileExtension string
	TimeFormat string
}

type Image struct {
	ImageMaxSize int
	ImageSavePath string
	ImageAllowExtension []string
	ImagePrefixUrl string
}

var AppSetting = &App{}
var ServerSetting = &Server{}
var DatabaseSetting = &Database{}
var LogSetting = &Log{}
var ImageSetting = &Image{}

var cfg *ini.File

// Setup initialize the configuration instance
func Setup() {
	var err error
	cfg, err = ini.Load("conf/app.ini")
	if err != nil {
		logrus.Fatalf("setting.Setup, fail to parse 'conf/app.ini': %v", err)
	}

	mapTo("app", AppSetting)
	mapTo("server", ServerSetting)
	mapTo("database", DatabaseSetting)
	mapTo("log", LogSetting)
	mapTo("image", ImageSetting)

	ImageSetting.ImageMaxSize = ImageSetting.ImageMaxSize * 1024 * 1024
	ServerSetting.ReadTimeout = ServerSetting.ReadTimeout * time.Second
	ServerSetting.WriteTimeout = ServerSetting.WriteTimeout * time.Second
}

// mapTo map section
func mapTo(section string, v interface{}) {
	err := cfg.Section(section).MapTo(v)
	if err != nil {
		logrus.Fatalf("Cfg.MapTo %s err: %v", section, err)
	}
}