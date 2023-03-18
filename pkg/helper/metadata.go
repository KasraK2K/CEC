package helper

import "app/pkg/config"

type metaData struct {
	BACKEND_VERSION  string `json:"backend_version" bson:"backend_version"`
	FRONTEND_VERSION string `json:"frontend_version" bson:"frontend_version"`
	APP_VERSION      string `json:"app_version" bson:"app_version"`
	MODE             string `json:"mode" bson:"mode"`
	SUCCESS          bool   `json:"success" bson:"success"`
	RESULT           any    `json:"result,omitempty" bson:"result,omitempty"`
	ERRORS           any    `json:"errors" bson:"errors,omitempty"`
}

func AddMetaData(data any, errors ...bool) metaData {
	response := metaData{
		BACKEND_VERSION:  config.AppConfig.BACKEND_VERSION,
		FRONTEND_VERSION: config.AppConfig.FRONTEND_VERSION,
		APP_VERSION:      config.AppConfig.APP_VERSION,
		MODE:             config.AppConfig.MODE,
	}

	if len(errors) > 0 && errors[0] {
		response.SUCCESS = false
		response.ERRORS = data
	} else {
		response.SUCCESS = true
		response.RESULT = data
	}

	return response
}
