package config

import (
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"time"

	"github.com/spf13/viper"
)

var (
	SECRET_JWT = ""
)

type HttpServer struct {
	AppVersion        string
	Port              string
	PprofPort         string
	Timeout           time.Duration
	ReadTimeout       time.Duration
	WriteTimeout      time.Duration
	CookieLifeTime    int
	SessionCookieName string
	CSRFHeader        string
}

type DatabaseConfig struct {
	Username string
	Password string
	Hostname string
	Port     int
	Name     string
}

type LoggerConfig struct {
	Development       bool
	DisableCaller     bool
	DisableStacktrace bool
	Encoding          string
	Level             string
}

type Metrics struct {
	Port        string
	URL         string
	ServiceName string
}

type JaegerConfig struct {
	Host        string
	ServiceName string
	LogSpans    bool
}

type AppConfig struct {
	AppPath      string
	JWTSecret    string
	AESGCMSecret string
	DB           DatabaseConfig
	HttpServer   HttpServer
	Logger       LoggerConfig
	Jaeger       JaegerConfig
	Metrics      Metrics
}

func InitConfig() *AppConfig {
	return ReadEnv()
}

func ReadEnv() *AppConfig {
	app := &AppConfig{}
	isRead := true

	if isRead {
		env := os.Getenv("APP_ENV")
		if env == "" {
			env = "localhost"
		}

		dir, _ := dirname()
		viper.SetConfigName(fmt.Sprintf("config-%s", env))
		viper.AddConfigPath(dir)
		viper.SetConfigType("yaml")

		err := viper.ReadInConfig()
		if err != nil {
			log.Println("error read config : ", err.Error())
			return nil
		}

		app.JWTSecret = viper.GetString("JWT_KEY")
		app.AESGCMSecret = viper.GetString("AESGCMSECRET")
		app.AppPath = viper.GetString("APP_PATH")

		app.DB.Username = viper.GetString("DATABASE.DBUSER")
		app.DB.Password = viper.GetString("DATABASE.DBPASS")
		app.DB.Hostname = viper.GetString("DATABASE.DBHOST")
		app.DB.Port = viper.GetInt("DATABASE.DBPORT")
		app.DB.Name = viper.GetString("DATABASE.DBNAME")

		app.HttpServer.AppVersion = viper.GetString("HTTP_SERVER.AppVersion")
		app.HttpServer.Port = viper.GetString("HTTP_SERVER.Port")
		app.HttpServer.PprofPort = viper.GetString("HTTP_SERVER.PprofPort")
		app.HttpServer.Timeout = time.Duration(viper.GetInt("HTTP_SERVER.Timeout")) * time.Second
		app.HttpServer.ReadTimeout = time.Duration(viper.GetInt("HTTP_SERVER.ReadTimeout")) * time.Second
		app.HttpServer.WriteTimeout = time.Duration(viper.GetInt("HTTP_SERVER.WriteTimeout")) * time.Second
		app.HttpServer.CookieLifeTime = viper.GetInt("HTTP_SERVER.CookieLifeTime")
		app.HttpServer.SessionCookieName = viper.GetString("HTTP_SERVER.SessionCookieName")
		app.HttpServer.CSRFHeader = viper.GetString("HTTP_SERVER.CSRFHeader")

		app.Logger.Development = viper.GetBool("LOGGER.DEVELOPMENT")
		app.Logger.DisableCaller = viper.GetBool("LOGGER.DISABLECALLER")
		app.Logger.DisableStacktrace = viper.GetBool("LOGGER.DISABLESTACKTRACE")
		app.Logger.Encoding = viper.GetString("LOGGER.ENCODING")
		app.Logger.Level = viper.GetString("LOGGER.LEVEL")

		app.Jaeger.Host = viper.GetString("JAEGER.HOST")
		app.Jaeger.ServiceName = viper.GetString("JAEGER.SERVICENAME")
		app.Jaeger.LogSpans = viper.GetBool("JAEGER.LOGSPANS")

		app.Metrics.Port = viper.GetString("METRICS.Port")
		app.Metrics.URL = viper.GetString("METRICS.URL")
		app.Metrics.ServiceName = viper.GetString("METRICS.ServiceName")
	}

	return app
}

func filename() (string, error) {
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		return "", errors.New("unable to get the current filename")
	}
	return filename, nil
}

func dirname() (string, error) {
	filename, err := filename()
	if err != nil {
		return "", err
	}
	return filepath.Dir(filename), nil
}
