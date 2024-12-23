package gemini

import (
	"archetype/app/shared/configuration"
	"context"

	ioc "github.com/Ignaciojeria/einar-ioc/v2"
	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
)

func init() {
	ioc.Registry(newClient, configuration.NewGeminiConfiguration)
}
func newClient(conf configuration.GeminiConfiguration) (*genai.Client, error) {
	return genai.NewClient(context.Background(), option.WithAPIKey(conf.GEMINI_API_KEY))
}
