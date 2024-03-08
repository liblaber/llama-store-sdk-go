package configmanager

import "github.com/liblaber/llama-store-sdk-go/pkg/llamastoreconfig"

type ConfigManager struct {
	LlamaPicture llamastoreconfig.Config
	Llama        llamastoreconfig.Config
	Token        llamastoreconfig.Config
	User         llamastoreconfig.Config
}

func NewConfigManager(config llamastoreconfig.Config) *ConfigManager {
	return &ConfigManager{
		LlamaPicture: config,
		Llama:        config,
		Token:        config,
		User:         config,
	}
}

func (c *ConfigManager) SetBaseUrl(baseUrl string) {
	c.LlamaPicture.SetBaseUrl(baseUrl)
	c.Llama.SetBaseUrl(baseUrl)
	c.Token.SetBaseUrl(baseUrl)
	c.User.SetBaseUrl(baseUrl)
}

func (c *ConfigManager) SetAccessToken(accessToken string) {
	c.LlamaPicture.SetAccessToken(accessToken)
	c.Llama.SetAccessToken(accessToken)
	c.Token.SetAccessToken(accessToken)
	c.User.SetAccessToken(accessToken)
}

func (c *ConfigManager) UpdateAccessToken(originalValue string, newValue string) {

	if c.LlamaPicture.AccessToken != nil && *c.LlamaPicture.AccessToken == originalValue {
		c.LlamaPicture.SetAccessToken(newValue)
	}

	if c.Llama.AccessToken != nil && *c.Llama.AccessToken == originalValue {
		c.Llama.SetAccessToken(newValue)
	}

	if c.Token.AccessToken != nil && *c.Token.AccessToken == originalValue {
		c.Token.SetAccessToken(newValue)
	}

	if c.User.AccessToken != nil && *c.User.AccessToken == originalValue {
		c.User.SetAccessToken(newValue)
	}
}

func (c *ConfigManager) GetLlamaPicture() *llamastoreconfig.Config {
	return &c.LlamaPicture
}
func (c *ConfigManager) GetLlama() *llamastoreconfig.Config {
	return &c.Llama
}
func (c *ConfigManager) GetToken() *llamastoreconfig.Config {
	return &c.Token
}
func (c *ConfigManager) GetUser() *llamastoreconfig.Config {
	return &c.User
}
