package configs

import (
	"time"

	"github.com/spf13/viper"
)

type Setting struct {
	vp *viper.Viper
}

type ServerConfig struct {
	RunMode      string
	HttpPort     string
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

type Sqlite3Connection struct {
	DbType           string
	ConnectionString string
}

func InitSetting() (*Setting, error) {
	vp := viper.New()
	vp.SetConfigFile("./configs/config.yaml")
	//vp.SetConfigName("config")
	//vp.SetConfigType("yaml")
	err := vp.ReadInConfig()
	if err != nil {
		return nil, err
	}
	return &Setting{vp}, nil
}

func (s *Setting) ReadConfigs(k string, v interface{}) error {
	err := s.vp.UnmarshalKey(k, v)
	if err != nil {
		return err
	}
	return nil
}
