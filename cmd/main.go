package main

import (
	"github.com/VadimRight/OzonTestTask/bootstrap"
)

func main() {
	cfg := bootstrap.LoadConfig()
	bootstrap.InitPostgresDatabase(cfg)
}
