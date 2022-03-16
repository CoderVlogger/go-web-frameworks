package internal

type Config struct {
	ServerPort int    `split_words:"true" default:"8080"`
	ServerHost string `split_words:"true" default:"localhost"`
	PageSize   int    `split_words:"true" default:"10"`
}

func (c Config) GetServerPort() int {
	return c.ServerPort
}

func (c Config) GetServerHost() string {
	return c.ServerHost
}

func (c Config) GetPageSize() int {
	return c.PageSize
}
