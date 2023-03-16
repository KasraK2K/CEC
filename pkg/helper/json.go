package helper

import (
	"CEC/pkg/config"
	"encoding/json"
)

func Marshal(v any) ([]byte, error) {
	if config.AppConfig.MODE == "development" {
		return json.MarshalIndent(v, "", "  ")
	} else {
		return json.Marshal(v)
	}
}
