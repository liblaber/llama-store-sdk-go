package llamastore

import (
	"github.com/liblaber/llama-store-sdk-go/internal/configmanager"
	"github.com/liblaber/llama-store-sdk-go/pkg/llama"
	"github.com/liblaber/llama-store-sdk-go/pkg/llamapicture"
	"github.com/liblaber/llama-store-sdk-go/pkg/llamastoreconfig"
	"github.com/liblaber/llama-store-sdk-go/pkg/token"
	"github.com/liblaber/llama-store-sdk-go/pkg/user"
)

type LlamaStore struct {
	LlamaPicture *llamapicture.LlamaPictureService
	Llama        *llama.LlamaService
	Token        *token.TokenService
	User         *user.UserService
	manager      *configmanager.ConfigManager
}

func NewLlamaStore(config llamastoreconfig.Config) *LlamaStore {
	manager := configmanager.NewConfigManager(config)
	return &LlamaStore{
		LlamaPicture: llamapicture.NewLlamaPictureService(manager),
		Llama:        llama.NewLlamaService(manager),
		Token:        token.NewTokenService(manager),
		User:         user.NewUserService(manager),
		manager:      manager,
	}
}

func (l *LlamaStore) SetBaseUrl(baseUrl string) {
	l.manager.SetBaseUrl(baseUrl)
}

func (l *LlamaStore) SetAccessToken(accessToken string) {
	l.manager.SetAccessToken(accessToken)
}
