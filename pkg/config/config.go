package Config

import (
	"time"

	"go.mongodb.org/mongo-driver/mongo/options"
)

type Config struct {
	Uri            string
	Direct         bool
	ConnectTimeout time.Duration
}

func New() *Config {
	c := new(Config)
	c.Uri = "mongodb://localhost:27017"
	c.Direct = true
	c.ConnectTimeout = 20
	return c
}

func (c *Config) SetUri(uri string) *Config {
	c.Uri = uri
	return c
}

func (c *Config) SetDirect(direct bool) *Config {
	c.Direct = direct
	return c
}

func (c *Config) SetConnectTimeout(t time.Duration) *Config {
	c.ConnectTimeout = t
	return c
}

func (c Config) GetOptions() *options.ClientOptions {
	opts := options.Client()
	opts.SetDirect(c.Direct)
	opts.ApplyURI(c.Uri)
	// opts.SetConnectTimeout(c.ConnectTimeout)

	return opts
}
