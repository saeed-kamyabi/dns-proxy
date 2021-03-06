package usecase

import (
	"encoding/json"
	"github.com/saeed-kamyabi/dns-proxy/config/domain"
	"io/ioutil"
	"regexp"
)

type config struct {
	path string
	cfg  *domain.Config
}

func NewConfig(path string) domain.UseCaseModel {
	cfg := new(domain.Config)
	b, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(b, cfg)
	if err != nil {
		panic(err)
	}
	cfg.Initialize()
	uc := &config{
		path: path,
		cfg:  cfg,
	}
	return uc
}

func (c *config) GetIP(domain string) (ip string) {
	for k, v := range c.cfg.Domains {
		match, _ := regexp.MatchString(k+"\\.", domain)
		if match {
			return v
		}
	}
	return ""
}

func (c *config) NextServer() interface{} {
	s := c.cfg.Servers.Next()
	if s == nil {
		return nil
	}
	return s
}

func (c *config) GetNameServer() domain.NameServer {
	return c.cfg.NameServers[0]
}

func (c *config) GetSecondaryNameServers() []domain.NameServer {
	return c.cfg.NameServers[1:]
}
