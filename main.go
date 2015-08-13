package main

import (
	"news_worker/lib/config"
	"news_worker/lib/utils"
)

// main entrypoint of the app
func main() {
	utils.Info("init!")
	config.Start()
}
