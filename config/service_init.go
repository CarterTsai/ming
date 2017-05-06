package config

import (
	"encoding/json"
	"time"

	"gopkg.in/yaml.v2"
)

// InitConfig output init config data for json or yaml
func InitConfig(format string) string {
	data := map[string]interface{}{
		"Project":   "EC",
		"Name":      "Hello",
		"Copyright": "Copyright © " + time.Now().Format("2006"),
		"Functions": []string{"Cache", "Content", "Directory"},
	}

	var convertData []byte
	var err error

	if format == "json" {
		convertData, err = json.Marshal(data)
		if err != nil {
			panic(err)
		}
	} else {
		convertData, err = yaml.Marshal(data)
		if err != nil {
			panic(err)
		}

	}
	return string(convertData)
}
