package main

import (
	"fmt"
	"ldap-auth/config"
	"ldap-auth/routes"
)

func main() {
	config.Load()

	router := routes.SetupRoutes()
	router.Run(fmt.Sprintf(":%s", config.C.Server.Port))
}
