package helper

import (
	"encoding/json"

	"app/pkg/config"
)

func Marshal(v any) ([]byte, error) {
	if config.AppConfig.MODE == "development" {
		return json.MarshalIndent(v, "", "  ")
	} else {
		return json.Marshal(v)
	}
}
