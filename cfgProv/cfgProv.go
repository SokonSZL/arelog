package cfgProv

import (
)

type Config struct {
	LogSavePath          string //if "" use the same path as config
	TXCall               string
	ButtonPrimaryColor   string
	ButtonSecondaryColor string
}

func (c *Config) GetTXCall() string {
	return c.TXCall
}
