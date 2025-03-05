package main

import (
	"ekko/pkg/carbon/pkg/config"

	_ "github.com/go-sql-driver/mysql"

	"ekko/internal/server"
)

func main() {
	flags := config.ParseFlags()
	server.Run(flags)
}
