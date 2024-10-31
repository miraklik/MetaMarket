package main

import (
	"admin_panel/config"
	"admin_panel/repository"
)

func main() {
	config.LoadConfig()
	repository.InitDB()
	repository.Migrate()
}
