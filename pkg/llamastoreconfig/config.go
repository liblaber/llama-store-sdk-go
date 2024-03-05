package llamastoreconfig

type Config struct {
	BaseUrl *string
}

func NewConfig() Config {
	baseUrl := ""
	newConfig := Config{
		BaseUrl: &baseUrl,
	}

	return newConfig
}

func (c *Config) SetBaseUrl(baseUrl string) {
	c.BaseUrl = &baseUrl
}
