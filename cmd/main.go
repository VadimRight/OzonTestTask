package main

import (
	"github.com/VadimRight/OzonTestTask/bootstrap"
)

func main() {
	cfg := bootstrap.LoadConfig()
	_ = bootstrap.InitPostgresDatabase(cfg)
}
