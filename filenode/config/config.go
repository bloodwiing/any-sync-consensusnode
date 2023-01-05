package config

import (
	commonaccount "github.com/anytypeio/any-sync/accountservice"
	"github.com/anytypeio/any-sync/app"
	"github.com/anytypeio/any-sync/metric"
	"github.com/anytypeio/any-sync/net"
	"gopkg.in/yaml.v3"
	"os"
)

const CName = "config"

func NewFromFile(path string) (c *Config, err error) {
	c = &Config{}
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	if err = yaml.Unmarshal(data, c); err != nil {
		return nil, err
	}
	return
}

type Config struct {
	Account         commonaccount.Config `yaml:"account"`
	GrpcServer      net.Config           `yaml:"grpcServer"`
	Metric          metric.Config        `yaml:"metric"`
	FileStorePogreb FileStorePogreb      `yaml:"fileStorePogreb"`
}

func (c *Config) Init(a *app.App) (err error) {
	return
}

func (c Config) Name() (name string) {
	return CName
}

func (c Config) GetAccount() commonaccount.Config {
	return c.Account
}

func (c Config) GetFileStorePogreb() FileStorePogreb {
	return c.FileStorePogreb
}

func (c Config) GetNet() net.Config {
	return c.GrpcServer
}

func (c Config) GetMetric() metric.Config {
	return c.Metric
}
