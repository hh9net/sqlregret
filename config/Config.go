package config

import (
	"io/ioutil"

	"github.com/cihub/seelog"

	"encoding/json"
)

type Config struct {
	Mode              string `json:"mode"` // online:实时同步 onfile:读取文件
	Destination       string `json:"destination"`
	SlaveId           int    `json:"slaveId"`
	BasePath          string `json:"basePath"`
	IndexFile         string `json:"indexFile"`
	MasterAddress     string `json:"masterAddress"`
	MasterPort        int    `json:"masterPort"`
	MasterJournalName string `json:"masterJournalName"`
	MasterPosition    int    `json:"masterPosition"`
	DbUsername        string `json:"dbUsername"`
	DbPassword        string `json:"dbPassword"`
	DefaultDbName     string `json:"defaultDbName"`
	LimitShowRow      int    `json:"limitShowRow"` // 在pre模式下，影响行数超过此值的予以显示
}

func ParseConfigData(data []byte) (*Config, error) {
	var cfg Config
	if err := json.Unmarshal([]byte(data), &cfg); err != nil {
		return nil, err
	}

	seelog.Debugf("运行模式:%s", cfg.Mode)
	return &cfg, nil
}

func ParseConfigFile(fileName string) (*Config, error) {
	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		return nil, err
	}

	return ParseConfigData(data)
}
