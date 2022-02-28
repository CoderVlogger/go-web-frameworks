package internal

type Config struct {
	ServerPort int    `split_words:"true" default:"8080"`
	ServerHost string `split_words:"true"`
	PageSize   int    `split_words:"true" default:"10"`
}

func (c Config) GetPageSize() int {
	return c.PageSize
}
