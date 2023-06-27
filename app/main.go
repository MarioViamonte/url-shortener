package main

import (
	"url-shortener/model"
	"url-shortener/server"
)

func main() {
	model.Setup()
	server.SetupAndList()
}