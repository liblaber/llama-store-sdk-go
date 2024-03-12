package llamastoreconfig

type Config struct {
	BaseUrl     *string
	AccessToken *string
}

func NewConfig() Config {
	baseUrl := "http://localhost:8080"
	newConfig := Config{
		BaseUrl: &baseUrl,
	}

	return newConfig
}

func (c *Config) SetBaseUrl(baseUrl string) {
	c.BaseUrl = &baseUrl
}

func (c *Config) SetAccessToken(accessToken string) {
	c.AccessToken = &accessToken
}
