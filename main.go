package main

import (
	"github.com/mipimentel/gin-api-rest/database"
	"github.com/mipimentel/gin-api-rest/routes"
)

func main() {
	database.ConectaComBancoDeDados()
	routes.HandleRequests()
}
