package configuration

import (
	ioc "github.com/Ignaciojeria/einar-ioc/v2"
)

type GeminiConfiguration struct {
	GEMINI_API_KEY string `required:"true"`
}

func init() {
	ioc.Registry(NewGeminiConfiguration, NewEnvLoader)
}
func NewGeminiConfiguration(env EnvLoader) (GeminiConfiguration, error) {
	conf := GeminiConfiguration{
		GEMINI_API_KEY: env.Get("GEMINI_API_KEY"),
	}
	return validateConfig(conf)
}
