package main

import (
	"github.com/Tverous/go-sss/config"
	"github.com/Tverous/go-sss/sss"
)

func main() {
	sss.GenerateShares(config.N, config.M, config.Secret)
}